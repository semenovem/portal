package people_provider

import (
	"context"
	"github.com/semenovem/portal/internal/abc/provider"
	"github.com/semenovem/portal/pkg/throw"
)

func (p *PeopleProvider) GetUserModel(ctx context.Context, userID uint32) (*UserModel, error) {
	var (
		sq = `SELECT id,
				firstname,
				surname,
				patronymic,
				status,
				avatar_id,
				note,
				expired_at,
				login,
				props
       		FROM people.users
       		WHERE id = $1 AND deleted = false`

		m = UserModel{}
	)

	err := p.db.QueryRow(ctx, sq, userID).
		Scan(
			&m.id,
			&m.firstname,
			&m.surname,
			&m.patronymic,
			&m.status,
			&m.avatarID,
			&m.note,
			&m.expiredAt,
			&m.login,
			&m.props,
		)
	if err != nil {
		if provider.IsNoRow(err) {
			return nil, throw.Err404User
		}
		p.logger.Named("GetUserProfile").DB(err)

		return nil, err
	}

	return &m, nil
}

func (p *PeopleProvider) GetEmployeeModel(ctx context.Context, userID uint32) (*EmployeeModel, error) {
	var (
		sq = `SELECT u.id,
					u.firstname,
					u.surname,
					u.patronymic,
					u.status,
					u.avatar_id,
					u.note,
					u.expired_at,
					u.login,
					u.props,
					e.dept_id,
					e.position_id,
					e.worked_at,
					e.fired_at
       		FROM people.employees  AS e
       		LEFT JOIN people.users AS u ON u.id = e.user_id
       		WHERE id = $1 AND deleted = false`

		m = EmployeeModel{}
	)

	err := p.db.QueryRow(ctx, sq, userID).
		Scan(
			&m.id,
			&m.firstname,
			&m.surname,
			&m.patronymic,
			&m.status,
			&m.avatarID,
			&m.note,
			&m.expiredAt,
			&m.login,
			&m.props,
			&m.props,
			&m.deptID,
			&m.positionID,
			&m.workedAt,
			&m.firedAt,
		)
	if err != nil {
		if provider.IsNoRow(err) {
			return nil, throw.Err404User
		}
		p.logger.Named("GetUserProfile").DB(err)

		return nil, err
	}

	return &m, nil
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
		return throw.Err404User
	}

	return nil
}

func (p *PeopleProvider) UndeleteUser(ctx context.Context, userID uint32) error {
	sq := `UPDATE people.users SET deleted = false WHERE deleted = true AND id = $1`

	result, err := p.db.Exec(ctx, sq, userID)
	if err != nil {
		p.logger.Named("UndeleteUser").DB(err)
		return err
	}

	if result.RowsAffected() == 0 {
		return throw.Err404User
	}

	return nil
}
