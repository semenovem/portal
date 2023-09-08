package throw

import "fmt"

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

func NewBadRequestTimeFieldErr(fieldName string) error {
	return &badRequestErr{
		msg: fmt.Sprintf("request field [%s] not time format (2001-03-24T00:00:00Z)", fieldName),
	}
}

func IsBadRequestErr(err error) bool {
	_, ok := err.(*badRequestErr)
	return ok
}
