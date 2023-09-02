package throw

// BadRequestErr ошибки в результате проверок и валидации
type BadRequestErr interface {
	Error() string
}

type badRequestErr struct {
	msg string
}

func (e badRequestErr) Error() string {
	return e.msg
}

func NewBadRequestErr(msg string) error {
	return &badRequestErr{msg: msg}
}

func IsBadRequestErr(err error) bool {
	_, ok := err.(*badRequestErr)
	return ok
}
