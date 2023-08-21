package media_provider

import (
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/semenovem/portal/pkg"
)

type MediaProvider struct {
	logger pkg.Logger
	db     *pgxpool.Pool
	redis  *redis.Client
}

func New(
	logger pkg.Logger,
	db *pgxpool.Pool,
	redisClient *redis.Client,
) *MediaProvider {
	return &MediaProvider{
		logger: logger.Named("MediaProvider"),
		db:     db,
		redis:  redisClient,
	}
}
