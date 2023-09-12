package people

import (
	"github.com/semenovem/portal/pkg/throw"
	"strings"
	"time"
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

func ParseUserStatusIfDefault(s string) UserStatus {
	st, err := ParseUserStatus(s)
	if err != nil {
		return UserStatusInactive
	}
	return st
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
