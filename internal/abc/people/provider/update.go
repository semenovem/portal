package people_provider

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/semenovem/portal/internal/abc/people/dto"
	"github.com/semenovem/portal/internal/abc/provider"
	"github.com/semenovem/portal/internal/util"
	"github.com/semenovem/portal/pkg/throw"
	"strings"
)

func (p *PeopleProvider) UpdateEmployee(
	ctx context.Context,
	userID uint32,
	dto *people_dto.EmployeeDTO,
) error {
	tx, err := p.db.Begin(ctx)
	if err != nil {
		p.logger.Named("UpdateEmployee.Begin").DB(err)
		return err
	}

	defer func() {
		if err1 := tx.Rollback(ctx); err1 != nil && !errors.Is(err1, pgx.ErrTxClosed) {
			p.logger.Named("Rollback").DB(err1)
		}
	}()

	if err = p.updateUserTx(ctx, tx, userID, &dto.UserDTO); err != nil {
		return err
	}

	if err = p.updateEmployeeTx(ctx, tx, userID, dto); err != nil {
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		p.logger.Named("Commit").DB(err)
		return err
	}

	return nil
}

func (p *PeopleProvider) updateUserTx(
	ctx context.Context,
	tx pgx.Tx,
	userID uint32,
	dto *people_dto.UserDTO,
) error {
	var (
		sq   = `UPDATE people.users SET updated_at = now()` + ", "
		set  = make([]string, 0)
		args = pgx.NamedArgs{
			"user_id":   userID,
			"firstname": dto.Firstname,
			"surname":   dto.Surname,
			"status":    dto.Status,
		}
	)

	if dto.Firstname != nil {
		set = append(set, "firstname = @firstname")
	}
	if dto.Surname != nil {
		set = append(set, "surname = @surname")
	}
	if dto.Status != nil {
		set = append(set, "status = @status")
	}
	if dto.Note != nil {
		set = append(set, "note = @note")
		args["note"] = util.ZeroStrNil(dto.Note)
	}
	if dto.Roles != nil {
		set = append(set, "roles = @roles::people.roles_enum[]")
		args["roles"] = util.ZeroArrStrNil(dto.Roles)
	}
	if dto.AvatarID != nil {
		set = append(set, "avatar_id = @avatar_id")
		args["avatar_id"] = util.ZeroUint32Nil(dto.AvatarID)
	}
	if dto.ExpiredAt != nil {
		set = append(set, "expired_at = @expired_at")
		args["expired_at"] = util.ZeroTimeNil(dto.ExpiredAt)
	}
	if dto.Login != nil {
		set = append(set, "login = @login")
		args["login"] = util.ZeroStrNil(dto.Login)
	}
	if dto.PasswdHash != nil {
		set = append(set, "passwd_hash = @passwd_hash")
		args["passwd_hash"] = util.ZeroStrNil(dto.PasswdHash)
	}

	if len(set) == 0 {
		return nil
	}

	sq += strings.Join(set, ",") + " WHERE id = @user_id AND deleted = false"

	res, err := tx.Exec(ctx, sq, args)
	if err != nil && !provider.IsDuplicateKeyErr(err) {
		if provider.IsDuplicateKeyErr(err) {
			// TODO тут привязать к названию ограничения и сообщить конкретную ошибку
			return throw.Err400DuplicateLogin
		}

		p.logger.Named("UpdateUserTx").DB(err)
		return err
	}

	if res.RowsAffected() == 0 {
		return throw.Err404User
	}

	return nil
}

func (p *PeopleProvider) updateEmployeeTx(
	ctx context.Context,
	tx pgx.Tx,
	userID uint32,
	dto *people_dto.EmployeeDTO,
) error {
	var (
		sq    = `UPDATE people.employees SET`
		where = ` WHERE user_id = @user_id AND (select deleted from people.users where id = @user_id) = false`
		set   = make([]string, 0)
		args  = pgx.NamedArgs{
			"user_id":     userID,
			"position_id": dto.PositionID,
			"dept_id":     dto.DeptID,
			"worked_at":   dto.WorkedAt,
		}
	)

	if dto.PositionID != nil {
		set = append(set, "position_id = @position_id")
	}
	if dto.DeptID != nil {
		set = append(set, "dept_id = @dept_id")
	}
	if dto.WorkedAt != nil {
		set = append(set, "worked_at = @worked_at")
	}
	if dto.FiredAt != nil {
		set = append(set, "fired_at = @fired_at")
		args["fired_at"] = util.ZeroTimeNil(dto.FiredAt)
	}

	if len(set) == 0 {
		return nil
	}

	sq += " " + strings.Join(set, ",") + where

	res, err := tx.Exec(ctx, sq, args)
	if err != nil {
		if provider.IsDuplicateKeyErr(err) {
			// TODO тут привязать к названию ограничения и сообщить конкретную ошибку
			return throw.Err400DuplicateLogin
		}

		if provider.IsCheckErr(err) {
			// TODO тут привязать к названию ограничения и сообщить конкретную ошибку
			return throw.NewBadRequestWithTargetErr(throw.Err400FiredBehind, err)
		}

		p.logger.Named("UpdateUserTx").DB(err)
		return err
	}

	if res.RowsAffected() == 0 {
		return throw.Err404User
	}

	return nil
}
