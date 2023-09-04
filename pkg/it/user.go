package it

import (
	"fmt"
	"strings"
	"time"
)

// UserCore основные данные сущности
type UserCore struct {
	ID     uint32
	Status UserStatus
	Roles  []UserRole
}

type UserProfile struct {
	UserCore
	AvatarID  uint32
	FirstName string
	Surname   string
	Note      string
	ExpiredAt *time.Time // Время автоматической блокировки
}

type EmployeeProfile struct {
	UserProfile
	Position    UserPosition // должность
	Boss        *UserBoss    // Руководитель
	StartWorkAt time.Time    // Дата начала работы
	FiredAt     *time.Time   // Дата увольнения
}

type UserBoss struct {
	ID        uint32
	Firstname string
	Surname   string
	UserPosition
}

func (u *UserProfile) ExpiredAtToString() string {
	if u.ExpiredAt == nil {
		return ""
	}

	return u.ExpiredAt.Format(time.RFC3339)
}

// ------------------------------------------------

type UserRole string
type UserStatus string

const (
	UserRoleSuperAdmin UserRole = "super-admin"
	UserRoleAdmin      UserRole = "admin"
	UserRoleUser       UserRole = "user"
)

const (
	UserStatusInactive UserStatus = "inactive"
	UserStatusActive   UserStatus = "active"
	UserStatusBlocked  UserStatus = "blocked"
)

type UserProps struct {
	Contacts []struct {
		Line1 string
		Note1 string
	}
}

func ParseUserStatus(s string) (UserStatus, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case string(UserStatusInactive):
		return UserStatusInactive, nil
	case string(UserStatusActive):
		return UserStatusActive, nil
	case string(UserStatusBlocked):
		return UserStatusBlocked, nil
	}

	return "", fmt.Errorf(msgErrUnknownUserStatus, s)
}

func ParseUserRole(s string) (UserRole, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case string(UserRoleSuperAdmin):
		return UserRoleSuperAdmin, nil
	case string(UserRoleAdmin):
		return UserRoleAdmin, nil
	case string(UserRoleUser):
		return UserRoleUser, nil
	}

	return "", fmt.Errorf(msgErrUnknownUserRole, s)
}

func ParseUserRoles(roles []string) ([]UserRole, error) {
	var (
		errs   = make([]string, 0)
		result = make([]UserRole, 0)
	)

	for _, role := range roles {
		if r, err := ParseUserRole(role); err != nil {
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

func StringifyUserRoles(a []UserRole) []string {
	b := make([]string, len(a))
	for i := range a {
		b[i] = string(a[i])
	}
	return b
}

func InflateUserRoles(a []string) []UserRole {
	b := make([]UserRole, len(a))
	for i := range a {
		b[i] = UserRole(a[i])
	}
	return b
}
