package entity

const (
	// Размерность пароля
	maxUserEmailLen = 256 // Максимальная длина

	minUserPasswordLen = 8  // Минимальная длина
	maxUserPasswordLen = 20 // Максимальная длина

	minUserLoginLen = 6
	maxUserLoginLen = 50 // TODO синхронизировать с типом столбца хранения
)

type User struct {
	ID     uint32
	Status UserStatus
	Roles  []UserRole
}

// LoggingUser входящий пользователь
type LoggingUser struct {
	User
	PasswordHash string
	Login        string
}
