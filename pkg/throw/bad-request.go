package throw

var (
	ErrBadRequestDuplicateLogin = NewBadRequestErr("duplicate user login")
	ErrBadRequestDuplicateEmail = NewBadRequestErr("duplicate user email")
)

// BadRequestErr ошибки в результате проверок и валидации
type BadRequestErr struct {
	msg string
}

func (e BadRequestErr) Error() string {
	return e.msg
}

func (e BadRequestErr) IsBadRequestErr() bool {
	return true
}

func NewBadRequestErr(msg string) error {
	return &BadRequestErr{msg: msg}
}
