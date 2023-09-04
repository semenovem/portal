package people_provider

import (
	"github.com/semenovem/portal/pkg/it"
	"time"
)

type UserModel struct {
	id         uint32
	firstname  string
	surname    string
	deleted    bool
	status     it.UserStatus
	note       *string
	roles      *[]string
	roles2     *[]string
	avatarID   *uint32
	expiredAt  *time.Time
	login      *string
	passwdHash *string
	props      *it.UserProps
}

func (u *UserModel) ID() uint32 {
	return u.id
}

func (u *UserModel) SetID(id uint32) {
	u.id = id
}

func (u *UserModel) Firstname() string {
	return u.firstname
}

func (u *UserModel) SetFirstname(name string) {
	u.firstname = name
}

func (u *UserModel) Surname() string {
	return u.surname
}

func (u *UserModel) SetSurname(name string) {
	u.surname = name
}

func (u *UserModel) Deleted() bool {
	return u.deleted
}

func (u *UserModel) SetDeleted(deleted bool) {
	u.deleted = deleted
}

func (u *UserModel) Note() string {
	if u.note != nil {
		return *u.note
	}
	return ""
}

func (u *UserModel) SetNote(note string) {
	if note == "" {
		u.note = nil
	} else {
		u.note = &note
	}
}

func (u *UserModel) Status() it.UserStatus {
	return u.status
}

func (u *UserModel) SetStatus(status it.UserStatus) {
	if status == "" {
		u.status = it.UserStatusInactive
	} else {
		u.status = status
	}
}

func (u *UserModel) Roles() []it.UserRole {
	if u.roles == nil {
		return nil
	}

	return it.InflateUserRoles(*u.roles)
}

func (u *UserModel) SetRoles(roles []it.UserRole) {
	if len(roles) == 0 {
		u.roles = nil
	} else {
		rr := it.StringifyUserRoles(roles)
		u.roles = &rr
	}
}

func (u *UserModel) AvatarID() uint32 {
	if u.avatarID != nil {
		return *u.avatarID
	}

	return 0
}

func (u *UserModel) SetAvatarID(avatarID uint32) {
	if avatarID > 0 {
		u.avatarID = &avatarID
	} else {
		u.avatarID = nil
	}
}

func (u *UserModel) ExpiredAt() *time.Time {
	return u.expiredAt
}

func (u *UserModel) SetExpiredAt(expiredAt *time.Time) {
	if expiredAt != nil && !expiredAt.IsZero() {
		u.expiredAt = expiredAt
	} else {
		u.expiredAt = nil
	}
}

func (u *UserModel) Login() string {
	if u.login != nil {
		return *u.login
	}

	return ""
}

func (u *UserModel) SetLogin(login string) {
	if login != "" {
		u.login = &login
	} else {
		u.login = nil
	}
}

func (u *UserModel) PasswdHash() string {
	if u.passwdHash != nil {
		return *u.passwdHash
	}

	return ""
}

func (u *UserModel) SetPasswdHash(passwdHash string) {
	if passwdHash != "" {
		u.passwdHash = &passwdHash
	} else {
		u.passwdHash = nil
	}
}

func (u *UserModel) Props() *it.UserProps {
	return u.props
}

func (u *UserModel) SetProps(props *it.UserProps) {
	u.props = props
}
