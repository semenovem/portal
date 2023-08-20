package it

import "time"

// UserAuth авторизующийся пользователь
type UserAuth struct {
	ID        uint32
	Status    UserStatus
	Roles     []UserRole
	ExpiredAt *time.Time // Активна до указанного времени
	WorkedAt  *time.Time // Дата начала работы
	FiredAt   *time.Time // Дата увольнения
}

func (u *UserAuth) CanLogging() error {
	if u.Status != UserStatusActive {
		return ErrUserHaveNotActiveStatus
	}
	now := time.Now()

	if u.ExpiredAt != nil && u.FiredAt.Before(now) {
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

// UserLoginAuth авторизующийся по логину пользователь
type UserLoginAuth struct {
	UserAuth
	PasswdHash string
}

func (u *UserLoginAuth) ToUserAuth() *UserAuth {
	return &u.UserAuth
}
