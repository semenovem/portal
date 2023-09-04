package it

import "errors"

var (
	e = errors.New

	ErrValidZeroValue   = NewValidateErr("zero value")         // Пустое значение
	ErrValidPasswdWeak  = NewValidateErr("password is weak")   // Простой пароль
	ErrValidIllegalChar = NewValidateErr("illegal characters") // Запрещенные символы
	ErrValidShort       = NewValidateErr("short")              // Короткий
	ErrValidLong        = NewValidateErr("long")               // Длинный

	ErrUserExpired      = e("user expired")
	ErrUserFired        = e("user fired")
	ErrUserNotStartWork = e("user not start work")
	ErrUserNotActive    = e("user have not active status")

	msgErrUnknownUserStatus = "unknown user status [%s]"
	msgErrUnknownUserRole   = "unknown user role [%s]"
)

type ValidateErr interface {
	Error() string
	isValidateErr() bool
}

type validateErr struct {
	msg string
}

func (e validateErr) Error() string {
	return e.msg
}

func (e validateErr) isAccessErr() bool {
	return true
}

func NewValidateErr(msg string) error {
	return &validateErr{msg: msg}
}

func IsValidateErr(err error) bool {
	_, ok := err.(*validateErr)
	return ok
}
