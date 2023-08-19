package auth_provider

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/semenovem/portal/pkg"
	"time"
)

type AuthProvider struct {
	logger                    pkg.Logger
	db                        *pgxpool.Pool
	redis                     *redis.Client
	jwtAccessTokenLifetimeMin time.Duration
}

func NewAuthPvd(
	logger pkg.Logger,
	db *pgxpool.Pool,
	redisClient *redis.Client,
	jwtAccessTokenLifetimeMin time.Duration,
) *AuthProvider {
	return &AuthProvider{
		logger:                    logger.Named("authPvd"),
		db:                        db,
		redis:                     redisClient,
		jwtAccessTokenLifetimeMin: jwtAccessTokenLifetimeMin,
	}
}

func (p *AuthProvider) Now(ctx context.Context) (time.Time, error) {
	var t time.Time

	err := p.db.QueryRow(ctx, "SELECT now()").Scan(&t)
	if err != nil {
		p.logger.Named("Now").DBTag().Error(err.Error())
	}

	return t, err
}
