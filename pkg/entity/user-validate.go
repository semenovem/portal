package entity

import (
	"regexp"
	"unicode"
	"unicode/utf8"
)

var (
	regValidateUserLogin = regexp.MustCompile(`(?i)^[a-z]+[\w_-]*[^_-]$`)
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
	l := utf8.RuneCountInString(email)
	if l > maxUserEmailLen {
		return ErrValidLong
	}

	if interValidator.Var(email, "email,ascii") != nil {
		return ErrValidIllegalChar
	}

	return nil
}
