package throw

var (
	Err400DuplicateLogin = NewBadRequestErr("duplicate user login")
	Err400DuplicateEmail = NewBadRequestErr("duplicate user email")
)

// BadRequestErr ошибки в результате проверок и валидации
type BadRequestErr interface {
	Error() string
	isBadRequestErr() bool
}

type badRequestErr struct {
	msg string
}

func (e badRequestErr) Error() string {
	return e.msg
}

func (e badRequestErr) isBadRequestErr() bool {
	return true
}

func NewBadRequestErr(msg string) error {
	return &badRequestErr{msg: msg}
}

func IsBadRequestErr(err error) bool {
	_, ok := err.(*badRequestErr)
	return ok
}
