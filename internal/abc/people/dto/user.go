package people_dto

import (
	"github.com/semenovem/portal/internal/abc/people"
	"time"
)

type UserDTO struct {
	ID         uint32
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

func (dto *EmployeeDTO) ToEmployeeSlim() *people.EmployeeSlim {
	return &people.EmployeeSlim{
		ID:          0,
		Status:      "",
		Roles:       nil,
		AvatarID:    0,
		FirstName:   "",
		Surname:     "",
		Note:        "",
		PositionID:  0,
		DeptID:      0,
		StartWorkAt: time.Time{},
		FiredAt:     nil,
	}
}
