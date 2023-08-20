package it

import "errors"

var (
	ern = errors.New

	ErrValidZeroValue   = ern("invalid: zero value")                  // Пустое значение
	ErrValidPasswdWeak  = ern("invalid: password is weak")            // Простой пароль
	ErrValidIllegalChar = ern("invalid: contains illegal characters") // Запрещенные символы
	ErrValidShort       = ern("invalid: short")                       // Короткий
	ErrValidLong        = ern("invalid: long")                        // Длинный

	ErrUserExpired             = ern("user expired")
	ErrUserFired               = ern("user fired")
	ErrUserNotStartWork        = ern("user not start work")
	ErrUserHaveNotActiveStatus = ern("user have not active status")
)
