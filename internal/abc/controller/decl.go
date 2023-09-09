package controller

import (
	"github.com/semenovem/portal/internal/audit"
	"github.com/semenovem/portal/internal/util"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/fail"
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

	ConditionalTimeVldTag = "conditional-time-vld-tag"
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

// ValidateConditionalTime Валидатор времени
func ValidateConditionalTime(s string) error {
	_, err := util.ParsePointerStringToTime(&s)
	return err
}
