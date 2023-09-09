package people_controller

import (
	"github.com/semenovem/portal/internal/util"
	"github.com/semenovem/portal/pkg/throw"
	"html"
	"strings"
	"time"
)

type userPathForm struct {
	UserID uint32 `json:"user_id" param:"user_id" validate:"required"`
}

type createUserForm struct {
	FirstName  string   `json:"firstname" validate:"required,user-name-vld-tag"`
	Surname    string   `json:"surname" validate:"required,user-name-vld-tag"`
	Note       string   `json:"note" validate:"omitempty"`
	Status     string   `json:"status" validate:"omitempty,user-status-vld-tag"`
	Roles      []string `json:"roles" validate:"omitempty,user-roles-vld-tag"`
	Login      string   `json:"login" validate:"omitempty,user-login-vld-tag"`
	Passwd     string   `json:"passwd" validate:"omitempty,user-password-vld-tag"`
	AvatarID   uint32   `json:"avatar_id" validate:"omitempty"`
	PositionID uint16   `json:"position_id" validate:"required,min=1"`
	DeptID     uint16   `json:"dept_id" validate:"required,min=1"`
	ExpiredAt  string   `json:"expired_at" validate:"omitempty,time-vld-tag"`
	WorkedAt   string   `json:"worked_at" validate:"required,time-vld-tag"`
	FiredAt    string   `json:"fired_at" validate:"omitempty,time-vld-tag"`
}

func (f *createUserForm) getFirstname() *string {
	return util.NormLowerStrToPointer(f.FirstName)
}

func (f *createUserForm) getSurname() *string {
	return util.NormLowerStrToPointer(f.Surname)
}

func (f *createUserForm) getNote() *string {
	note := html.EscapeString(strings.TrimSpace(f.Note))
	return &note
}

func (f *createUserForm) getLogin() *string {
	return util.NormLowerStrToPointer(f.Login)
}

func (f *createUserForm) getExpiredAt() *time.Time {
	t, err := util.ParseNotEmptyStrToTime(&f.ExpiredAt)
	if err != nil {
		panic(throw.NewInvalidTimeFieldErr("expired_at", err))
	}
	return t
}

func (f *createUserForm) getWorkedAt() *time.Time {
	t, err := util.ParseNotEmptyStrToTime(&f.WorkedAt)
	if err != nil {
		panic(throw.NewInvalidTimeFieldErr("worked_at", err))
	}
	return t
}

func (f *createUserForm) getFiredAt() *time.Time {
	t, err := util.ParseNotEmptyStrToTime(&f.FiredAt)
	if err != nil {
		panic(throw.NewInvalidTimeFieldErr("fired_at", err))
	}
	return t
}

type freeLoginNameForm struct {
	LoginName string `param:"login_name" validate:"required"`
}

type handbookForm struct {
	DeptID     []uint16 `json:"dept_id[]" query:"dept_id[]" validate:"omitempty"`
	PositionID []uint16 `json:"position_id[]" query:"position_id[]" validate:"omitempty"`
	Order      []string `json:"order[]" query:"order[]" validate:"omitempty"`
}

type employeeUpdateForm struct {
	UserID     uint32    `json:"user_id" param:"user_id" validate:"required" swaggerignore:"true"`
	Firstname  *string   `json:"firstname" validate:"omitempty,user-name-vld-tag"`
	Surname    *string   `json:"surname" validate:"omitempty,user-name-vld-tag"`
	Note       *string   `json:"note"`
	Status     *string   `json:"status" validate:"omitempty,user-status-vld-tag"`
	Roles      *[]string `json:"roles" validate:"omitempty,user-roles-vld-tag"`
	AvatarID   *uint32   `json:"avatar_id"`
	Login      *string   `json:"login" validate:"omitempty,user-login-vld-tag"`
	PositionID *uint16   `json:"position_id" validate:"omitempty"`
	DeptID     *uint16   `json:"dept_id" validate:"omitempty"`
	ExpiredAt  *string   `json:"expired_at" validate:"omitempty,time-vld-tag"`
	WorkedAt   *string   `json:"worked_at" validate:"omitempty,time-vld-tag"`
	FiredAt    *string   `json:"fired_at" validate:"omitempty,time-vld-tag"`
}

func (f *employeeUpdateForm) getExpiredAt() *time.Time {
	t, err := util.ParsePointerStrToTime(f.ExpiredAt)
	if err != nil {
		panic(throw.NewInvalidTimeFieldErr("expired_at", err))
	}
	return t
}

func (f *employeeUpdateForm) getWorkedAt() *time.Time {
	t, err := util.ParsePointerStrToTime(f.WorkedAt)
	if err != nil {
		panic(throw.NewInvalidTimeFieldErr("worked_at", err))
	}
	return t
}

func (f *employeeUpdateForm) getFiredAt() *time.Time {
	t, err := util.ParsePointerStrToTime(f.FiredAt)
	if err != nil {
		panic(throw.NewInvalidTimeFieldErr("fired_at", err))
	}
	return t
}
