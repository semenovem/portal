package config

import (
	"github.com/semenovem/portal/internal/zoo/conn"
	"strings"
	"time"
)

type Base struct {
	Env      string `env:"ENV,required"`
	CliMode  bool   `env:"CLI_MODE,required"`
	LogLevel int8   `env:"LOG_LEVEL,required"`
}

func (b *Base) IsDev() bool {
	return strings.EqualFold(b.Env, "DEV")
}

type RedisConn struct {
	Host     string `env:"REDIS_HOST,required"`
	Password string `env:"REDIS_PASSWORD,required"`
	DBName   uint16 `env:"REDIS_DB_NAME,required"`
}

func (c *RedisConn) ConvTo() *conn.RedisProps {
	return &conn.RedisProps{
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

func (c *DatabaseCoreConn) ConvTo() *conn.DBPGProps {
	return &conn.DBPGProps{
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

type Rest struct {
	Port              uint16 `env:"REST_PORT,required"`
	CorsAllowedHosts  string `env:"REST_CORS_ALLOWED_HOSTS,required"`
	CorsMaxAgeSeconds uint32 `env:"REST_CORS_MAX_AGE_SECONDS,required"`
}

func (r *Rest) GetCorsMaxAge() time.Duration {
	return time.Second * time.Duration(r.CorsMaxAgeSeconds)
}

type GrpcServer struct {
	Port                 string `env:"GRPC_SERVER_PORT,required"`
	MaxProcessingTimeSec uint32 `env:"GRPC_SERVER_MAX_PROCESSING_TIME_SEC" envDefault:"5"`
}

func (s *GrpcServer) GetMaxProcessingTimeSec() time.Duration {
	return time.Second * time.Duration(s.MaxProcessingTimeSec)
}

type GrpcClient struct {
	Host string
}

type S3Conn struct {
	URL                string `env:"S3_URL,required"`
	AccessKey          string `env:"S3_ACCESS_KEY,required"`
	SecretKey          string `env:"S3_SECRET_KEY,required"`
	UseSSL             bool   `env:"S3_USE_SSL,required"`
	InsecureSkipVerify bool   `env:"S3_INSECURE_SKIP_VERIFY,required"`
}

type FileUpload struct {
	AvatarMaxMB uint16 `env:"UPLOAD_AVATAR_MAX_MB,required"` // Максимальный размер файла аватарки
	ImageMaxMB  uint16 `env:"UPLOAD_IMAGE_MAX_MB,required"`  // Максимальный размер файла фото
	VideoMaxMB  uint16 `env:"UPLOAD_VIDEO_MAX_MB,required"`  // Максимальный размер файла видео
	DocMaxMB    uint16 `env:"UPLOAD_DOC_MAX_MB,required"`    // Максимальный размер файла документа
}
