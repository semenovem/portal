package logger

import (
	"fmt"
	"github.com/semenovem/portal/pkg"
	"io"
	"os"
)

type pen struct {
	names    []string
	params   []string
	level    int8
	isNested bool
	hideTime bool
	cli      bool
	outErr   io.Writer
	outInf   io.Writer
	outDeb   io.Writer
	tags     []string
}

func New() (pkg.Logger, *Setter) {
	st := Setter{
		logger: &pen{
			outErr: os.Stdout,
			outInf: os.Stdout,
			outDeb: os.Stdout,
		},
	}

	return st.logger, &st
}

func (p *pen) Named(n string) pkg.Logger {
	a := p.copy()

	if a.names == nil {
		a.names = make([]string, 0)
	}

	a.names = append(a.names, n)

	return a
}

func (p *pen) With(k string, v interface{}) pkg.Logger {
	a := p.copy()

	if a.params == nil {
		a.params = make([]string, 0)
	}

	a.params = append(a.params, fmt.Sprintf("%s:%+v", k, v))

	return a
}

func (p *pen) Error(format string) {
	p.save(Error, format)
}

func (p *pen) Errorf(format string, v ...any) {
	p.save(Error, format, v...)
}

func (p *pen) Info(format string) {
	p.save(Info, format)
}

func (p *pen) Infof(format string, v ...any) {
	p.save(Info, format, v...)
}

func (p *pen) Debug(format string) {
	p.save(Debug, format)
}

func (p *pen) Debugf(format string, v ...any) {
	p.save(Debug, format, v...)
}

func (p *pen) DebugOrErr(isDebug bool, format string) {
	if isDebug {
		p.save(Debug, format)
	} else {
		p.save(Error, format)
	}
}

func (p *pen) DebugOrErrf(isDebug bool, format string, v ...any) {
	if isDebug {
		p.save(Debug, format, v...)
	} else {
		p.save(Error, format, v...)
	}
}

func (p *pen) Nested(format string) {
	p.isNested = true
	p.save(Debug, format)
}

func (p *pen) Nestedf(format string, v ...any) {
	p.isNested = true
	p.save(Debug, format, v...)
}

func (p *pen) Tags(tags ...string) pkg.Logger {
	if p.tags == nil {
		p.tags = make([]string, 0)
	}

	p.tags = append(p.tags, tags...)

	return p
}

func (p *pen) DBTag() pkg.Logger {
	return p.Tags(DatabaseTag)
}

func (p *pen) RedisTag() pkg.Logger {
	return p.Tags(RedisTag)
}

func (p *pen) AuthTag() pkg.Logger {
	return p.Tags(AuthTag)
}

func (p *pen) ClientTag() pkg.Logger {
	return p.Tags(ClientTag)
}
