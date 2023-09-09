package people_dto

import "time"

type UserDTO struct {
	Firstname  *string
	Surname    *string
	Note       *string
	Status     *string
	Roles      *[]string
	AvatarID   *uint32
	ExpiredAt  *time.Time
	Login      *string
	PasswdHash *string
}

type EmployeeDTO struct {
	UserDTO
	PositionID *uint16
	DeptID     *uint16
	WorkedAt   *time.Time
	FiredAt    *time.Time
}
