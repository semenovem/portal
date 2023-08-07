package config

type API struct {
	Base Base

	RestPort string `env:"REST_PORT,required"`

	DBCoreConn DatabaseCoreConn
	RedisConn  RedisConn

	DBMigrationsDir string `env:"DB_MIGRATIONS_DIR,required"`
}
