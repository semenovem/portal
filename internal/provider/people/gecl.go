package peopleprovider

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/semenovem/portal/pkg"
)

type Provider struct {
	db     *pgxpool.Pool
	logger pkg.Logger
}

func New(db *pgxpool.Pool, logger pkg.Logger) *Provider {
	return &Provider{
		db:     db,
		logger: logger.Named("peopleprovider"),
	}
}
