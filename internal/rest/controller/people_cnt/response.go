package people_cnt

import (
	"github.com/semenovem/portal/pkg/it"
)

// Общедоступный профиль пользователя
type userProfileView struct {
	ID        uint32 `json:"id"`
	FirstName string `json:"first_name,omitempty"`
	SurName   string `json:"sur_name,omitempty"`
	Position  string `json:"position,omitempty"`
	Avatar    string `json:"avatar,omitempty"` // TODO нужно собрать uri на загрузку аватара
}

func newUserProfileView(u *it.UserProfile) *userProfileView {
	r := &userProfileView{
		ID:        u.ID,
		FirstName: u.FirstName,
		SurName:   u.Surname,
		Position:  u.PositionName,
	}
	if u.Avatar != nil {
		r.Avatar = *u.Avatar
	}

	return r
}
