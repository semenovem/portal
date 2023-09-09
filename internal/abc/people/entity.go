package people

import (
	"time"
)

type UserPosition struct {
	ID       uint16
	Title    string
	ParentID uint16
}

type UserDept struct {
	ID       uint16
	Title    string
	ParentID uint16
}

type Employee struct {
	ID        uint32
	Status    UserStatus
	Roles     []UserRole
	AvatarID  uint32
	FirstName string
	Surname   string
	Note      string
	ExpiredAt *time.Time

	PositionID  uint16     // должность
	DeptID      uint16     // отдел
	StartWorkAt time.Time  // Дата начала работы
	FiredAt     *time.Time // Дата увольнения
}

type UserBoss struct {
	ID            uint32
	Firstname     string
	Surname       string
	PositionTitle string
}
