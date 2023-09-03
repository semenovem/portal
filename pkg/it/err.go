package it

import "errors"

var (
	e = errors.New

	ErrValidZeroValue   = e("invalid: zero value")                  // Пустое значение
	ErrValidPasswdWeak  = e("invalid: password is weak")            // Простой пароль
	ErrValidIllegalChar = e("invalid: contains illegal characters") // Запрещенные символы
	ErrValidShort       = e("invalid: short")                       // Короткий
	ErrValidLong        = e("invalid: long")                        // Длинный

	ErrUserExpired             = e("user expired")
	ErrUserFired               = e("user fired")
	ErrUserNotStartWork        = e("user not start work")
	ErrUserHaveNotActiveStatus = e("user have not active status")
	msgErrUserStatusInvalid    = "user status [%s] invalid"
	msgErrUserRoleInvalid      = "user role [%s] invalid"
)
