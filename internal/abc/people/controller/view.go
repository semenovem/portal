package people_controller

import (
	"github.com/semenovem/portal/pkg/it"
	"time"
)

// Общедоступный профиль пользователя
type userPublicProfileView struct {
	ID        uint32 `json:"id"`
	Firstname string `json:"firstname,omitempty"`
	Surname   string `json:"surname,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
}

// Общедоступный профиль сотрудника
type publicEmployeeView struct {
	userPublicProfileView
	StartWorkAt  time.Time `json:"start_work_at"` // Дата начала работы
	PositionName string    `json:"position_name"`
	DeptName     string    `json:"dept_name"`
	Note         string    `json:"note,omitempty"`
	BossID       uint32    `json:"boss_id"`
}

type userProfileView struct {
	userPublicProfileView
	Note      string   `json:"note,omitempty"`
	ExpiredAt string   `json:"expired_at,omitempty"`
	Status    string   `json:"status,omitempty"`
	Roles     []string `json:"roles,omitempty"`
}

func newUserPublicProfileView(u *it.UserProfile) *userPublicProfileView {
	r := &userPublicProfileView{
		ID:        u.ID,
		Firstname: u.FirstName,
		Surname:   u.Surname,
	}
	if u.AvatarID != 0 {
		r.Avatar = "asdfasf/asdfasdf/"
	}

	return r
}

func newUserProfileView(u *it.UserProfile) *userProfileView {
	r := &userProfileView{
		userPublicProfileView: *newUserPublicProfileView(u),
		Note:                  u.Note,
		ExpiredAt:             u.ExpiredAtToString(),
		Status:                string(u.Status),
		Roles:                 it.StringifyUserRoles(u.Roles),
	}

	return r
}
