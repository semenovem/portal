package it

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	// Размерность пароля
	maxUserEmailLen = 256 // Максимальная длина

	minUserPasswordLen = 6   // Минимальная длина
	maxUserPasswordLen = 128 // Максимальная длина

	minUserLoginLen = 6
	maxUserLoginLen = 64 // TODO синхронизировать с типом столбца хранения

	minUserNameLen = 2
	maxUserNameLen = 128
)

var (
	regValidateUserLogin = regexp.MustCompile(`(?i)^[a-zа-яйё]+[\wа-яйё_-]*[^_-]$`)
	regValidateUserName  = regexp.MustCompile(`^[a-zа-яёй0-9_-]*$`)
)

func ValidateUserLogin(login string) error {
	if l := utf8.RuneCountInString(login); l < minUserLoginLen {
		return ErrValidShort
	} else if l > maxUserLoginLen {
		return ErrValidLong
	}

	if !regValidateUserLogin.MatchString(login) {
		return ErrValidIllegalChar
	}

	return nil
}

// ValidateUserPassword одна цифра, заглавная, строчная буква, специальный символ и нет пробелов
func ValidateUserPassword(password string) error {
	if interValidator.Var(password, "ascii") != nil {
		return ErrValidIllegalChar
	}

	var (
		hasNum, hasLower, hasUpper, hasSpecial bool
	)

	if l := utf8.RuneCountInString(password); l < minUserPasswordLen {
		return ErrValidShort
	} else if l > maxUserPasswordLen {
		return ErrValidLong
	}

	for _, char := range password {
		switch {
		case unicode.IsNumber(char):
			hasNum = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsSymbol(char) || unicode.IsPunct(char):
			hasSpecial = true
		case unicode.IsSpace(char):
			return ErrValidIllegalChar
		}
	}

	v := hasNum && hasLower && hasUpper && hasSpecial
	if !v {
		return ErrValidPasswdWeak
	}

	return nil
}

func ValidateUserEmail(email string) error {
	if l := utf8.RuneCountInString(email); l > maxUserEmailLen {
		return ErrValidLong
	}

	if interValidator.Var(email, "email,ascii") != nil {
		return ErrValidIllegalChar
	}

	return nil
}

func ValidateUserName(name string) error {
	n := strings.TrimSpace(strings.ToLower(name))

	if l := utf8.RuneCountInString(n); l > maxUserNameLen {
		return ErrValidLong
	} else if l < minUserNameLen {
		return ErrValidShort
	}

	if !regValidateUserName.MatchString(n) {
		return ErrValidIllegalChar
	}

	return nil
}

func ValidateUserStatus(status string) error {
	_, err := ParseUserStatus(status)
	return err
}

func ValidateUserRole(role string) error {
	_, err := ParseUserRole(role)
	return err
}

func ValidateUserRoles(roles []string) error {
	_, err := ParseUserRoles(roles)
	return err
}
