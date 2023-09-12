package people_controller

import (
	"github.com/semenovem/portal/pkg/it"
	"time"
)

// Общедоступный профиль пользователя
type userPublicProfileView struct {
	ID         uint32 `json:"id"`
	Firstname  string `json:"firstname,omitempty"`
	Surname    string `json:"surname,omitempty"`
	Patronymic string `json:"patronymic,omitempty"`
	Avatar     string `json:"avatar,omitempty"`
}

// Профиль сотрудника
type employeeProfileView struct {
	userPublicProfileView
	WorkedAt     *time.Time `json:"worked_at"`          // Дата начала работы
	FiredAt      *time.Time `json:"fired_at,omitempty"` // Дата увольнения
	DeptName     string     `json:"dept_name"`
	PositionName string     `json:"position_name"`
	Note         string     `json:"note,omitempty"`
	BossID       uint32     `json:"boss_id"`
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

func newUserPublicProfileView2(u *it.UserProfile) *userPublicProfileView {
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
