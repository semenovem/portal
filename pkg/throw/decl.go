package throw

import (
	"errors"
	"fmt"
)

var (
	e = errors.New

	ErrOverNote           = e("more than one note value passed")
	ErrNoFile             = e("file not sent")
	ErrOverFile           = e("more than one file sent")
	ErrFileTooBig         = e("file too big")
	ErrFileEmpty          = e("file empty")
	ErrUnsupportedContent = e("unsupported content type")
)

func NewWithTargetErr(target, err error) error {
	return NewWithTargetErrf(err, err.Error())
}

func NewWithTargetErrf(target error, msg string) error {
	switch t := target.(type) {
	case accessErr:
		return &accessErr{
			msg:    msg,
			target: t,
		}
	case authErr:
		return &authErr{
			msg:    msg,
			target: t,
		}
	case badRequestErr:
		return &badRequestErr{
			msg:    msg,
			target: t,
		}

	case invalidErr:
		return &invalidErr{
			msg:    msg,
			target: t,
		}
	}

	return fmt.Errorf("%s: %s", target.Error(), msg)
}
