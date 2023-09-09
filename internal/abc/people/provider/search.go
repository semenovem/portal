package people_provider

import (
	"context"
	"github.com/semenovem/portal/internal/abc/people"
	"github.com/semenovem/portal/internal/abc/people/dto"
)

func (p *PeopleProvider) SearchEmployees(
	ctx context.Context,
	opts *people_dto.EmployeesSearchOpts,
) ([]*people.EmployeeSlim, uint32, error) {
	var (
		// total uint32 TODO пока кол-во пользователей меньше лимита, не делать подсчет кол-ва
		ls = make([]*people.EmployeeSlim, 0)

		sq = `SELECT u.id, u.firstname, u.surname, u.note, u.status, u.roles, u.avatar_id,
		     e.dept_id, e.position_id, e.worked_at, e.fired_at
		FROM       people.users AS u
		LEFT JOIN  people.employees AS e ON e.user_id = u.id AND (e.fired_at IS NULL OR e.fired_at > now())
		WHERE u.deleted = false AND (u.expired_at IS NULL OR u.expired_at > now()) AND u.status = 'active'
		LIMIT $1 OFFSET $2`
	)

	rows, err := p.db.Query(ctx, sq, opts.Limit, opts.Offset)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	for rows.Next() {
		var o = people_dto.EmployeeDTO{}

		if err = rows.Scan(
			&o.ID,
			&o.Firstname,
			&o.Surname,
			&o.Note,
			&o.Status,
			&o.Roles,
			&o.AvatarID,
			&o.DeptID,
			&o.PositionID,
			&o.WorkedAt,
			&o.FiredAt,
		); err != nil {
			p.logger.Named("SearchEmployees.scan").DB(err)
			return nil, 0, err
		}

		ls = append(ls, o.ToEmployeeSlim())
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return ls, uint32(len(ls)), nil
}
