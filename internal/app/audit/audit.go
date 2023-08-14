package auditapp

import (
	"context"
	"github.com/semenovem/portal/config"
	"github.com/semenovem/portal/pkg"
)

type appAudit struct {
	ctx    context.Context
	logger pkg.Logger
}

func New(ctx context.Context, logger pkg.Logger, cfg config.Audit) error {
	var (
		ll  = logger.Named("appAudit.New")
		app = appAudit{
			ctx:    ctx,
			logger: logger,
		}
	)

	ll.Info("start")
	app.logger.Info("start")
	//
	//// Провайдеры данных
	//var (
	//	audit   = provider.NewAudit(app.db, logger)
	//	authPvd = provider.NewAuthPvd(
	//		logger,
	//		app.db,
	//		app.redis,
	//		time.Minute*time.Duration(cfg.JWT.AccessTokenLifetimeMin),
	//	)
	//	peoplePvd = provider.NewPeoplePvd(app.db, logger)
	//)

	return nil
}
