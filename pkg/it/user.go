package it

import (
	"time"
)

const (
	// Размерность пароля
	maxUserEmailLen = 256 // Максимальная длина

	minUserPasswordLen = 6   // Минимальная длина
	maxUserPasswordLen = 128 // Максимальная длина

	minUserLoginLen = 6
	maxUserLoginLen = 50 // TODO синхронизировать с типом столбца хранения
)

type User struct {
	ID          uint32
	Status      UserStatus
	Roles       []UserRole
	StartWorkAt time.Time // Время начала работы
	FiredAt     *time.Time
	Avatar      *string
	Note        string
	FirstName   string
	Surname     string
}

// IsWorks работает ли сотрудник
func (u *User) IsWorks() error {
	if u.Status != UserStatusActive {
		return ErrUserHaveNotActiveStatus
	}
	now := time.Now()

	if u.StartWorkAt.After(now) {
		return ErrUserNotStartWork
	}

	if u.FiredAt != nil && u.FiredAt.Before(now) {
		return ErrUserFired
	}

	return nil
}

// LoggingUser авторизующийся по логину пользователь
type LoggingUser struct {
	User
	PasswdHash string
	Login      string
}

func (u *LoggingUser) ToUser() *User {
	return &u.User
}

func (u *LoggingUser) CanLogging() *User {
	return &User{
		ID:     u.ID,
		Status: u.Status,
		Roles:  u.Roles,
	}
}
