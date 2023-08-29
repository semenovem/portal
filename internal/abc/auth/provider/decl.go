package auth_provider

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/semenovem/portal/pkg"
	"time"
)

type AuthProvider struct {
	logger                 pkg.Logger
	db                     *pgxpool.Pool
	redis                  *redis.Client
	jwtAccessTokenLifetime time.Duration
	onetimeEntryLifetime   time.Duration
}

func New(
	logger pkg.Logger,
	db *pgxpool.Pool,
	redisClient *redis.Client,
	jwtAccessTokenLifetimeMin time.Duration,
	onetimeEntryLifetime time.Duration,
) *AuthProvider {
	return &AuthProvider{
		logger:                 logger.Named("AuthProvider"),
		db:                     db,
		redis:                  redisClient,
		jwtAccessTokenLifetime: jwtAccessTokenLifetimeMin,
		onetimeEntryLifetime:   onetimeEntryLifetime,
	}
}

func (p *AuthProvider) Now(ctx context.Context) (time.Time, error) {
	var (
		t   time.Time
		err = p.db.QueryRow(ctx, "SELECT now()").Scan(&t)
	)

	if err != nil {
		p.logger.Named("Now").DB(err)
	}

	return t, err
}

func getOnetimeEntryKeyName(id uuid.UUID) string {
	return fmt.Sprintf("onetime_entry_%s", id.String())
}
