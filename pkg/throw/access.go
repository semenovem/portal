package throw

/* AccessErr ошибки в результате нарушения доступа */

type AccessErr interface {
	Error() string
	isAccessErr() bool
}

type accessErr struct {
	msg    string
	target error
}

func NewAccessErr(msg string) error {
	return &accessErr{msg: msg}
}

func IsAccessErr(err error) bool {
	_, ok := err.(*accessErr)
	return ok
}

func (e accessErr) Error() string {
	if e.target != nil {
		if e.msg == "" {
			return e.target.Error()
		}
		return e.target.Error() + ": " + e.msg
	}

	return e.msg
}

func (e accessErr) Target() error {
	if e.target != nil {
		return e.target
	}

	return e
}

func (e accessErr) isAccessErr() bool {
	return true
}
