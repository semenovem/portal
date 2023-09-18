package people_provider

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/semenovem/portal/internal/abc/provider"
	"github.com/semenovem/portal/pkg/throw"
	"strings"
)

func (p *PeopleProvider) UpdateEmployee(
	ctx context.Context,
	userID uint32,
	dto *EmployeeCreateModel,
) error {
	ll := p.logger.Func(ctx, "UpdateEmployee")

	tx, err := p.db.Begin(ctx)
	if err != nil {
		ll.Named("Begin").DB(err)
		return err
	}

	defer func() {
		if err1 := tx.Rollback(ctx); err1 != nil && !errors.Is(err1, pgx.ErrTxClosed) {
			ll.Named("Rollback").DB(err1)
		}
	}()

	if err = p.updateUserTx(ctx, tx, userID, &dto.UserCreateModel); err != nil {
		return err
	}

	if err = p.updateEmployeeTx(ctx, tx, userID, dto); err != nil {
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		ll.Named("Commit").DB(err)
		return err
	}

	return nil
}

func (p *PeopleProvider) updateUserTx(
	ctx context.Context,
	tx pgx.Tx,
	userID uint32,
	dto *UserCreateModel,
) error {
	var (
		ll    = p.logger.Func(ctx, "updateUserTx")
		sq    = `UPDATE people.users SET updated_at = now()` + ", "
		where = " WHERE id = @user_id AND deleted = false"
		set   = make([]string, 0)
		args  = pgx.NamedArgs{
			"user_id": userID,
		}
	)

	if dto.Firstname != nil {
		set = append(set, "firstname = LOWER(@firstname)")
		args["firstname"] = dto.getFirstname()
	}
	if dto.Surname != nil {
		set = append(set, "surname = LOWER(@surname)")
		args["surname"] = dto.getSurname()
	}
	if dto.Patronymic != nil {
		set = append(set, "surname = LOWER(@patronymic)")
		args["patronymic"] = dto.getPatronymic()
	}
	if dto.Status != nil {
		set = append(set, "status = @status")
		args["status"] = dto.getStatus()
	}
	if dto.Note != nil {
		set = append(set, "note = @note")
		args["note"] = dto.getNote()
	}
	if dto.AvatarID != nil {
		set = append(set, "avatar_id = @avatar_id")
		args["avatar_id"] = dto.getAvatarID()
	}
	if dto.ExpiredAt != nil {
		set = append(set, "expired_at = @expired_at")
		args["expired_at"] = dto.getExpiredAt()
	}
	if dto.Login != nil {
		set = append(set, "login = @login")
		args["login"] = dto.getLogin()
	}
	if dto.PasswdHash != nil {
		set = append(set, "passwd_hash = @passwd_hash")
		args["passwd_hash"] = dto.getPasswdHash()
	}

	if len(set) == 0 {
		return nil
	}

	sq += strings.Join(set, ",") + where

	res, err := tx.Exec(ctx, sq, args)
	if err != nil && !provider.IsDuplicateKeyErr(err) {
		if n, ok := provider.ConstraintErr(err); ok {
			return constraintErr(n, err)
		}

		ll.Named("Exec").DB(err)
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
	dto *EmployeeCreateModel,
) error {
	var (
		ll    = p.logger.Func(ctx, "updateEmployeeTx")
		sq    = `UPDATE people.employees SET`
		where = ` WHERE user_id = @user_id AND (select deleted from people.users where id = @user_id) = false`
		set   = make([]string, 0)
		args  = pgx.NamedArgs{
			"user_id": userID,
		}
	)

	if dto.PositionID != nil {
		set = append(set, "position_id = @position_id")
		args["position_id"] = dto.getPositionID()
	}
	if dto.DeptID != nil {
		set = append(set, "dept_id = @dept_id")
		args["dept_id"] = dto.getDeptID()
	}
	if dto.WorkedAt != nil {
		set = append(set, "worked_at = @worked_at")
		args["worked_at"] = dto.getWorkedAt()
	}
	if dto.FiredAt != nil {
		set = append(set, "fired_at = @fired_at")
		args["fired_at"] = dto.getFiredAt()
	}

	if len(set) == 0 {
		return nil
	}

	sq += " " + strings.Join(set, ",") + where

	res, err := tx.Exec(ctx, sq, args)
	if err != nil {
		if n, ok := provider.ConstraintErr(err); ok {
			return constraintErr(n, err)
		}

		ll.Named("UpdateUserTx").DB(err)
		return err
	}

	if res.RowsAffected() == 0 {
		return throw.Err404User
	}

	return nil
}
