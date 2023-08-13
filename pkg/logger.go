package pkg

type Logger interface {
	Named(string) Logger
	With(key string, val interface{}) Logger
	Error(msg string)
	Errorf(format string, v ...any)
	Info(msg string)
	Infof(format string, v ...any)
	Debug(msg string)
	Debugf(format string, v ...any)
	Nested(format string)
	Nestedf(format string, v ...any)
	Tags(tags ...string) Logger
	DBTag() Logger
	AuthTag() Logger
}
