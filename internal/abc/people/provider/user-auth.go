package people_provider

import (
	"context"
	"errors"
	"github.com/semenovem/portal/internal/abc/provider"
	"github.com/semenovem/portal/pkg/it"
)

// GetUserAuth данные пользователя для авторизации по userID
func (p *PeopleProvider) GetUserAuth(ctx context.Context, userID uint32) (*it.UserAuth, error) {
	if userID == 0 {
		p.logger.Named("GetUserAuth").Error(userIDIsEmpty)
		return nil, errors.New(userIDIsEmpty)
	}

	u, err := p.getUserAuth(ctx, "", userID)
	if err != nil {
		p.logger.Named("GetUserAuth").Nested(err.Error())
		return nil, err
	}

	return u.ToUserAuth(), nil
}

// GetUserByLogin данные пользователя для авторизации по логину
func (p *PeopleProvider) GetUserByLogin(ctx context.Context, login string) (*it.UserLoginAuth, error) {
	if login == "" {
		p.logger.Named("GetUserByLogin").Error(userLoginIsEmpty)
		return nil, errors.New(userLoginIsEmpty)
	}

	u, err := p.getUserAuth(ctx, login, 0)
	if err != nil {
		p.logger.Named("GetUserByLogin").Nested(err.Error())
		return nil, err
	}

	return u, nil
}

// Данные пользователя для авторизации по логину или ID пользователя
func (p *PeopleProvider) getUserAuth(
	ctx context.Context,
	login string,
	userID uint32,
) (*it.UserLoginAuth, error) {
	var (
		sq = `SELECT id, status, roles, expired_at, COALESCE(passwd_hash, '')
				FROM people.users
				WHERE deleted = false`

		arg interface{}
		u   = it.UserLoginAuth{}
	)

	if login != "" {
		sq += " AND login = LOWER($1)"
		arg = login
	} else if userID != 0 {
		sq += " AND id = $1"
		arg = userID
	} else {
		panic("incorrect method call (login and userID is empty")
	}

	err := p.db.QueryRow(ctx, sq, arg).Scan(&u.ID, &u.Status, &u.Roles, &u.ExpiredAt, &u.PasswdHash)
	if err != nil {
		if !provider.IsNoRows(err) {
			p.logger.Named("getUserByLogin").DB(err)
		}

		return nil, err
	}

	sq = `SELECT worked_at, fired_at FROM people.employees WHERE user_id = $1`

	err = p.db.QueryRow(ctx, sq, u.ID).Scan(&u.WorkedAt, &u.FiredAt)
	if err != nil {
		if !provider.IsNoRows(err) {
			p.logger.Named("getUserByLogin").DB(err)
			return nil, err
		}
	}

	return &u, nil
}
