package people

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"regexp"
)

var (
	ErrUserExpired      = errors.New("user expired")
	ErrUserFired        = errors.New("user fired")
	ErrUserNotStartWork = errors.New("user not start work")
	ErrUserNotActive    = errors.New("user have not active status")
)

const (
	// Размерность пароля
	maxUserEmailLen = 256 // Максимальная длина

	minUserPasswordLen = 6   // Минимальная длина
	maxUserPasswordLen = 128 // Максимальная длина

	minUserLoginLen = 3
	maxUserLoginLen = 64 // TODO синхронизировать с типом столбца хранения

	minUserNameLen = 2
	maxUserNameLen = 128
)

var (
	regValidateUserLogin = regexp.MustCompile(`(?i)^[a-zа-яйё]+[\wа-яйё_-]*[^_-]$`)
	regValidateUserName  = regexp.MustCompile(`^[a-zа-яёй0-9_-]*$`)

	interValidator = validator.New()
)
