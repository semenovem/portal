package people_provider

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/semenovem/portal/internal/abc/people/dto"
	"github.com/semenovem/portal/internal/abc/provider"
	"github.com/semenovem/portal/pkg/throw"
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
		if provider.IsDuplicateKeyErr(err) {
			return 0, throw.Err400DuplicateLogin
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
		ON CONFLICT (user) DO UPDATE SET
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
		if provider.IsCheckErr(err) {
			return throw.Err400FiredBehind
		}

		if provider.IsConstrainForeignKeyErr(err) {
			return throw.NewBadRequestErr(err.Error())
		}

		p.logger.Named("createEmployeeTx").DB(err)
	}

	return nil
}
