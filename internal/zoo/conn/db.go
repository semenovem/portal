package conn

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/semenovem/portal/pkg"
	"time"
)

type DBPGProps struct {
	Host            string
	Port            uint16
	Name            string
	User            string
	Password        string
	SSLMode         string
	Schema          string
	AppName         string
	MaxIdleConns    uint16
	MaxOpenConns    uint16
	MaxLifetime     time.Duration
	MaxIdleLifetime time.Duration
}

func (c *DBPGProps) validate() error {
	if c.MaxLifetime == 0 {
		return errors.New("no value specified for MaxLifetime")
	}

	if c.MaxIdleLifetime == 0 {
		return errors.New("no value specified for MaxIdleLifetime")
	}

	if c.MaxIdleConns == 0 {
		return errors.New("no value specified for MaxIdleConns")
	}

	if c.MaxOpenConns == 0 {
		return errors.New("no value specified for MaxOpenConns")
	}

	return nil
}

func ConnectDBPostgres(ctx context.Context, cfg *DBPGProps) (*pgxpool.Pool, error) {
	if err := cfg.validate(); err != nil {
		return nil, err
	}

	dataSourceName := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=%s search_path=%s application_name=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SSLMode,
		cfg.Schema,
		cfg.AppName,
	)

	c, err := pgxpool.ParseConfig(dataSourceName)
	if err != nil {
		return nil, err
	}

	c.MaxConns = int32(cfg.MaxOpenConns)
	c.MinConns = int32(cfg.MaxIdleConns)
	c.MaxConnLifetime = cfg.MaxLifetime
	c.MaxConnIdleTime = cfg.MaxIdleLifetime

	pool, err := pgxpool.NewWithConfig(context.Background(), c)
	if err != nil {
		return nil, err
	}

	go func() {
		<-ctx.Done()
		pool.Close()
	}()

	return pool, nil
}

func ConnectDBPostgresSQL(ctx context.Context, logger pkg.Logger, cfg *DBPGProps) (*sql.DB, error) {
	ll := logger.Named("ConnectDBPostgresSQL")

	if err := cfg.validate(); err != nil {
		ll.Named("validate").Error(err.Error())
		return nil, err
	}

	dataSourceName := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=%s search_path=%s application_name=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SSLMode,
		cfg.Schema,
		cfg.AppName,
	)

	ll.Debugf(dataSourceName)

	c, err := pgx.ParseConfig(dataSourceName)
	if err != nil {
		ll.Named("ParseConfig").Error(err.Error())
		return nil, err
	}

	db := stdlib.OpenDB(*c)

	go func() {
		<-ctx.Done()
		if err1 := db.Close(); err1 != nil {
			ll.Named("ConnectDBPostgresSQL.DATABASE").Error(err.Error())
		}
	}()

	return db, nil
}
