package config

import (
	"github.com/caarlos0/env/v9"
	"strings"
	"time"
)

type Platform struct {
	Base

	Rest       Rest
	DBCoreConn DatabaseCoreConn
	RedisConn  RedisConn

	DBMigrationsDir string `env:"DB_MIGRATIONS_DIR,required"`
	UserPasswdSalt  string `env:"USER_PASSWD_SALT,required"`

	Auth struct {
		// Длительность сессии одноразовой ссылки
		OnetimeEntryLifetime struct {
			Raw uint32 `env:"AUTH_ONETIME_ENTRY_LIFETIME_MIN,required"`
			Val time.Duration
		}

		JWT struct {
			AccessTokenSecret  string `env:"AUTH_JWT_ACCESS_TOKEN_SECRET,required"`
			RefreshTokenSecret string `env:"AUTH_JWT_REFRESH_TOKEN_SECRET,required"`

			AccessTokenLifetime struct {
				RawMin uint32 `env:"AUTH_JWT_ACCESS_TOKEN_LIFETIME_MIN,required"`
				Val    time.Duration
			}

			RefreshTokenLifetime struct {
				RawDay uint32 `env:"AUTH_JWT_REFRESH_TOKEN_LIFETIME_DAY,required"`
				Val    time.Duration
			}

			// Домены для рефреш токена. Список через пробел
			ServedDomains struct {
				Raw string `env:"AUTH_JWT_SERVED_DOMAINS,required"`
				Val []string
			}

			RefreshTokenCookieName string `env:"AUTH_JWT_REFRESH_TOKEN_COOKIE_NAME,required"`
		}
	}

	GrpcAuditClient struct {
		Host string `env:"GRPC_AUDIT_CLIENT_HOST,required"`
	}

	Media struct {
		Avatar struct {
			// Максимальный размер файла
			MaxSizeMB struct {
				RawMb uint8 `env:"MEDIA_AVATAR_MAX_SIZE_MB,required"`
				Bytes uint32
			}

			MaxSizeMB22 uint16 `env:"MEDIA_AVATAR_MAX_SIZE_MB,required"` // Максимальный размер файла
			// Минимальное разрешение исходника 50
			MinResolution uint16 `env:"MEDIA_AVATAR_MIN_RESOLUTION_PX,required"`
			// Разрешение картинки предварительного просмотра 50
			PreviewResolution uint16 `env:"MEDIA_AVATAR_PREVIEW_RESOLUTION,required"`
			// Разрешение картинки до которого сжать картинку для хранения 300
			Resolution uint16 `env:"MEDIA_AVATAR_RESOLUTION,required"`
		}

		Image struct {
			// Максимальный размер файла
			MaxSize struct {
				RawMb uint8 `env:"MEDIA_IMAGE_MAX_SIZE_MB,required"`
				Bytes uint32
			}
		}
	}

	S3Conn S3Conn

	Controller struct {
		MinTimeContextMs uint32 `env:"CONTROLLER_MIN_CONTEXT_MS,required"`
	}
}

func ParseAPI() (*Platform, error) {
	var cfg Platform

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	cfg.Auth.OnetimeEntryLifetime.Val = time.Minute * time.Duration(cfg.Auth.OnetimeEntryLifetime.Raw)
	cfg.Auth.JWT.AccessTokenLifetime.Val = time.Minute * time.Duration(cfg.Auth.JWT.AccessTokenLifetime.RawMin)
	cfg.Auth.JWT.RefreshTokenLifetime.Val = time.Hour * 24 * time.Duration(cfg.Auth.JWT.RefreshTokenLifetime.RawDay)
	cfg.Auth.JWT.ServedDomains.Val = strings.Fields(cfg.Auth.JWT.ServedDomains.Raw)

	cfg.Media.Avatar.MaxSizeMB.Bytes = uint32(cfg.Media.Avatar.MaxSizeMB.RawMb) * 1024 * 1024
	cfg.Media.Image.MaxSize.Bytes = uint32(cfg.Media.Image.MaxSize.RawMb) * 1024 * 1024

	return &cfg, nil
}

func (c *Platform) GetGRPCAuditConfig() *GrpcClient {
	return &GrpcClient{
		Host: c.GrpcAuditClient.Host,
	}
}
