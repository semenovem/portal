package people_controller

import (
	people_dto "github.com/semenovem/portal/internal/abc/people/dto"
	"github.com/semenovem/portal/internal/util"
)

type userCreateResponse struct {
	UserID uint32 `json:"user_id"`
}

type freeLoginNameResponse struct {
	Free        bool   `json:"free"`
	ValidateErr string `json:"validate_err,omitempty"`
}

type publicHandbookResponse struct {
	Total     uint32                `json:"total"`
	Employees []*publicEmployeeView `json:"employees"`
}

type employeeUpdateResponse struct {
	FirstnameErr  *string `json:"firstname_err,omitempty"`
	SurnameErr    *string `json:"surname,omitempty"`
	NoteErr       *string `json:"note,omitempty"`
	StatusErr     *string `json:"status,omitempty"`
	RolesErr      *string `json:"roles,omitempty"`
	AvatarIDErr   *string `json:"avatar_id,omitempty"`
	LoginErr      *string `json:"login,omitempty"`
	PositionIDErr *string `json:"position_id,omitempty"`
	DeptIDErr     *string `json:"dept_id,omitempty"`
	ExpiredAtErr  *string `json:"expired_at,omitempty"`
	WorkedAtErr   *string `json:"worked_at,omitempty"`
	FiredAtErr    *string `json:"fired_at,omitempty"`
}

func newEmployeeUpdateResponse(dto *people_dto.UserProcessingErrDTO) *employeeUpdateResponse {
	return &employeeUpdateResponse{
		FirstnameErr:  util.ErrToPointStr(dto.FirstnameErr),
		SurnameErr:    util.ErrToPointStr(dto.SurnameErr),
		NoteErr:       util.ErrToPointStr(dto.NoteErr),
		StatusErr:     util.ErrToPointStr(dto.StatusErr),
		RolesErr:      util.ErrToPointStr(dto.RolesErr),
		AvatarIDErr:   util.ErrToPointStr(dto.AvatarIDErr),
		LoginErr:      util.ErrToPointStr(dto.LoginErr),
		PositionIDErr: util.ErrToPointStr(dto.PositionIDErr),
		DeptIDErr:     util.ErrToPointStr(dto.DeptIDErr),
		ExpiredAtErr:  util.ErrToPointStr(dto.ExpiredAtErr),
		WorkedAtErr:   util.ErrToPointStr(dto.WorkedAtErr),
		FiredAtErr:    util.ErrToPointStr(dto.FiredAtErr),
	}
}
