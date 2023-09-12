package people_provider

import (
	"github.com/semenovem/portal/internal/abc/people"
	"github.com/semenovem/portal/pkg/it"
	"time"
)

type UserModel struct {
	id         *uint32
	firstname  *string
	surname    *string
	patronymic *string
	status     *string
	note       *string
	roles      *[]string
	avatarID   *uint32
	expiredAt  *time.Time
	login      *string
	props      *it.UserProps
	updateAt   time.Time
}

func (u *UserModel) ID() uint32 {
	if u.id == nil {
		return 0
	}
	return *u.id
}

func (u *UserModel) Firstname() string {
	if u.firstname == nil {
		return ""
	}
	return *u.firstname
}

func (u *UserModel) Surname() string {
	if u.surname == nil {
		return ""
	}
	return *u.surname
}

func (u *UserModel) Patronymic() string {
	if u.patronymic == nil {
		return ""
	}
	return *u.patronymic
}

func (u *UserModel) Note() string {
	if u.note == nil {
		return ""
	}
	return *u.note
}

func (u *UserModel) Status() people.UserStatus {
	if u.status == nil {
		return people.UserStatusInactive
	}

	return people.ParseUserStatusIfDefault(*u.status)
}

func (u *UserModel) Roles() []it.UserRole {
	if u.roles == nil {
		return nil
	}

	return it.InflateUserRoles(*u.roles)
}

func (u *UserModel) AvatarID() uint32 {
	if u.avatarID == nil {
		return 0
	}
	return *u.avatarID
}

func (u *UserModel) ExpiredAt() *time.Time {
	if u.expiredAt == nil || u.expiredAt.IsZero() {
		return nil
	}
	return u.expiredAt
}

func (u *UserModel) Login() string {
	if u.login == nil {
		return ""
	}

	return *u.login
}

func (u *UserModel) Props() *it.UserProps {
	return u.props
}

func (u *UserModel) UpdatedAt() time.Time {
	return u.updateAt
}

type EmployeeModel struct {
	UserModel
	updateAt   time.Time
	positionID *uint16
	deptID     *uint16
	workedAt   *time.Time
	firedAt    *time.Time
}

func (m *EmployeeModel) UpdatedAt() time.Time {
	return m.updateAt
}

func (m *EmployeeModel) PositionID() uint16 {
	if m.positionID == nil {
		return 0
	}
	return *m.positionID
}

func (m *EmployeeModel) DeptID() uint16 {
	if m.deptID == nil {
		return 0
	}
	return *m.deptID
}

func (m *EmployeeModel) WorkedAt() *time.Time {
	if m.workedAt == nil || m.workedAt.IsZero() {
		return nil
	}
	return m.workedAt
}

func (m *EmployeeModel) FiredAt() *time.Time {
	if m.firedAt == nil || m.firedAt.IsZero() {
		return nil
	}
	return m.firedAt
}
