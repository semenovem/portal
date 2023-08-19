package config

type API struct {
	Base

	Rest       Rest
	DBCoreConn DatabaseCoreConn
	RedisConn  RedisConn

	DBMigrationsDir string `env:"DB_MIGRATIONS_DIR,required"`
	UserPasswdSalt  string `env:"USER_PASSWD_SALT,required"`

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
		OnetimeEntryLifetimeMin uint32 `env:"AUTH_ONETIME_ENTRY_LIFETIME_MIN,required"`
	}

	GrpcAuditClient struct {
		Host string `env:"GRPC_AUDIT_CLIENT_HOST,required"`
	}
}

func (c *API) GetGRPCAuditConfig() *GrpcClient {
	return &GrpcClient{
		Host: c.GrpcAuditClient.Host,
	}
}
