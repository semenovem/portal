package people_controller

import (
	"github.com/semenovem/portal/pkg/it"
	"strings"
	"time"
)

type UserForm struct {
	UserID uint32 `param:"user_id" validate:"required"`
}

type createUserForm struct {
	FirstName string     `json:"firstname" validate:"required,user-name-vld-tag"`
	Surname   string     `json:"surname" validate:"required,user-name-vld-tag"`
	Note      string     `json:"note" validate:"omitempty,user-name-vld-tag"`
	Status    string     `json:"status" validate:"omitempty,user-status-vld-tag"`
	Roles     []string   `json:"roles" validate:"omitempty,user-roles-vld-tag"`
	ExpiredAt *time.Time `json:"expired_at" binding:"datetime=2006-01-02T15:04:05Z07:00" validate:"omitempty"`
	Login     string     `json:"login" validate:"omitempty,user-login-vld-tag"`
	Passwd    string     `json:"passwd" validate:"omitempty,user-password-vld-tag"`
	AvatarID  uint32     `json:"avatar_id" validate:"omitempty"`
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
	return strings.ToLower(strings.TrimSpace(f.Note))
}

func (f *createUserForm) getLogin() string {
	return strings.ToLower(strings.TrimSpace(f.Note))
}
