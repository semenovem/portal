package conn

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/semenovem/portal/pkg"
)

type RedisProps struct {
	Host     string
	Password string
	DBName   uint16
}

func InitRedis(ctx context.Context, logger pkg.Logger, c *RedisProps) (*redis.Client, error) {
	options := redis.Options{
		Addr:     c.Host,
		Password: c.Password,
		DB:       int(c.DBName),
	}

	client := redis.NewClient(&options)

	go func() {
		<-ctx.Done()
		if err := client.Close(); err != nil {
			logger.Named("InitRedis.CLose").Error(err.Error())
		}
	}()

	return client, nil
}
