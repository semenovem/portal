package throw

import (
	"errors"
	"fmt"
)

var (
	e = errors.New

	ErrOverNote           = e("more than one note value passed")
	ErrNoFile             = e("file not sent")
	ErrOverFile           = e("received more than one file")
	ErrFileTooBig         = e("file too big")
	ErrFileEmpty          = e("file empty")
	ErrUnsupportedContent = e("unsupported content type")
	ErrUnknownContentType = e("unknown content type")
)

func NewWithTargetErr(target, err error) error {
	return NewWithTargetErrf(target, err.Error())
}

func NewWithTargetErrf(target error, msg string, v ...any) error {
	if len(v) != 0 {
		msg = fmt.Sprintf(msg, v...)
	}

	switch t := target.(type) {
	case accessErr, *accessErr:
		return &accessErr{
			msg:    msg,
			target: t,
		}
	case authErr, *authErr:
		return &authErr{
			msg:    msg,
			target: t,
		}
	case badRequestErr, *badRequestErr:
		return &badRequestErr{
			msg:    msg,
			target: t,
		}

	case invalidErr, *invalidErr:
		return &invalidErr{
			msg:    msg,
			target: t,
		}
	}

	return fmt.Errorf("%s: %s", target.Error(), msg)
}
