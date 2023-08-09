package entity

const (
	// Размерность пароля
	minUserPasswordLen = 8  // Минимальная длина
	maxUserPasswordLen = 20 // Максимальная длина

)

type User struct {
	ID uint32
}
