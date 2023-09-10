package people_provider

import (
	"context"
	"github.com/semenovem/portal/internal/abc/people/dto"
)

func (p *PeopleProvider) SearchEmployees(
	ctx context.Context,
	opts *people_dto.EmployeesSearchOpts,
) ([]*EmployeeModel, uint32, error) {
	var (
		// total uint32 TODO пока кол-во пользователей меньше лимита, не делать подсчет кол-ва
		ls = make([]*EmployeeModel, 0)

		sq = `SELECT u.id, u.firstname, u.surname, u.note, u.status, u.roles, u.avatar_id,
		     e.dept_id, e.position_id, e.worked_at, e.fired_at
		FROM       people.employees AS e
		LEFT JOIN  people.users     AS u ON e.user_id = u.id AND (e.fired_at IS NULL OR e.fired_at > now())
		WHERE u.deleted = false AND (u.expired_at IS NULL OR u.expired_at > now()) AND u.status = 'active'
		LIMIT $1 OFFSET $2`
	)

	rows, err := p.db.Query(ctx, sq, opts.Limit, opts.Offset)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	for rows.Next() {
		var o EmployeeModel

		if err = rows.Scan(
			&o.id,
			&o.firstname,
			&o.surname,
			&o.note,
			&o.status,
			&o.roles,
			&o.avatarID,
			&o.deptID,
			&o.positionID,
			&o.workedAt,
			&o.firedAt,
		); err != nil {
			p.logger.Named("SearchEmployees.scan").DB(err)
			return nil, 0, err
		}

		ls = append(ls, &o)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return ls, uint32(len(ls)), nil
}
