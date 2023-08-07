package logger

import (
    "fmt"
    "github.com/semenovem/portal/pkg"
    "io"
    "os"
    "strings"
    "time"
)

type Pen struct {
    names    []string
    params   []string
    level    int8
    isNested bool
    hideTime bool
    cli      bool
    outErr   io.Writer
    outInf   io.Writer
    outDeb   io.Writer
}

func New() (*Pen, *Setter) {
    st := Setter{
        logger: &Pen{
            outErr: os.Stdout,
            outInf: os.Stdout,
            outDeb: os.Stdout,
        },
    }

    return st.logger, &st
}

func (p *Pen) Named(n string) pkg.Logger {
    a := p.copy()

    if a.names == nil {
        a.names = make([]string, 0)
    }

    a.names = append(a.names, n)

    return a
}

func (p *Pen) With(k string, v interface{}) pkg.Logger {
    a := p.copy()

    if a.params == nil {
        a.params = make([]string, 0)
    }

    a.params = append(a.params, fmt.Sprintf("%s:%v", k, v))

    return a
}

func (p *Pen) Error(format string) {
    p.save(Error, format)
}

func (p *Pen) Errorf(format string, v ...any) {
    p.save(Error, format, v...)
}

func (p *Pen) Info(format string) {
    p.save(Info, format)
}

func (p *Pen) Infof(format string, v ...any) {
    p.save(Info, format, v...)
}

func (p *Pen) Debug(format string) {
    p.save(Debug, format)
}

func (p *Pen) Debugf(format string, v ...any) {
    p.save(Debug, format, v...)
}

func (p *Pen) Nested(format string) {
    p.isNested = true
    p.save(Debug, format)
}

func (p *Pen) Nestedf(format string, v ...any) {
    p.isNested = true
    p.save(Debug, format, v...)
}

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

    return &a
}
