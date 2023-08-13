package provider

import (
	"context"
	"github.com/semenovem/portal/pkg/it"
)

func (p *PeoplePvd) GetUser(ctx context.Context, userID uint32) (*it.User, error) {
	var (
		sq = `SELECT id, status, roles, start_work_at, fired_at, firstname, surname, avatar
       		FROM people.users
       		WHERE id = $1 AND deleted = false`
		u = it.User{}
	)

	err := p.db.QueryRow(ctx, sq, userID).
		Scan(&u.ID, &u.Status, &u.Roles, &u.StartWorkAt, &u.FiredAt, &u.FirstName, &u.Surname, &u.Avatar)
	if err != nil {
		if !IsNoRows(err) {
			p.logger.Named("GetUser").Error(err.Error())
		}

		return nil, err
	}

	return &u, nil
}

func (p *PeoplePvd) GetUserByLogin(ctx context.Context, loginQuery string) (*it.LoggingUser, error) {
	var (
		sq = `SELECT id, login, passwd_hash
				FROM people.users
				WHERE login = LOWER($1) AND deleted = false`

		userID        uint32
		login, passwd string
	)

	err := p.db.QueryRow(ctx, sq, loginQuery).Scan(&userID, &login, &passwd)
	if err != nil {
		if !IsNoRows(err) {
			p.logger.DBTag().Named("GetUserByLogin").Error(err.Error())
		}

		return nil, err
	}

	user, err := p.GetUser(ctx, userID)
	if err != nil {
		if !IsNoRows(err) {
			p.logger.DBTag().Named("GetUser").Nested(err.Error())
		}

		return nil, err
	}

	return &it.LoggingUser{
		User:       *user,
		PasswdHash: passwd,
		Login:      login,
	}, nil
}
