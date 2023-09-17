package people

import (
	"fmt"
	"github.com/semenovem/portal/pkg/throw"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

var (
	ErrUnknownUserStatus = throw.NewInvalidErr("unknown user status")
	ErrUnknownUserGroup  = throw.NewInvalidErr("unknown user group")
)

type UserGroup string
type UserStatus string

const (
	UserGroupSuperAdmin UserGroup = "super-admin"
	UserGroupAdmin      UserGroup = "admin"
	UserGroupUser       UserGroup = "user"
)

const (
	UserStatusInactive UserStatus = "inactive"
	UserStatusActive   UserStatus = "active"
	UserStatusBlocked  UserStatus = "blocked"
)

func ParseUserStatus(s string) (UserStatus, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case string(UserStatusInactive):
		return UserStatusInactive, nil
	case string(UserStatusActive):
		return UserStatusActive, nil
	case string(UserStatusBlocked):
		return UserStatusBlocked, nil
	}

	return "", throw.NewWithTargetErrf(ErrUnknownUserStatus, "origin:[%s]", s)
}

func ParseUserStatusIfDefault(s string) UserStatus {
	st, err := ParseUserStatus(s)
	if err != nil {
		return UserStatusInactive
	}
	return st
}

func ParseUserGroup(s string) (UserGroup, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case string(UserGroupSuperAdmin):
		return UserGroupSuperAdmin, nil
	case string(UserGroupAdmin):
		return UserGroupAdmin, nil
	case string(UserGroupUser):
		return UserGroupUser, nil
	}

	return "", throw.NewWithTargetErrf(ErrUnknownUserGroup, "origin:[%s]", s)
}

func ParseUserGroups(groups []string) ([]UserGroup, error) {
	var (
		errs   = make([]string, 0)
		result = make([]UserGroup, 0)
	)

	for _, group := range groups {
		if r, err := ParseUserGroup(group); err != nil {
			errs = append(errs, err.Error())
		} else {
			result = append(result, r)
		}
	}

	if len(errs) != 0 {
		return nil, fmt.Errorf(strings.Join(errs, "; "))
	}

	return result, nil
}

type UserAuth struct {
	ID        uint32
	Status    UserStatus
	ExpiredAt *time.Time
	WorkedAt  *time.Time
	FiredAt   *time.Time
}

func (u *UserAuth) CanLogging() error {
	if u.Status != UserStatusActive {
		return ErrUserNotActive
	}
	now := time.Now()

	if u.ExpiredAt != nil && u.ExpiredAt.Before(now) {
		return ErrUserExpired
	}

	if u.WorkedAt != nil && u.WorkedAt.After(now) {
		return ErrUserNotStartWork
	}

	if u.FiredAt != nil && u.FiredAt.Before(now) {
		return ErrUserFired
	}

	return nil
}

func ValidateUserName(name string) error {
	n := strings.TrimSpace(strings.ToLower(name))

	if l := utf8.RuneCountInString(n); l > maxUserNameLen {
		return throw.ErrInvalidLong
	} else if l < minUserNameLen {
		return throw.ErrInvalidShort
	}

	if !regValidateUserName.MatchString(n) {
		return throw.ErrInvalidIllegalChar
	}

	return nil
}

func ValidateUserStatus(status string) error {
	_, err := ParseUserStatus(status)
	return err
}

func ValidateUserGroup(group string) error {
	_, err := ParseUserGroup(group)
	return err
}

func ValidateUserGroups(groups []string) error {
	_, err := ParseUserGroups(groups)
	return err
}

func ValidateUserLogin(login string) error {
	if l := utf8.RuneCountInString(login); l < minUserLoginLen {
		return throw.ErrInvalidShort
	} else if l > maxUserLoginLen {
		return throw.ErrInvalidLong
	}

	if !regValidateUserLogin.MatchString(login) {
		return throw.ErrInvalidIllegalChar
	}

	return nil
}

// ValidateUserPassword одна цифра, заглавная, строчная буква, специальный символ и нет пробелов
func ValidateUserPassword(password string) error {
	if interValidator.Var(password, "ascii") != nil {
		return throw.ErrInvalidIllegalChar
	}

	var (
		hasNum, hasLower, hasUpper, hasSpecial bool
	)

	if l := utf8.RuneCountInString(password); l < minUserPasswordLen {
		return throw.ErrInvalidShort
	} else if l > maxUserPasswordLen {
		return throw.ErrInvalidLong
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
			return throw.ErrInvalidIllegalChar
		}
	}

	v := hasNum && hasLower && hasUpper && hasSpecial
	if !v {
		return throw.ErrInvalidPasswdWeak
	}

	return nil
}

func StringifyUserGroups(a []UserGroup) []string {
	b := make([]string, len(a))
	for i := range a {
		b[i] = string(a[i])
	}
	return b
}

func InflateUserGroups(a []string) []UserGroup {
	b := make([]UserGroup, len(a))
	for i := range a {
		b[i] = UserGroup(a[i])
	}
	return b
}
