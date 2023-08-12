package peopleprovider

import (
	"context"
	"github.com/semenovem/portal/internal/provider"
	"github.com/semenovem/portal/pkg/entity"
)

func (p *Provider) GetUser(ctx context.Context, userID uint32) (*entity.User, error) {
	var (
		sq = `SELECT id, status, roles FROM people.users WHERE id = $1`
		u  = entity.User{}
	)

	err := p.db.QueryRow(ctx, sq, userID).
		Scan(&u.ID, &u.Status, &u.Roles)
	if err != nil {
		if !provider.IsNoRows(err) {
			p.logger.Named("GetUser").Error(err.Error())
		}

		return nil, err
	}

	return &u, nil
}

func (p *Provider) GetUserByLogin(ctx context.Context, loginQuery string) (*entity.LoggingUser, error) {
	var (
		sq = `SELECT id, login, passwd_hash FROM people.logged_users WHERE login = LOWER($1)`

		userID        uint32
		login, passwd string
	)

	err := p.db.QueryRow(ctx, sq, loginQuery).Scan(&userID, &login, &passwd)
	if err != nil {
		if !provider.IsNoRows(err) {
			p.logger.DBTag().Named("GetUserByLogin").Error(err.Error())
		}

		return nil, err
	}

	user, err := p.GetUser(ctx, userID)
	if err != nil {
		if !provider.IsNoRows(err) {
			p.logger.DBTag().Named("GetUserByLogin").Nested(err.Error())
		}

		return nil, err
	}

	return &entity.LoggingUser{
		User:         *user,
		PasswordHash: passwd,
		Login:        login,
	}, nil
}
