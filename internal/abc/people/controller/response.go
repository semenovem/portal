package people_controller

import (
	"github.com/semenovem/portal/pkg/it"
)

// Общедоступный профиль пользователя
type userPublicProfileView struct {
	ID        uint32 `json:"id"`
	Firstname string `json:"firstname,omitempty"`
	Surname   string `json:"surname,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
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

type userProfileView struct {
	userPublicProfileView
	Note      string `json:"note,omitempty"`
	ExpiredAt string `json:"expired_at"`
}

func newUserProfileView(u *it.UserProfile) *userProfileView {
	r := &userProfileView{
		userPublicProfileView: *newUserPublicProfileView(u),
		Note:                  u.Note,
		ExpiredAt:             u.ExpiredAtToString(),
	}

	return r
}
