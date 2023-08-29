package logger

import (
	"fmt"
	"github.com/pkg/errors"
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
	a.with(k, fmt.Sprintf("%+v", v))

	return a
}

func (p *pen) Error(format string) {
	p.save(Error, format)
}

func (p *pen) ErrorE(err error) {
	p.save(Error, err.Error())
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

func (p *pen) Nested(err error) {
	p.isNested = true
	p.save(Debug, err.Error())
}

func (p *pen) NestedWith(err error, msg string) error {
	p.isNested = true

	if err == nil && msg == "" {
		p.save(Error, "logger argument is Nil (cause - developer)")
		return nil
	}

	if msg == "" {
		p.save(Debug, err.Error())
	} else {
		if err == nil {
			err = errors.New(msg)
		} else {
			p.with(nesterErrMsg, " "+err.Error())
			err = errors.WithMessage(err, msg)
		}

		p.save(Debug, msg)
	}

	return err
}

func (p *pen) Nestedf(format string, v ...any) {
	p.isNested = true
	p.save(Debug, format, v...)
}

func (p *pen) addTag(tags ...string) *pen {
	if p.tags == nil {
		p.tags = make([]string, 0)
	}

	p.tags = append(p.tags, tags...)

	return p
}

// ----------------------------------------

func (p *pen) DB(err error) {
	p.copy().addTag(databaseTag).save(Error, err.Error())
}

func (p *pen) DBStr(msg string) {
	p.copy().addTag(databaseTag).save(Error, msg)
}

func (p *pen) DBf(format string, v ...any) {
	p.copy().addTag(databaseTag).save(Error, format, v...)
}

func (p *pen) Redis(err error) {
	p.copy().addTag(redisTag).save(Error, err.Error())
}

func (p *pen) RedisStr(msg string) {
	p.copy().addTag(redisTag).save(Error, msg)
}

func (p *pen) Redisf(format string, v ...any) {
	p.copy().addTag(redisTag).save(Error, format, v...)
}

func (p *pen) BadRequest(err error) {
	p.copy().addTag(badRequestTag).save(Debug, err.Error())
}

func (p *pen) BadRequestStr(msg string) {
	p.copy().addTag(badRequestTag).save(Debug, msg)
}

func (p *pen) BadRequestStrRetErr(msg string) error {
	p.copy().addTag(badRequestTag).BadRequestStr(msg)
	return errors.New(msg)
}

func (p *pen) NotFound(err error) {
	p.copy().addTag(notFound).save(Info, err.Error())
}

func (p *pen) NotFoundStr(msg string) {
	p.copy().addTag(notFound).save(Info, msg)
}

func (p *pen) Deny(err error) {
	p.copy().addTag(denyTag).save(Info, err.Error())
}

func (p *pen) Auth(err error) {
	p.copy().addTag(authTag).save(Info, err.Error())
}

func (p *pen) AuthStr(msg string) {
	p.copy().addTag(authTag).save(Info, msg)
}

func (p *pen) AuthDebug(err error) {
	p.copy().addTag(authTag).save(Debug, err.Error())
}

func (p *pen) AuthDebugStr(msg string) {
	p.copy().addTag(authTag).save(Debug, msg)
}
