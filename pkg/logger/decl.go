package logger

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	prefixErr   = "[ERRO] "
	prefixInfo  = "[INFO] "
	prefixDebug = "[DEBU] "
)

var (
	buf = new(bytes.Buffer)

	baseLoggerErr   = log.New(buf, prefixErr, 0)
	baseLoggerInfo  = log.New(buf, prefixInfo, 0)
	baseLoggerDebug = log.New(buf, prefixDebug, 0)

	loggerErr   = baseLoggerErr
	loggerInfo  = baseLoggerInfo
	loggerDebug = baseLoggerDebug

	loggerNoop = log.New(io.Discard, "", 0)

	cliMode     = false
	hasSetLevel = false
	hideTime    = false

	logLevel = -1

	outfile *os.File
	synced  = false
)

func SetLevel(level int) error {
	if level < -1 || level > 1 {
		return fmt.Errorf("invalid value of log level [%d]. Can be one of [-1|0|1]", level)
	}

	if level == logLevel {
		return nil
	}

	logLevel = level
	loggerDebug = baseLoggerDebug
	loggerInfo = baseLoggerInfo
	loggerErr = baseLoggerErr

	if level >= 0 {
		loggerDebug = loggerNoop

		if level != 0 {
			loggerInfo = loggerNoop
		}
	}
	hasSetLevel = true

	return nil
}

func GetLevel() int {
	return logLevel
}

func SetLevelIfNot(level int) {
	if hasSetLevel {
		return
	}

	SetLevel(level)
}

func SetCLIMode(mode bool) {
	if cliMode == mode {
		return
	}

	cliMode = mode

	if mode {
		loggerErr.SetPrefix("\033[0;31m" + prefixErr + "\033[0m")
		loggerInfo.SetPrefix("\033[0;32m" + prefixInfo + "\u001B[0m")
		loggerDebug.SetPrefix("\033[1;34m" + prefixDebug + "\u001B[0m")
	} else {
		loggerErr.SetPrefix(prefixErr)
		loggerInfo.SetPrefix(prefixInfo)
		loggerDebug.SetPrefix(prefixDebug)
	}
}

func GetCliMode() bool {
	return cliMode
}

func SetLogFile(logfile string) {
	defer Sync()

	f, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		Named("logger").Error("can't open file for logs: %v", err)
		return
	}

	outfile = f
}

func Sync() {
	if synced {
		return
	}

	synced = true

	if outfile == nil {
		_, _ = os.Stdout.Write(buf.Bytes())

		baseLoggerInfo.SetOutput(os.Stdout)
		baseLoggerDebug.SetOutput(os.Stdout)
		baseLoggerErr.SetOutput(os.Stdout)
	} else {
		_, _ = outfile.Write(buf.Bytes())

		loggerErr.SetOutput(outfile)
		loggerInfo.SetOutput(outfile)
		loggerDebug.SetOutput(outfile)
	}

}

func Close() {
	if outfile != nil {
		outfile.Close()
	}
}

func SetHideTime(v bool) {
	hideTime = v
}

func Error(format string, v ...any) {
	save(loggerErr, "", nil, format, v...)

}

func Info(format string, v ...any) {
	save(loggerInfo, "", nil, format, v...)
}

func InfoForce(format string, v ...any) {
	save(baseLoggerInfo, "", nil, format, v...)
}

func Debug(format string, v ...any) {
	save(loggerDebug, "", nil, format, v...)
}

func DebugForce(format string, v ...any) {
	save(baseLoggerDebug, "", nil, format, v...)
}

func Named(format string, v ...any) *Expander {
	return &Expander{name: fmt.Sprintf(format, v...)}
}

func With(k string, v interface{}) *Expander {
	p := &Expander{}
	p.addParam(k, v)
	return p
}
