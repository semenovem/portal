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
	Nested(format string)
	Nestedf(format string, v ...any)
	DBTag() Logger // Все методы тегирования не создают новый объект, а мутирует текущий
	RedisTag() Logger
	AuthTag() Logger
	ClientTag() Logger
	DenyTag() Logger
	NotFoundTag() Logger

	DB(err error)
	DBStr(msg string)
	DBf(format string, v ...any)
}
