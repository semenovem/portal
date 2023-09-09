package people_dto

import "time"

type EmployeeDTO struct {
	UserDTO
	PositionID *uint16
	DeptID     *uint16
	WorkedAt   *time.Time
	FiredAt    *time.Time
}
