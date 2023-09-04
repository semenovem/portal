package people_controller

import (
	"github.com/semenovem/portal/pkg/it"
	"html"
	"strings"
	"time"
)

type userForm struct {
	UserID uint32 `param:"user_id" validate:"required"`
}

type createUserForm struct {
	FirstName string    `json:"firstname" validate:"required,user-name-vld-tag"`
	Surname   string    `json:"surname" validate:"required,user-name-vld-tag"`
	Note      string    `json:"note" validate:"omitempty"`
	Status    string    `json:"status" validate:"omitempty,user-status-vld-tag"`
	Roles     []string  `json:"roles" validate:"omitempty,user-roles-vld-tag"`
	ExpiredAt time.Time `json:"expired_at" validate:"omitempty"`
	Login     string    `json:"login" validate:"omitempty,user-login-vld-tag"`
	Passwd    string    `json:"passwd" validate:"omitempty,user-password-vld-tag"`
	AvatarID  uint32    `json:"avatar_id" validate:"omitempty"`
}

func (f *createUserForm) getStatus() it.UserStatus {
	st, _ := it.ParseUserStatus(f.Status)
	return st
}

func (f *createUserForm) getRoles() []it.UserRole {
	roles, _ := it.ParseUserRoles(f.Roles)
	return roles
}

func (f *createUserForm) getFirstname() string {
	return strings.ToLower(strings.TrimSpace(f.FirstName))
}

func (f *createUserForm) getSurname() string {
	return strings.ToLower(strings.TrimSpace(f.Surname))
}

func (f *createUserForm) getNote() string {
	return html.EscapeString(strings.TrimSpace(f.Note))
}

func (f *createUserForm) getLogin() string {
	return strings.ToLower(strings.TrimSpace(f.Login))
}

type freeLoginNameForm struct {
	LoginName string `param:"login_name" validate:"required"`
}

type handbookForm struct {
	DeptID     []uint16 `json:"dept_id[]" query:"dept_id[]" validate:"omitempty"`
	PositionID []uint16 `json:"position_id[]" query:"position_id[]" validate:"omitempty"`
	Order      []string `json:"order[]" query:"order[]" validate:"omitempty"`
}
