package store_provider

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/semenovem/portal/internal/abc/provider"
	"github.com/semenovem/portal/pkg"
	"time"
)

type StoreProvider struct {
	logger                    pkg.Logger
	db                        *pgxpool.Pool
	redis                     *redis.Client
	draftStoreDefaultLifetime time.Duration // Время хранения по умолчанию, если не указано
	draftStoreMaxLifetime     time.Duration // Максимальное время хранения
}

func New(
	logger pkg.Logger,
	db *pgxpool.Pool,
	redisClient *redis.Client,
	draftStoreDefaultLifetime time.Duration,
	draftStoreMaxLifetime time.Duration,
) *StoreProvider {
	return &StoreProvider{
		logger:                    logger.Named("StoreProvider"),
		db:                        db,
		redis:                     redisClient,
		draftStoreDefaultLifetime: draftStoreDefaultLifetime,
		draftStoreMaxLifetime:     draftStoreMaxLifetime,
	}
}

func getArbitraryDataKeyName(userID uint32, path string) string {
	return fmt.Sprintf("user_storage_%d_%s", userID, path)
}

func (p *StoreProvider) StoreArbitraryData(ctx context.Context, userID uint32, path, payload string) error {
	key := getArbitraryDataKeyName(userID, path)

	err := p.redis.Set(ctx, key, payload, p.draftStoreDefaultLifetime).Err()
	if err != nil {
		p.logger.Named("Set").Redis(err)
	}

	return err
}

func (p *StoreProvider) LoadArbitraryData(ctx context.Context, userID uint32, path string) (string, error) {
	key := getArbitraryDataKeyName(userID, path)

	payload, err := p.redis.Get(ctx, key).Result()
	if err != nil {
		if !provider.IsNoRec(err) {
			p.logger.Named("Get").Redis(err)
		}

		return "", err
	}

	return payload, nil
}

func (p *StoreProvider) DeleteArbitraryData(ctx context.Context, userID uint32, path string) error {
	key := getArbitraryDataKeyName(userID, path)

	res, err := p.redis.Del(ctx, key).Result()
	if err != nil {
		p.logger.Named("Del").Redis(err)
		return err
	}

	if res == 0 {
		return redis.Nil
	}

	return err
}
