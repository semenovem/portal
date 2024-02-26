package app_platform

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/semenovem/portal/config"
	"github.com/semenovem/portal/internal/abc"
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

type platformApp struct {
	ctx             context.Context
	logger          pkg.Logger
	db              *pgxpool.Pool
	redis           *redis.Client
	s3              *s3.Service
	router          *router.Router
	config          *config.Platform
	auditService    *audit.AuditProvider
	failService     *fail.Service
	jwtService      *jwtoken.Service
	loginPasswdAuth it.LoginPasswdAuthenticator

	providers abc.Providers
	actions   abc.Actions
}

func New(ctx context.Context, logger pkg.Logger, cfg *config.Platform) error {
	var (
		ll  = logger.Named("platformApp.New")
		app = platformApp{
			ctx:             ctx,
			logger:          logger,
			config:          cfg,
			loginPasswdAuth: it.NewLoginPasswdAuthenticator(cfg.UserPasswdSalt),
		}
		err error

		dbc = cfg.DBCoreConn.ConvTo()
	)

	if app.db, err = conn.ConnectDBPostgres(ctx, dbc); err != nil {
		ll.Named("ConnectDBPostgres").Error(err.Error())
		return err
	}

	cstr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s&search_path=%s&application_name=%spool_max_conns=%d"+
			"",
		dbc.User,
		dbc.Password,
		dbc.Host,
		dbc.Port,
		dbc.Name,
		dbc.SSLMode,
		dbc.Schema,
		dbc.AppName,
		dbc.MaxOpenConns,
		//dbc.MaxLifetime,
		//dbc.MaxIdleLifetime,
	)

	//&pool_max_conn_lifetime=%s
	//&pool_max_conn_idle_time=%s

	stats := app.db.Stat()
	ll.Info(fmt.Sprintf(">>>> IdleConns  = %v", stats.IdleConns()))
	ll.Info(fmt.Sprintf(">>>> MaxConns   = %v", stats.MaxConns()))
	ll.Info(fmt.Sprintf(">>>> TotalConns = %v", stats.TotalConns()))
	ll.Info(fmt.Sprintf(">>>> TotalConns = %v", stats.TotalConns()))

	if err = conn.Migrate(ll, cstr, cfg.DBMigrationsDir); err != nil {
		ll.Named("Migrate").Nested(err)
	}

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
		AccessTokenSecret:    cfg.Auth.JWT.AccessTokenSecret,
		RefreshTokenSecret:   cfg.Auth.JWT.RefreshTokenSecret,
		AccessTokenLifetime:  cfg.Auth.JWT.AccessTokenLifetime.Val,
		RefreshTokenLifetime: cfg.Auth.JWT.RefreshTokenLifetime.Val,
	})

	app.auditService = audit.New(ctx, app.db, logger, cfg.GetGRPCAuditConfig())

	// Инициализация доменных областей
	if err = app.initDomains(); err != nil {
		ll.Named("initDomains").ErrorE(err)
		return err
	}

	// Router
	if app.router, err = router.New(
		ctx,
		app.logger,
		app.config,
		app.auditService,
		app.jwtService,
		app.loginPasswdAuth,

		&app.providers,
		&app.actions,
	); err != nil {
		ll.Named("router").Nested(err)
		return err
	}

	go app.router.Start()

	// Проверки
	t, err := app.providers.Auth.Now(ctx)
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
