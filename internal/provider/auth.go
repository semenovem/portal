package provider

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/it"
)

type AuthPvd struct {
	db     *pgxpool.Pool
	logger pkg.Logger
}

func NewAuthPvd(db *pgxpool.Pool, logger pkg.Logger) *AuthPvd {
	return &AuthPvd{
		db:     db,
		logger: logger.Named("authPvd"),
	}
}

func (p *AuthPvd) GetUserByEmail(ctx context.Context) (*it.User, error) {

	return nil, nil
}
