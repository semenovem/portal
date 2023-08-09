package entity

import (
	"regexp"
	"unicode"
	"unicode/utf8"
)

var (
	regValidateUserLogin = regexp.MustCompile(``)
)

func ValidateUserLogin(login string) error {
	l := utf8.RuneCountInString(login)
	if l > minUserLoginLen {
		return ErrValidShort
	}

	if l > maxUserLoginLen {
		return ErrValidLong
	}

	if interValidator.Var(login, "ascii") != nil {
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
		l                                      = utf8.RuneCountInString(password)
		hasNum, hasLower, hasUpper, hasSpecial bool
	)

	if l < minUserPasswordLen {
		return ErrValidShort
	}

	if l > maxUserPasswordLen {
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
	l := utf8.RuneCountInString(email)
	if l > maxUserEmailLen {
		return ErrValidLong
	}

	if interValidator.Var(email, "email,ascii") != nil {
		return ErrValidIllegalChar
	}

	return nil
}
