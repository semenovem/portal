package people_provider

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/semenovem/portal/internal/abc/provider"
	"github.com/semenovem/portal/pkg/it"
)

func (p *PeopleProvider) GetUserProfile(ctx context.Context, userID uint32) (*it.UserProfile, error) {
	var (
		sq = `SELECT id,
					status,
					roles,
					firstname,
					surname,
					avatar_id,
					note
       		FROM people.users
       		WHERE id = $1 AND deleted = false`

		m = UserModel{}
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
		if !provider.IsNoRow(err) {
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
			  firstname, surname, login, note,
			  status, roles, avatar_id,
			  expired_at, passwd_hash, props
			) VALUES (
				LOWER(@firstname), LOWER(@surname), LOWER(@login), @note,
				@status, @roles::people.roles_enum[], @avatar_id,
				@expired_at, @passwd_hash, @props
			) returning id`

		args = pgx.NamedArgs{
			"firstname":   m.firstname,
			"surname":     m.surname,
			"login":       m.login,
			"note":        m.note,
			"status":      m.status,
			"roles":       m.roles,
			"avatar_id":   m.avatarID,
			"expired_at":  m.expiredAt,
			"passwd_hash": m.passwdHash,
			"props":       m.props,
		}
	)

	if err = p.db.QueryRow(ctx, sq, args).Scan(&userID); err != nil && !provider.IsDuplicateKeyErr(err) {
		p.logger.Named("CreateUser").DB(err)
	}

	return userID, err
}

func (p *PeopleProvider) ExistsLoginName(ctx context.Context, loginName string) (exists bool, err error) {
	sq := `SELECT NOT EXISTS (select id from people.users where login = $1)`
	err = p.db.QueryRow(ctx, sq, loginName).Scan(&exists)

	return
}

func (p *PeopleProvider) DeleteUser(ctx context.Context, userID uint32) error {
	sq := `UPDATE people.users SET deleted = true WHERE deleted = false AND id = $1`

	result, err := p.db.Exec(ctx, sq, userID)
	if err != nil {
		p.logger.Named("DeleteUser").DB(err)
		return err
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}
