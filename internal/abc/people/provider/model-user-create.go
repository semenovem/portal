package people_provider

import (
	"github.com/semenovem/portal/internal/abc/people"
	"html"
	"strings"
	"time"
)

type UserCreateModel struct {
	Firstname  *string
	Surname    *string
	Patronymic *string
	Status     *string
	Note       *string
	AvatarID   *uint32
	ExpiredAt  *time.Time
	Login      *string
	PasswdHash *string
	Props      *people.UserProps
}

func (u *UserCreateModel) getFirstname() string {
	if u.Firstname == nil {
		return ""
	}
	return strings.TrimSpace(*u.Firstname)
}

func (u *UserCreateModel) getSurname() string {
	if u.Surname == nil {
		return ""
	}
	return strings.TrimSpace(*u.Surname)
}

func (u *UserCreateModel) getPatronymic() string {
	if u.Patronymic == nil {
		return ""
	}
	return strings.TrimSpace(*u.Patronymic)
}

func (u *UserCreateModel) getNote() *string {
	if u.Note == nil {
		return nil
	}

	note := strings.TrimSpace(*u.Note)
	if note == "" {
		return nil
	}

	ss := html.EscapeString(note)
	return &ss
}

func (u *UserCreateModel) getStatus() string {
	if u.Status == nil {
		return string(people.UserStatusInactive)
	}

	return *u.Status
}

func (u *UserCreateModel) getAvatarID() *uint32 {
	if u.AvatarID == nil || *u.AvatarID == 0 {
		return nil
	}
	return u.AvatarID
}

func (u *UserCreateModel) getExpiredAt() *time.Time {
	if u.ExpiredAt == nil || u.ExpiredAt.IsZero() {
		return nil
	}
	return u.ExpiredAt
}

func (u *UserCreateModel) getLogin() *string {
	if u.Login == nil {
		return nil
	}

	login := strings.TrimSpace(*u.Login)

	if login == "" {
		return nil
	}

	return &login
}

func (u *UserCreateModel) getPasswdHash() *string {
	if u.PasswdHash == nil || *u.PasswdHash == "" {
		return nil
	}
	return u.PasswdHash
}

func (u *UserCreateModel) getProps() *people.UserProps {
	return u.Props
}

type EmployeeUpdateModel struct {
	UserCreateModel
	PositionID *uint16
	DeptID     *uint16
	WorkedAt   *time.Time
	FiredAt    *time.Time
}

func (m *EmployeeUpdateModel) getPositionID() uint16 {
	if m.PositionID == nil {
		return 0
	}
	return *m.PositionID
}

func (m *EmployeeUpdateModel) getDeptID() uint16 {
	if m.DeptID == nil {
		return 0
	}
	return *m.DeptID
}

func (m *EmployeeUpdateModel) getWorkedAt() time.Time {
	if m.WorkedAt == nil {
		return time.Time{}
	}
	return *m.WorkedAt
}

func (m *EmployeeUpdateModel) getFiredAt() *time.Time {
	if m.FiredAt == nil || m.FiredAt.IsZero() {
		return nil
	}
	return m.FiredAt
}
