package entity

import (
	"unicode"
	"unicode/utf8"
)

func ValidateUserLogin(login string) error {

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

func ValidateUserEmail(login string) error {

	return nil
}
