package pkg

import "context"

type Logger interface {
	Func(context.Context, string) Logger
	Named(string) Logger
	With(key string, val interface{}) Logger // Не создает новый объект, мутирует текущий
	Error(msg string)
	ErrorE(err error)
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

	DB(err error)
	DBStr(msg string)
	DBf(format string, v ...any)

	Redis(err error)
	RedisStr(msg string)
	Redisf(format string, v ...any)

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
