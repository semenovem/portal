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
)

var (
	ErrOverNote   = e("more than one note value passed")
	ErrNoFile     = e("file not sent")
	ErrOverFile   = e("more than one file sent")
	ErrFileTooBig = e("file too big")
	ErrFileEmpty  = e("file empty")

	ErrAccessTokenExp  = e("access token expired")
	ErrInvalidBearer   = e("invalid bearer token")
	ErrAuthCookieEmpty = e("empty header [Authorization] token")

	ErrUserLogouted = e("user is logouted")
)
