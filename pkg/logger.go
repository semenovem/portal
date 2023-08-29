package pkg

type Logger interface {
	Named(string) Logger
	With(key string, val interface{}) Logger // Не создает новый объект, мутирует текущий
	Error(msg string)
	Errorf(format string, v ...any)
	Info(msg string)
	Infof(format string, v ...any)
	Debug(msg string)
	Debugf(format string, v ...any)
	DebugOrErr(isErr bool, format string)
	DebugOrErrf(isErr bool, format string, v ...any)
	Nested(err error)
	NestedWith(err error, msg string) error
	Nestedf(format string, v ...any)

	/* методы тегирования не создают новый объект, а мутирует текущий */

	AuthTag() Logger

	DB(err error)
	DBStr(msg string)
	DBf(format string, v ...any)

	Redis(err error)
	RedisStr(msg string)
	Redisf(format string, v ...any)

	Client(err error)
	ClientStr(msg string)
	Clientf(format string, v ...any)

	BadRequest(err error)
	BadRequestStr(msg string)
	BadRequestStrRetErr(msg string) error

	NotFound(err error)
	NotFoundStr(msg string)

	Deny(err error)

	Auth(err error)
	AuthStr(msg string)
	AuthDebug(err error)
	AuthDebugStr(msg string)
}
