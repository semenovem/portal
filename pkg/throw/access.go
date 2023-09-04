package throw

// AccessErr ошибки в результате нарушения доступа

type AccessErr interface {
	Error() string
	isAccessErr() bool
}

type accessErr struct {
	msg string
}

func (e accessErr) Error() string {
	return e.msg
}

func (e accessErr) isAccessErr() bool {
	return true
}

func NewAccessErr(msg string) error {
	return &accessErr{msg: msg}
}

func IsAccessErr(err error) bool {
	_, ok := err.(*accessErr)
	return ok
}
