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

// UserCore основные данные сущности
type UserCore struct {
	ID     uint32
	Status UserStatus
	Roles  []UserRole
}

type UserProfile struct {
	UserCore
	Avatar        *string
	FirstName     string
	Surname       string
	PositionTitle string
	Note          string
}

type EmployeeProfile struct {
	UserProfile
	Position    UserPosition // должность
	Boss        *UserBoss    // Руководитель
	StartWorkAt *time.Time   // Дата начала работы
	FiredAt     *time.Time   // Дата увольнения
	ExpiredAt   *time.Time   // Время автоматической блокировки
}

type UserBoss struct {
	ID        uint32
	Firstname string
	Surname   string
	UserPosition
}
