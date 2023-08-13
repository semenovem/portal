package logger

import (
	"fmt"
	"github.com/semenovem/portal/pkg"
	"io"
	"os"
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
	tags     []string
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

	a.params = append(a.params, fmt.Sprintf("%s:%+v", k, v))

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

func (p *Pen) Tags(tags ...string) pkg.Logger {
	if p.tags == nil {
		p.tags = make([]string, 0)
	}

	p.tags = append(p.tags, tags...)

	return p
}

func (p *Pen) DBTag() pkg.Logger {
	return p.Tags(DatabaseTag)
}

func (p *Pen) AuthTag() pkg.Logger {
	return p.Tags(AuthTag)
}
