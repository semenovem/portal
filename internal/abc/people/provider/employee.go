package people_provider

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/semenovem/portal/internal/abc/people/dto"
	"github.com/semenovem/portal/internal/abc/provider"
)

func (p *PeopleProvider) CreateEmployee(
	ctx context.Context,
	dto *people_dto.EmployeeDTO,
) (userID uint32, err error) {
	tx, err := p.db.Begin(ctx)
	if err != nil {
		p.logger.Named("CreateEmployee.Begin").DB(err)
		return 0, err
	}

	defer func() {
		if err1 := tx.Rollback(ctx); err1 != nil && !errors.Is(err1, pgx.ErrTxClosed) {
			p.logger.Named("Rollback").DB(err1)
		}
	}()

	userID, err = p.createUserTx(ctx, tx, &dto.UserDTO)
	if err != nil {
		return 0, err
	}

	if err = p.createEmployeeTx(ctx, tx, userID, dto); err != nil {
		return 0, err
	}

	if err = tx.Commit(ctx); err != nil {
		p.logger.Named("Commit").DB(err)
		return 0, err
	}

	return userID, err
}

func (p *PeopleProvider) createUserTx(
	ctx context.Context,
	tx pgx.Tx,
	dto *people_dto.UserDTO,
) (userID uint32, err error) {
	var (
		sq = `INSERT INTO people.users (
			  firstname, surname, login, note,
			  status, roles, avatar_id,
			  expired_at, passwd_hash
			) VALUES (
				LOWER(@firstname), LOWER(@surname), LOWER(@login), @note,
				@status, @roles::people.roles_enum[], @avatar_id,
				@expired_at, @passwd_hash
			) returning id`

		args = pgx.NamedArgs{
			"firstname":   *dto.Firstname,
			"surname":     *dto.Surname,
			"login":       dto.Login,
			"note":        dto.Note,
			"status":      dto.Status,
			"roles":       dto.Roles,
			"avatar_id":   dto.AvatarID,
			"expired_at":  dto.ExpiredAt,
			"passwd_hash": dto.PasswdHash,
		}
	)

	if err = tx.QueryRow(ctx, sq, args).Scan(&userID); err != nil {
		if n, ok := provider.ConstraintErr(err); ok {
			return 0, constraintErr(n, err)
		}

		p.logger.Named("CreateUser").DB(err)
	}

	return userID, nil
}

func (p *PeopleProvider) createEmployeeTx(
	ctx context.Context,
	tx pgx.Tx,
	userID uint32,
	dto *people_dto.EmployeeDTO,
) error {
	var (
		sq = `INSERT INTO people.employees
		       ( user_id,  position_id,  dept_id,  worked_at,  fired_at)
		VALUES (@user_id, @position_id, @dept_id, @worked_at, @fired_at)
		ON CONFLICT (user_id) DO UPDATE SET
			position_id = @position_id,
			dept_id     = @dept_id,
			worked_at   = @worked_at,
			fired_at    = @fired_at`

		args = pgx.NamedArgs{
			"user_id":     userID,
			"position_id": *dto.PositionID,
			"dept_id":     *dto.DeptID,
			"worked_at":   *dto.WorkedAt,
			"fired_at":    dto.FiredAt,
		}
	)

	if _, err := tx.Exec(ctx, sq, args); err != nil {
		if n, ok := provider.ConstraintErr(err); ok {
			return constraintErr(n, err)
		}

		p.logger.Named("createEmployeeTx").DB(err)
	}

	return nil
}

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

// Поиск руководителя отдела
// 1. получаем руководителя из таблицы `people.head_of_dept` по отделу пользователя
// 2. если не найдено: рекурсивно получим отдел из parent_id отдела пользователя

func (p *PeopleProvider) GetBoss(ctx context.Context, deptID uint16) (*EmployeeModel, error) {

	return nil, nil
}
