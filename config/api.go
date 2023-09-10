package config

type API struct {
	Base

	Rest       Rest
	DBCoreConn DatabaseCoreConn
	RedisConn  RedisConn

	DBMigrationsDir    string `env:"DB_MIGRATIONS_DIR,required"`
	DBMigrationsDirDev string `env:"DB_MIGRATIONS_DIR_DEV,required"` // Локальные данные для разработки
	UserPasswdSalt     string `env:"USER_PASSWD_SALT,required"`

	JWT struct {
		AccessTokenSecret       string `env:"JWT_ACCESS_TOKEN_SECRET,required"`
		RefreshTokenSecret      string `env:"JWT_REFRESH_TOKEN_SECRET,required"`
		AccessTokenLifetimeMin  uint32 `env:"JWT_ACCESS_TOKEN_LIFETIME_MIN,required"`
		RefreshTokenLifetimeDay uint32 `env:"JWT_REFRESH_TOKEN_LIFETIME_DAY,required"`

		// Домены для рефреш токена. Список через запятую
		ServedDomains          string `env:"JWT_SERVED_DOMAINS,required"`
		RefreshTokenCookieName string `env:"JWT_REFRESH_TOKEN_COOKIE_NAME,required"`
	}

	Auth struct {
		// Длительность сессии одноразовой ссылки
		OnetimeEntryLifetimeMin uint32 `env:"AUTH_ONETIME_ENTRY_LIFETIME_MIN,required"`
	}

	GrpcAuditClient struct {
		Host string `env:"GRPC_AUDIT_CLIENT_HOST,required"`
	}

	Upload struct {
		ImageMaxMB uint16 `env:"UPLOAD_IMAGE_MAX_MB,required"` // Максимальный размер файла фото
		VideoMaxMB uint16 `env:"UPLOAD_VIDEO_MAX_MB,required"` // Максимальный размер файла видео
		DocMaxMB   uint16 `env:"UPLOAD_DOC_MAX_MB,required"`   // Максимальный размер файла документа
	}

	S3Conn S3Conn
}

func (c *API) GetGRPCAuditConfig() *GrpcClient {
	return &GrpcClient{
		Host: c.GrpcAuditClient.Host,
	}
}
