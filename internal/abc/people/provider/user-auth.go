package people_provider

import (
	"context"
	"github.com/semenovem/portal/internal/abc/people"
	"github.com/semenovem/portal/internal/abc/provider"
	"github.com/semenovem/portal/pkg/throw"
	"github.com/semenovem/portal/pkg/txt"
)

// GetUserAuth данные пользователя для авторизации по userID
func (p *PeopleProvider) GetUserAuth(ctx context.Context, userID uint32) (*people.UserAuth, error) {
	return p.getUserAuth(ctx, userID, "", "")
}

// GetUserByLogin данные пользователя для авторизации по логину
func (p *PeopleProvider) GetUserByLogin(
	ctx context.Context,
	login, passwdHash string,
) (*people.UserAuth, error) {
	return p.getUserAuth(ctx, 0, login, passwdHash)
}

// Данные пользователя для авторизации по логину или ID пользователя
func (p *PeopleProvider) getUserAuth(
	ctx context.Context,
	userID uint32,
	login, passwdHash string,
) (*people.UserAuth, error) {
	const label = "PeopleProvider.getUserAuth"
	var (
		sq = `SELECT u.id, u.status, u.expired_at, e.worked_at, e.fired_at
		FROM      people.users     AS u
		LEFT JOIN people.employees AS e ON e.user_id = u.id
		WHERE u.deleted = false
		  AND (
		      (u.id = $1 AND u.id <> 0)
		           OR
		      (u.login = LOWER($2) AND u.passwd_hash = LOWER($3) AND u.passwd_hash <> '')
		  )`

		u people.UserAuth
	)

	err := p.db.QueryRow(ctx, sq, userID, login, passwdHash).Scan(
		&u.ID,
		&u.Status,
		&u.ExpiredAt,
		&u.WorkedAt,
		&u.FiredAt,
	)
	if err != nil {
		if provider.IsNoRow(err) {
			err = throw.NewNotFound(txt.NotFoundUser)
		}

		return nil, throw.Trace(err, label, map[string]any{
			"userID": userID,
		})
	}

	return &u, nil
}
