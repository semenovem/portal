package people_provider

import (
	"context"
	"fmt"
	"github.com/semenovem/portal/internal/abc/provider"
	"github.com/semenovem/portal/pkg/it"
)

func (p *PeopleProvider) GetUserProfile(ctx context.Context, userID uint32) (*it.UserProfile, error) {
	var (
		sq = `SELECT u.id,
					u.status,
					u.roles,
					u.firstname,
					u.surname,
					u.avatar_id,
					u.note
       		FROM people.users* AS u
       		LEFT JOIN people.employees AS em ON em.user_id = u.id
       		WHERE u.id = $1 AND u.deleted = false`

		m = UserModel{
			id:         0,
			firstname:  "",
			surname:    "",
			deleted:    false,
			status:     "",
			note:       nil,
			roles:      nil,
			avatarID:   nil,
			expiredAt:  nil,
			login:      nil,
			passwdHash: nil,
			props:      nil,
		}
	)

	err := p.db.QueryRow(ctx, sq, userID).
		Scan(
			&m.id,
			&m.status,
			&m.roles,
			&m.firstname,
			&m.surname,
			&m.avatarID,
			&m.note,
		)
	if err != nil {
		if !provider.IsNoRows(err) {
			p.logger.Named("GetUserProfile").DB(err)
		}

		return nil, err
	}

	return &it.UserProfile{
		UserCore: it.UserCore{
			ID:     m.id,
			Status: m.status,
			Roles:  m.Roles(),
		},
		AvatarID:  m.AvatarID(),
		FirstName: m.Firstname(),
		Surname:   m.surname,
		Note:      m.Note(),
		ExpiredAt: m.expiredAt,
	}, nil
}

func (p *PeopleProvider) CreateUser(ctx context.Context, m *UserModel) (userID uint32, err error) {
	var (
		sq = `INSERT INTO people.users (
			  firstname,
			  surname,
			  login,
			  note,
			  status,
			  roles,
			  avatar_id,
			  expired_at,
			  passwd_hash,
			  props
			  ) VALUES (LOWER($1), LOWER($2), LOWER($3), $4, $5, $6, $7, $8, $9, $10) returning id`
	)

	fmt.Println(">>>>>>>>>>>>>>>>> id ", m.ExpiredAt())
	fmt.Println(">>>>>>>>>>>>>>>>> ExpiredAt = ", m.ExpiredAt())

	err = p.db.QueryRow(ctx, sq,
		m.firstname,
		m.surname,
		m.login,
		m.note,
		m.status,
		m.roles,
		m.avatarID,
		m.expiredAt,
		m.passwdHash,
		m.props,
	).Scan(&userID)
	if err != nil && !provider.IsDuplicateKeyError(err) {
		p.logger.Named("CreateUser").DB(err)
	}

	return userID, err
}
