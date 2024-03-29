package people_action

import (
	"context"
	"github.com/semenovem/portal/internal/abc/people"
)

func (a *PeopleAction) CheckLoginName(
	ctx context.Context,
	thisUserID uint32,
	loginName string,
) (loginNameIsFree bool, err error) {
	ll := a.logger.Func(ctx, "CheckLoginName")

	// TODO Должна быть проверка на право проверки свободного логина пользователя
	// Требуется право создания нового пользователя
	// Требуется право менять логин у пользователя

	if err = people.ValidateUserLogin(loginName); err != nil {
		ll.Named("ValidateUserLogin").BadRequest(err)
		return
	}

	exists, err := a.peoplePvd.ExistsLoginName(ctx, loginName)
	if err != nil {
		ll.Named("peoplePvd.ExistsLoginName").Nested(err)
		return false, err
	}

	return exists, nil
}

func (a *PeopleAction) DeleteUser(
	ctx context.Context,
	thisUserID, userID uint32,
) error {
	ll := a.logger.Func(ctx, "DeleteUser")

	// TODO Должна быть проверка на право удаления пользователя

	if err := a.peoplePvd.DeleteUser(ctx, userID); err != nil {
		ll.Named("peoplePvd.DeleteUser").Nested(err)
		return err
	}

	return nil
}

func (a *PeopleAction) UndeleteUser(
	ctx context.Context,
	thisUserID, userID uint32,
) error {
	ll := a.logger.Func(ctx, "UndeleteUser")

	// TODO Должна быть проверка на право восстановления пользователя

	if err := a.peoplePvd.UndeleteUser(ctx, userID); err != nil {
		ll.Named("peoplePvd.UndeleteUser").Nested(err)
		return err
	}

	return nil
}
