package throw

import "fmt"

/* InvalidErr ошибки в результате не соответствия типов, не валидных данных */

var (
	ErrInvalidZeroValue   = NewInvalidErr("zero value")         // Пустое значение
	ErrInvalidPasswdWeak  = NewInvalidErr("password is weak")   // Простой пароль
	ErrInvalidIllegalChar = NewInvalidErr("illegal characters") // Запрещенные символы
	ErrInvalidShort       = NewInvalidErr("short")              // Короткий
	ErrInvalidLong        = NewInvalidErr("long")               // Длинный
)

func NewInvalidErr(msg string) error {
	return &invalidErr{msg: msg}
}

// NewInvalidTimeFieldErr ошибка парсинга времени
func NewInvalidTimeFieldErr(fieldName string, err error) error {
	return &invalidErr{
		msg: fmt.Sprintf(
			"request field [%s] not time format (2001-03-24T00:00:00Z): %s",
			fieldName,
			err.Error(),
		),
	}
}

// NewInvalidTimeErr ошибка парсинга времени
func NewInvalidTimeErr(err error) error {
	return &invalidErr{
		msg: fmt.Sprintf("not time format (2001-03-24T00:00:00Z): %s", err.Error()),
	}
}

func IsInvalidErr(err error) bool {
	_, ok := err.(*invalidErr)
	return ok
}

type InvalidErr interface {
	Error() string
	isInvalidErr() bool
}

type invalidErr struct {
	msg    string
	target error
}

func (e invalidErr) Error() string {
	return e.msg
}

func (e invalidErr) Target() error {
	if e.target != nil {
		return e.target
	}

	return e
}

func (e invalidErr) isValidErr() bool {
	return true
}
