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
	ErrOverNote           = e("more than one note value passed")
	ErrNoFile             = e("file not sent")
	ErrOverFile           = e("more than one file sent")
	ErrFileTooBig         = e("file too big")
	ErrFileEmpty          = e("file empty")
	ErrUnsupportedContent = e("unsupported content type")

	ErrAccessTokenExp  = e("access token expired")
	ErrInvalidBearer   = e("invalid bearer token")
	ErrAuthCookieEmpty = e("empty header [Authorization] token")

	ErrUserLogouted = e("user is logouted")
)

var (
	ErrAccess = NewAccessErr("access denied")

	ErrNotFound = NewNotFoundErr("not found")
)

// AccessErr ошибки в результате нарушения доступа
// --------------------------------------------------------------
type AccessErr interface {
	Error() string
}

type accessErr struct {
	msg string
}

func (e accessErr) Error() string {
	return e.msg
}

func IsAccessErr(err error) bool {
	_, ok := err.(*accessErr)
	return ok
}

func NewAccessErr(msg string) error {
	return &accessErr{msg: msg}
}

// NotFoundErr ошибки в результате отсутствия запрошенной сущности
// --------------------------------------------------------------
type NotFoundErr interface {
	Error() string
}

type notFoundErr struct {
	msg string
}

func (e notFoundErr) Error() string {
	return e.msg
}

func NewNotFoundErr(msg string) NotFoundErr {
	return &notFoundErr{msg: msg}
}

// AuthErr ошибки в результате нарушения при авторизации
// --------------------------------------------------------------
type AuthErr interface {
	Error() string
}

type authErr struct {
	msg string
}

func (e authErr) Error() string {
	return e
}

func NewAuthErr(msg string) AuthErr {
	return &authErr{msg: msg}
}
