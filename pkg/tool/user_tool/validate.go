package user_tool

import (
	"github.com/semenovem/portal/pkg/tool/user_tool/user_err"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
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
	regValidatePassword  = regexp.MustCompile(`^[a-zA-Zа-яёйА-ЯЁЙ0-9_!@#$%^&*()=+"№:,\[\]\\{}-]*$`)
)

func ValidateStatus(status string) error {
	_, ok := ParseStatus(status)
	if !ok {
		return user_err.UnknownStatus.Err()
	}
	return nil
}

func ValidateGroup(group string) error {
	_, ok := ParseGroup(group)
	if !ok {
		return user_err.UnknownGroup.Err()
	}
	return nil
}

func ValidateUserGroups(groups []string) error {
	_, ok := ParseUserGroups(groups)
	if !ok {
		return user_err.UnknownGroup.Err()
	}
	return nil
}

func ValidateUserName(name string) error {
	n := strings.TrimSpace(strings.ToLower(name))

	if l := utf8.RuneCountInString(n); l > maxUserNameLen {
		return user_err.NameTooLong.Err()
	} else if l < minUserNameLen {
		return user_err.NameTooShort.Err()
	}

	if !regValidateUserName.MatchString(n) {
		return user_err.NameContainsIllegalChars.Err()
	}

	return nil
}

func ValidateUserLogin(login string) error {
	if l := utf8.RuneCountInString(login); l < minUserLoginLen {
		return user_err.LoginTooShort.Err()
	} else if l > maxUserLoginLen {
		return user_err.LoginTooLong.Err()
	}

	if !regValidateUserLogin.MatchString(login) {
		return user_err.LoginContainsIllegalChars.Err()
	}

	return nil
}

// ValidateUserPassword одна цифра, заглавная, строчная буква, специальный символ и нет пробелов
func ValidateUserPassword(password string) []error {
	if !regValidatePassword.MatchString(password) {
		return []error{user_err.PasswdContainsIllegalChars.Err()}
	}

	var (
		hasNum, hasLower, hasUpper, hasSpecial bool
		errs                                   []error
	)

	if l := utf8.RuneCountInString(password); l < minUserPasswordLen {
		errs = append(errs, user_err.PasswdTooShort.Err())
	} else if l > maxUserPasswordLen {
		errs = append(errs, user_err.PasswdTooLong.Err())
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
		}
	}

	if !hasNum {
		errs = append(errs, user_err.PasswdNotContainDigit.Err())
	}

	if !hasLower {
		errs = append(errs, user_err.PasswdNotContainLower.Err())
	}

	if !hasUpper {
		errs = append(errs, user_err.PasswdNotContainUpper.Err())
	}

	if !hasSpecial {
		errs = append(errs, user_err.PasswdNotContainSpecialSymbol.Err())
	}

	return nil
}
