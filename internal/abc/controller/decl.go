package controller

import (
	"github.com/semenovem/portal/config"
	"github.com/semenovem/portal/internal/audit"
	"github.com/semenovem/portal/internal/util"
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
	UserGroupVldTag  = "user-group-vld-tag"
	UserGroupsVldTag = "user-groups-vld-tag"
	TimeVldTag       = "time-vld-tag"
)

type This struct {
	UserID uint32
}

type InitArgs struct {
	Logger         pkg.Logger
	FailureService *fail.Service
	MainConfig     *config.Main
	Audit          *audit.AuditProvider
	Common         *Common
}

// ValidateConditionalTime Валидатор времени
func ValidateConditionalTime(s string) error {
	_, err := util.ParsePointerStrToTime(&s)
	return err
}

func TimeToString(t *time.Time) string {
	if t == nil || t.IsZero() {
		return ""
	}

	return t.Format(time.RFC3339)
}
