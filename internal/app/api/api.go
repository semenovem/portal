package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/semenovem/portal/config"
	"github.com/semenovem/portal/internal/action"
	"github.com/semenovem/portal/internal/provider"
	"github.com/semenovem/portal/internal/rest/router"
	"github.com/semenovem/portal/internal/zoo/conn"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/it"
	"math"
	"time"
)

type appAPI struct {
	ctx    context.Context
	logger pkg.Logger
	db     *pgxpool.Pool
	redis  *redis.Client
	router *router.Router
}

func New(ctx context.Context, logger pkg.Logger, cfg config.API) error {
	var (
		ll  = logger.Named("appAPI.NewPeoplePvd")
		app = appAPI{
			ctx:    ctx,
			logger: logger,
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
			ll.Named("ConnectDBPostgresSQL").DBTag().Error(err.Error())
		}

		if err = conn.Migrate(ll, db, cfg.DBMigrationsDir); err != nil {
			ll.Named("Migrate").Nested(err.Error())
		}
	}()

	// Redis
	if app.redis, err = conn.InitRedis(ctx, ll, cfg.RedisConn.ConvTo()); err != nil {
		ll.Named("InitRedis").Nested(err.Error())
		return err
	}

	// Провайдеры данных
	var (
		audit   = provider.NewAudit(app.db, logger)
		authPvd = provider.NewAuthPvd(
			logger,
			app.db,
			app.redis,
			time.Minute*time.Duration(cfg.JWT.AccessTokenLifetimeMin),
		)
		peoplePvd = provider.NewPeoplePvd(app.db, logger)
	)

	t, err := authPvd.Now(ctx)
	if err != nil {
		ll.Named("authPvd.Now").Nested(err.Error())
		return err
	}

	if math.Abs(float64(t.Unix()-time.Now().Unix())) > 60 {
		err = errors.New("the time in the database differs by more than a minute")
		ll.Error(err.Error())
		return err
	}

	// Экшены
	authAct := action.NewAuth(
		logger,
		it.NewUserPasswdAuth(cfg.PasswdSalt),
		audit,
		authPvd,
		peoplePvd,
	)

	// Router
	if app.router, err = router.New(
		ctx,
		logger,
		cfg,
		app.redis,
		authPvd,
		peoplePvd,
		authAct,
	); err != nil {
		ll.Named("router").Nested(err.Error())
		return err
	}

	go app.router.Start()

	return nil
}
