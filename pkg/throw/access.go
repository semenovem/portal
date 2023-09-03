package throw

// AccessErr ошибки в результате нарушения доступа
type AccessErr struct {
	msg string
}

func (e AccessErr) Error() string {
	return e.msg
}

func NewAccessErr(msg string) error {
	return &AccessErr{msg: msg}
}

func IsAccessErr(err error) bool {
	_, ok := err.(*AccessErr)
	return ok
}
