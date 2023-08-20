package logger

const (
	prefixErr   = "[ERRO]"
	prefixInfo  = "[INFO]"
	prefixDebug = "[DEBU]"
)

var (
	prefixBytesErr   = []byte(prefixErr)
	prefixBytesInfo  = []byte(prefixInfo)
	prefixBytesDebug = []byte(prefixDebug)

	nestedBytes = []byte(".NESTED")

	prefixLen = len(prefixBytesErr)
)

const (
	Debug Level = iota - 1
	Info
	Error
)

type Level int8

type Tag string

const (
	databaseTag = "DATABASE"
	redisTag    = "REDIS"
	authTag     = "AUTH"
	clientTag   = "CLIENT"
	denyTag     = "DENY"
	notFound    = "NOT_FOUND"
)
