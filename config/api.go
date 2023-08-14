package config

type API struct {
	Base

	RestPort string `env:"REST_PORT,required"`

	DBCoreConn DatabaseCoreConn
	RedisConn  RedisConn

	DBMigrationsDir string `env:"DB_MIGRATIONS_DIR,required"`
	PasswdSalt      string `env:"PASSWD_SALT,required"`

	JWT struct {
		AccessTokenSecret       string `env:"JWT_ACCESS_TOKEN_SECRET,required"`
		RefreshTokenSecret      string `env:"JWT_REFRESH_TOKEN_SECRET,required"`
		AccessTokenLifetimeMin  uint32 `env:"JWT_ACCESS_TOKEN_LIFETIME_MIN,required"`
		RefreshTokenLifetimeDay uint32 `env:"JWT_REFRESH_TOKEN_LIFETIME_DAY,required"`

		// Домены для рефреш токена. Список через запятую
		ServedDomains          string `env:"JWT_SERVED_DOMAINS,required"`
		RefreshTokenCookieName string `env:"JWT_REFRESH_TOKEN_COOKIE_NAME,required"`
	}
}
