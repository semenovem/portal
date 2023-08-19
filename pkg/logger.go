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
	Tags(tags ...string) Logger // Не создает новый объект, мутирует текущий
	DBTag() Logger
	RedisTag() Logger
	AuthTag() Logger
	ClientTag() Logger
}
