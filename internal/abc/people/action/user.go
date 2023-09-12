package people_action

import (
	"context"
	people_provider "github.com/semenovem/portal/internal/abc/people/provider"
	"github.com/semenovem/portal/pkg/it"
)

func (a *PeopleAction) CheckLoginName(
	ctx context.Context,
	thisUserID uint32,
	loginName string,
) (loginNameIsFree bool, err error) {
	ll := a.logger.Named("CheckLoginName")

	// TODO Должна быть проверка на право проверки свободного логина пользователя
	// Требуется право создания нового пользователя
	// Требуется право менять логин у пользователя

	if err = it.ValidateUserLogin(loginName); err != nil {
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

func (a *PeopleAction) CreateEmployee(
	ctx context.Context,
	thisUserID uint32,
	dto *people_provider.EmployeeUpdateModel,
) (userID uint32, err error) {
	ll := a.logger.Named("CreateUser")

	// TODO Должна быть проверка на право создания пользователя
	// ----
	// ----
	// ----

	userID, err = a.peoplePvd.CreateEmployee(ctx, dto)
	if err != nil {
		ll.Named("peoplePvd.CreateUser").Nested(err)
		return
	}

	return
}

func (a *PeopleAction) DeleteUser(
	ctx context.Context,
	thisUserID, userID uint32,
) error {
	ll := a.logger.Named("DeleteUser")

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
	ll := a.logger.Named("UndeleteUser")

	// TODO Должна быть проверка на право восстановления пользователя

	if err := a.peoplePvd.UndeleteUser(ctx, userID); err != nil {
		ll.Named("peoplePvd.UndeleteUser").Nested(err)
		return err
	}

	return nil
}
