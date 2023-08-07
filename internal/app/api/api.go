package api

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/semenovem/portal/config"
	"github.com/semenovem/portal/internal/rest/router"
	"github.com/semenovem/portal/internal/zoo/conn"
	"github.com/semenovem/portal/pkg"
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
		ll  = logger.Named("appAPI.New")
		app = appAPI{
			ctx:    ctx,
			logger: logger,
		}
		err error
	)

	app.db, err = conn.ConnectDBPostgres(ctx, cfg.DBCoreConn.ConvTo())
	if err != nil {
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
			ll.Named("ConnectDBPostgresSQL").Error(err.Error())
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

	// Router
	app.router = router.New(&router.Config{
		Ctx:    ctx,
		Logger: logger,
		DB:     app.db,
		Redis:  app.redis,
		Global: &cfg,
	})

	go app.router.Start()

	return nil
}
