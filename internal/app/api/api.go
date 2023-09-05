package apiapp

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/semenovem/portal/config"
	"github.com/semenovem/portal/internal/abc/auth/action"
	"github.com/semenovem/portal/internal/abc/auth/provider"
	"github.com/semenovem/portal/internal/abc/media/action"
	"github.com/semenovem/portal/internal/abc/media/provider"
	"github.com/semenovem/portal/internal/abc/people/action"
	"github.com/semenovem/portal/internal/abc/people/provider"
	"github.com/semenovem/portal/internal/abc/store/action"
	"github.com/semenovem/portal/internal/abc/store/provider"
	"github.com/semenovem/portal/internal/audit"
	"github.com/semenovem/portal/internal/router"
	"github.com/semenovem/portal/internal/s3"
	"github.com/semenovem/portal/internal/zoo/conn"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/fail"
	"github.com/semenovem/portal/pkg/it"
	"github.com/semenovem/portal/pkg/jwtoken"
	"math"
	"time"
)

type appAPI struct {
	ctx            context.Context
	logger         pkg.Logger
	db             *pgxpool.Pool
	redis          *redis.Client
	s3             *s3.Service
	router         *router.Router
	config         *config.API
	auditService   *audit.AuditProvider
	failService    *fail.Service
	jwtService     *jwtoken.Service
	userPasswdAuth it.UserPasswdAuthenticator

	providers struct {
		auth   *auth_provider.AuthProvider
		people *people_provider.PeopleProvider
		store  *store_provider.StoreProvider
		media  *media_provider.MediaProvider
	}

	actions struct {
		auth   *auth_action.AuthAction
		people *people_action.PeopleAction
		store  *store_action.StoreAction
		media  *media_action.MediaAction
	}
}

func New(ctx context.Context, logger pkg.Logger, cfg config.API) error {
	var (
		ll  = logger.Named("appAPI.New")
		app = appAPI{
			ctx:            ctx,
			logger:         logger,
			config:         &cfg,
			userPasswdAuth: it.NewUserPasswdAuth(cfg.UserPasswdSalt),
		}
		err error
	)

	if app.db, err = conn.ConnectDBPostgres(ctx, cfg.DBCoreConn.ConvTo()); err != nil {
		ll.Named("ConnectDBPostgres").Error(err.Error())
		return err
	}

	stats := app.db.Stat()
	ll.Info(fmt.Sprintf(">>>> IdleConns  = %v", stats.IdleConns()))
	ll.Info(fmt.Sprintf(">>>> MaxConns   = %v", stats.MaxConns()))
	ll.Info(fmt.Sprintf(">>>> TotalConns = %v", stats.TotalConns()))
	ll.Info(fmt.Sprintf(">>>> TotalConns = %v", stats.TotalConns()))

	// Миграции БД
	func() {
		var (
			ll              = ll.Named("migration")
			ctxnest, cancel = context.WithCancel(ctx)
		)
		defer cancel()

		db, err := conn.ConnectDBPostgresSQL(ctxnest, ll, cfg.DBCoreConn.ConvTo())
		if err != nil {
			ll.Named("ConnectDBPostgresSQL").DB(err)
			return
		}

		if err = conn.Migrate(ll, db, cfg.DBMigrationsDir); err != nil {
			ll.Named("Migrate").Nested(err)
		}
	}()

	// Redis
	if app.redis, err = conn.InitRedis(ctx, ll, cfg.RedisConn.ConvTo()); err != nil {
		ll.Named("InitRedis").Nested(err)
		return err
	}

	// S3
	if app.s3, err = s3.New(&s3.Props{
		Ctx:    app.ctx,
		Logger: app.logger,
		S3Conn: &app.config.S3Conn,
	}); err != nil {
		return ll.Named("s3").NestedWith(err, "can't create new S3 service")
	}

	app.jwtService = jwtoken.New(&jwtoken.Config{
		AccessTokenSecret:    cfg.JWT.AccessTokenSecret,
		RefreshTokenSecret:   cfg.JWT.RefreshTokenSecret,
		AccessTokenLifetime:  time.Minute * time.Duration(cfg.JWT.AccessTokenLifetimeMin),
		RefreshTokenLifetime: time.Hour * 24 * time.Duration(cfg.JWT.RefreshTokenLifetimeDay),
	})

	app.auditService = audit.New(ctx, app.db, logger, cfg.GetGRPCAuditConfig())

	// Провайдеры данных
	app.initProviders()

	// Экшены
	app.initActions()

	// Router
	if app.router, err = router.New(
		ctx,
		app.logger,
		app.config,
		app.auditService,
		app.jwtService,

		app.providers.auth,
		app.providers.people,

		app.actions.auth,
		app.actions.people,
		app.actions.store,
		app.actions.media,
	); err != nil {
		ll.Named("router").Nested(err)
		return err
	}

	go app.router.Start()

	// Проверки
	t, err := app.providers.auth.Now(ctx)
	if err != nil {
		ll.Named("authPvd.Now").Nested(err)
		return err
	}

	if math.Abs(float64(t.Unix()-time.Now().Unix())) > 60 {
		err = errors.New("the time in the database differs by more than a minute")
		ll.Error(err.Error())
		return err
	}

	return nil
}
