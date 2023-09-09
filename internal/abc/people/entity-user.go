package people

import (
	"github.com/semenovem/portal/pkg/throw"
	"strings"
)

var (
	ErrUnknownUserStatus = throw.NewInvalidErr("unknown user status")
	ErrUnknownUserRole   = throw.NewInvalidErr("unknown user role")
)

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

func ParseUserRole(s string) (UserRole, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case string(UserRoleSuperAdmin):
		return UserRoleSuperAdmin, nil
	case string(UserRoleAdmin):
		return UserRoleAdmin, nil
	case string(UserRoleUser):
		return UserRoleUser, nil
	}

	return "", throw.NewWithTargetErrf(ErrUnknownUserRole, "origin:[%s]", s)
}
