package people_action

import (
	"context"
	"github.com/semenovem/portal/internal/abc/people/provider"
)

func (a *PeopleAction) CreateEmployee(
	ctx context.Context,
	thisUserID uint32,
	dto *people_provider.EmployeeCreateModel,
) (userID uint32, err error) {
	ll := a.logger.Func(ctx, "CreateUser")

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

func (a *PeopleAction) UpdateEmployee(
	ctx context.Context,
	thisUserID, userID uint32,
	dto *people_provider.EmployeeCreateModel,
) error {
	var (
		ll = a.logger.Func(ctx, "UpdateEmployee").With("userID", userID)
	)

	// TODO проверка права редактирования пользователя
	// Проверка права приема/увольнения пользователя

	// валидация

	err := a.peoplePvd.UpdateEmployee(ctx, userID, dto)
	if err != nil {
		ll.Named("peoplePvd.UpdateEmployee").Nested(err)
		return err
	}

	return nil
}
