package user_tool

import (
	"github.com/semenovem/portal/pkg/tool/user_tool/user_err"
	"time"
)

type UserAuth struct {
	ID        uint32
	Status    Status
	ExpiredAt *time.Time
	WorkedAt  *time.Time
	FiredAt   *time.Time
}

func (u *UserAuth) CanLogging() error {
	if u.Status != StatusActive {
		return user_err.NotActive.Err()
	}
	now := time.Now()

	if u.ExpiredAt != nil && u.ExpiredAt.Before(now) {
		return user_err.Expired.Err()
	}

	if u.WorkedAt != nil && u.WorkedAt.After(now) {
		return user_err.NotStartWork.Err()
	}

	if u.FiredAt != nil && u.FiredAt.Before(now) {
		return user_err.Fired.Err()
	}

	return nil
}
