package people_dto

import "time"

type UserDTO struct {
	Firstname  *string
	Surname    *string
	Note       *string
	Status     *string
	Roles      *[]string
	AvatarID   *uint32
	Login      *string
	PositionID *uint16
	DeptID     *uint16
	ExpiredAt  *time.Time
	WorkedAt   *time.Time
	FiredAt    *time.Time
}

type UserProcessingErrDTO struct {
	FirstnameErr  error
	SurnameErr    error
	NoteErr       error
	StatusErr     error
	RolesErr      error
	AvatarIDErr   error
	LoginErr      error
	PositionIDErr error
	DeptIDErr     error
	ExpiredAtErr  error
	WorkedAtErr   error
	FiredAtErr    error
}
