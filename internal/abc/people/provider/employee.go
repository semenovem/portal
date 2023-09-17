package people_provider

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/semenovem/portal/internal/abc/provider"
)

func (p *PeopleProvider) CreateEmployee(
	ctx context.Context,
	dto *EmployeeUpdateModel,
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

	userID, err = p.createUserTx(ctx, tx, &dto.UserCreateModel)
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
	dto *UserCreateModel,
) (userID uint32, err error) {
	var (
		sq = `INSERT INTO people.users (
			  firstname, surname, patronymic, login, note,
			  status, avatar_id, expired_at, passwd_hash
			) VALUES (
				LOWER(@firstname), LOWER(@surname), LOWER(@patronymic), LOWER(@login), @note,
				@status, @avatar_id, @expired_at, @passwd_hash
			) returning id`

		args = pgx.NamedArgs{
			"firstname":   dto.getFirstname(),
			"surname":     dto.getSurname(),
			"patronymic":  dto.getPatronymic(),
			"login":       dto.getLogin(),
			"note":        dto.getNote(),
			"status":      dto.getStatus(),
			"avatar_id":   dto.getAvatarID(),
			"expired_at":  dto.getExpiredAt(),
			"passwd_hash": dto.getPasswdHash(),
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
	dto *EmployeeUpdateModel,
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
			"position_id": dto.getPositionID(),
			"dept_id":     dto.getDeptID(),
			"worked_at":   dto.getWorkedAt(),
			"fired_at":    dto.getFiredAt(),
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
	opts *EmployeesSearchOpts,
) ([]*EmployeeModel, uint32, error) {
	var (
		// Total uint32 TODO пока кол-во пользователей меньше лимита, не делать подсчет кол-ва
		ls = make([]*EmployeeModel, 0)

		sq = `SELECT u.id, u.firstname, u.surname, u.note, u.status, u.avatar_id,
		     e.dept_id, e.position_id, e.worked_at, e.fired_at
		FROM       people.employees AS e
		LEFT JOIN  people.users     AS u
		    ON e.user_id = u.id AND (e.fired_at IS NULL OR e.fired_at > now())
		WHERE u.deleted = false
		  AND (u.expired_at IS NULL OR u.expired_at > now()) AND u.status = 'active'
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
