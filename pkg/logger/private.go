package logger

import (
	"fmt"
	"io"
	"strings"
	"time"
)

func (p *Pen) save(level Level, format string, v ...any) {
	var b = make([]byte, prefixLen)

	if !p.hideTime {
		b = append(b, []byte(fmt.Sprintf(" [%s]", time.Now().Format(time.RFC3339)))...)
	}

	if p.names != nil {
		b = append(b, []byte(" "+strings.Join(p.names, "."))...)
		if p.isNested {
			b = append(b, nestedBytes...)
		}
	}

	if p.tags != nil {
		b = append(b, []byte(" ["+strings.Join(p.tags, ".")+"]")...)
		if p.isNested {
			b = append(b, nestedBytes...)
		}
	}

	if p.params != nil {
		b = append(b, []byte(" ["+strings.Join(p.params, ", ")+"]")...)
	}

	b = append(b, []byte(" "+fmt.Sprintf(format, v...))...)

	b = append(b, []byte("\n")...)

	var (
		writer io.Writer
		prefix []byte
	)

	switch level {
	case Error:
		prefix = prefixBytesErr
		writer = p.outErr
	case Info:
		prefix = prefixBytesInfo
		writer = p.outInf
	case Debug:
		prefix = prefixBytesDebug
		writer = p.outDeb
	}

	_ = copy(b[0:prefixLen], prefix)

	if _, err := writer.Write(b); err != nil {
		fmt.Println(prefixErr, string(b))
	}
}

func (p *Pen) copy() *Pen {
	a := Pen{
		level:    p.level,
		isNested: p.isNested,
		hideTime: p.hideTime,
		cli:      p.cli,
		outErr:   p.outErr,
		outInf:   p.outInf,
		outDeb:   p.outDeb,
	}

	if p.names != nil {
		a.names = make([]string, len(p.names))
		copy(a.names, p.names)
	}

	if p.params != nil {
		a.params = make([]string, len(p.params))
		copy(a.params, p.params)
	}

	if p.tags != nil {
		a.tags = make([]string, len(p.tags))
		copy(a.tags, p.tags)
	}

	return &a
}
