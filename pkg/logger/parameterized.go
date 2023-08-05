package logger

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Expander struct {
	name   string
	params []*param
}

func New() *Expander {
	return &Expander{}
}

type param struct {
	key, value string
}

func (l *Expander) Named(n string) *Expander {
	l2 := l.copy()

	if l2.name != "" {
		l2.name += "."
	}
	l2.name += strings.TrimSpace(n)

	return l2
}

func (l *Expander) With(k string, v interface{}) *Expander {
	l2 := l.copy()

	l2.addParam(k, v)

	return l2
}

func (l *Expander) Error(format string, v ...any) {
	save(loggerErr, l.name, l.params, format, v...)
}

func (l *Expander) Info(format string, v ...any) {
	save(loggerInfo, l.name, l.params, format, v...)
}

func (l *Expander) Debug(format string, v ...any) {
	save(loggerDebug, l.name, l.params, format, v...)
}

func (l *Expander) Nested(format string, v ...any) {
	var name string
	if l.name != "" {
		name = l.name + ".NESTED"
	} else {
		name = "NESTED"
	}
	save(loggerDebug, name, l.params, format, v...)
}

func save(lg *log.Logger, name string, params []*param, format string, v ...any) {
	var with string

	if name != "" {
		name = name + ": "
	}

	if len(params) != 0 {
		sep := ""
		for _, p := range params {
			if with != "" && sep == "" {
				sep = ","
			}
			with += fmt.Sprintf("%s%s:%v", sep, p.key, p.value)
		}
		with = "[" + with + "] "
	}

	if hideTime {
		lg.Printf("%s%s%s", name, with, fmt.Sprintf(format, v...))
	} else {
		lg.Printf("[%s] %s%s%s", time.Now().Format(time.RFC3339), name, with, fmt.Sprintf(format, v...))
	}
}

func (l *Expander) copy() *Expander {
	var w []*param

	if len(l.params) != 0 {
		w = make([]*param, len(l.params))
		for i, v := range l.params {
			w[i] = new(param)
			*w[i] = *v
		}
	}

	return &Expander{
		name:   l.name,
		params: w,
	}
}

func (l *Expander) addParam(k string, v interface{}) {
	if l.params == nil {
		l.params = make([]*param, 0)
	}
	l.params = append(l.params, &param{
		key:   k,
		value: fmt.Sprintf("%+v", v),
	})
}
