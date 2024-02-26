package user_err

import (
	"errors"
	"strings"
)

type Err string

func (e Err) S() string {
	return string(e)
}

func (e Err) String() string {
	return string(e)
}

func (e Err) Err() error {
	if msg, ok := g[e]; ok {
		return errors.New(msg)
	}
	return errors.New(strings.ToLower(e.String()))
}

const (
	UnknownStatus Err = "UNKNOWN_USER_STATUS"
	UnknownGroup  Err = "UNKNOWN_USER_GROUP"

	Expired      Err = "USER_EXPIRED"
	Fired        Err = "USER_FIRED"
	NotStartWork Err = "USER_NOT_START_WORK"
	NotActive    Err = "USER_NOT_ACTIVE"

	NameContainsIllegalChars Err = "USER_NAME_CONTAINS_ILLEGAL_CHARS"
	NameTooLong              Err = "USER_NAME_TOO_LONG"
	NameTooShort             Err = "USER_NAME_TOO_SHORT"

	LoginContainsIllegalChars Err = "USER_LOGIN_CONTAINS_ILLEGAL_CHARS"
	LoginTooLong              Err = "USER_LOGIN_TOO_LONG"
	LoginTooShort             Err = "USER_LOGIN_TOO_SHORT"
)

var g = map[Err]string{
	LoginTooShort:             "user login too short",
	LoginTooLong:              "user login too long",
	LoginContainsIllegalChars: "user login contains illegal chars",

	NameTooShort:             "user name too short",
	NameTooLong:              "user name too long",
	NameContainsIllegalChars: "user name contains illegal chars",

	NotActive:     "user have not active status",
	NotStartWork:  "user not start work",
	Fired:         "user fired",
	Expired:       "user expired",
	UnknownGroup:  "unknown user group",
	UnknownStatus: "unknown user status",
}
