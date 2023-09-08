package controller

import (
	"github.com/semenovem/portal/internal/audit"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/fail"
	"time"
)

const (
	ThisUserID = "this_user_id"
)

/* Теги валидаторов */
const (
	UserLoginVldTag  = "user-login-vld-tag"
	UserPasswordTag  = "user-password-vld-tag"
	UserEmailVldTag  = "user-email-vld-tag"
	UserNameVldTag   = "user-name-vld-tag"
	UserStatusVldTag = "user-status-vld-tag"
	UserRoleVldTag   = "user-role-vld-tag"
	UserRolesVldTag  = "user-roles-vld-tag"
)

type This struct {
	UserID uint32
}

type CntArgs struct {
	Logger         pkg.Logger
	FailureService *fail.Service
	Audit          *audit.AuditProvider
	Common         *Common
}

func ParseTime(s *string) (*time.Time, error) {
	if s == nil {
		return nil, nil
	}

	if *s == "" {
		return &time.Time{}, nil
	}

	t, err := time.Parse(time.RFC3339, *s)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
