package config

import (
	"github.com/semenovem/portal/internal/zoo/conn"
	"time"
)

type Base struct {
	Env      string `env:"ENV,required"`
	CliMode  bool   `env:"CLI_MODE,required"`
	LogLevel int8   `env:"LOG_LEVEL,required"`
}

type RedisConn struct {
	Host     string `env:"REDIS_HOST,required"`
	Password string `env:"REDIS_PASSWORD,required"`
	DBName   uint16 `env:"REDIS_DB_NAME,required"`
}

func (c *RedisConn) ConvTo() *conn.RedisConfig {
	return &conn.RedisConfig{
		Host:     c.Host,
		Password: c.Password,
		DBName:   c.DBName,
	}
}

type DatabaseCoreConn struct {
	Host               string `env:"DB_CORE_HOST,required"`
	Port               uint16 `env:"DB_CORE_PORT,required"`
	Name               string `env:"DB_CORE_NAME,required"`
	User               string `env:"DB_CORE_USER,required"`
	Password           string `env:"DB_CORE_PASSWORD,required"`
	SSLMode            string `env:"DB_CORE_SSL_MODE,required"`
	MaxIdleConns       uint16 `env:"DB_CORE_MAX_IDLE_CONNS,required"`
	MaxOpenConns       uint16 `env:"DB_CORE_MAX_OPEN_CONNS,required"`
	MaxLifetimeSec     uint16 `env:"DB_CORE_MAX_LIFETIME_SEC,required"`
	MaxIdleLifetimeSec uint16 `env:"DB_CORE_MAX_IDLE_LIFETIME_SEC,required"`
	Schema             string `env:"DB_CORE_SCHEMA,required"`
	AppName            string `env:"DB_CORE_APPLICATION_NAME,required"`
}

func (c *DatabaseCoreConn) ConvTo() *conn.DBPGConfig {
	return &conn.DBPGConfig{
		Host:            c.Host,
		Port:            c.Port,
		Name:            c.Name,
		User:            c.User,
		Password:        c.Password,
		SSLMode:         c.SSLMode,
		Schema:          c.Schema,
		AppName:         c.AppName,
		MaxIdleConns:    c.MaxIdleConns,
		MaxOpenConns:    c.MaxOpenConns,
		MaxLifetime:     time.Second * time.Duration(c.MaxLifetimeSec),
		MaxIdleLifetime: time.Second * time.Duration(c.MaxIdleLifetimeSec),
	}
}
