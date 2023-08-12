package authprovider

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/entity"
)

type Provider struct {
	db     *pgxpool.Pool
	logger pkg.Logger
}

func New(db *pgxpool.Pool, logger pkg.Logger) *Provider {
	return &Provider{
		db:     db,
		logger: logger.Named("authprovider"),
	}
}

func (p *Provider) GetUserByEmail(ctx context.Context) (*entity.User, error) {

	return nil, nil
}
