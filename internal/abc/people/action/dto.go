package people_action

import (
	people_provider "github.com/semenovem/portal/internal/abc/people/provider"
	"github.com/semenovem/portal/pkg/it"
	"time"
)

type CreateUserDTO struct {
	FirstName string
	Surname   string
	Note      string
	Status    it.UserStatus
	Roles     []it.UserRole
	ExpiredAt time.Time
	Login     string
	Passwd    string
	AvatarID  uint32
}

func (dto *CreateUserDTO) toPvdModel(passwdHash string) *people_provider.UserModel {
	m := people_provider.UserModel{}

	m.SetFirstname(dto.FirstName)
	m.SetSurname(dto.Surname)
	m.SetNote(dto.Note)
	m.SetStatus(dto.Status)
	m.SetRoles(dto.Roles)
	m.SetExpiredAt(&dto.ExpiredAt)
	m.SetLogin(dto.Login)
	m.SetPasswdHash(passwdHash)
	m.SetAvatarID(dto.AvatarID)

	return &m
}

type LoginNameCheckDTO struct {
	Exists      bool
	ValidateErr error
}