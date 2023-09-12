package people_action

import (
	"context"
	"github.com/semenovem/portal/internal/abc/people/provider"
)

func (a *PeopleAction) UpdateEmployee(
	ctx context.Context,
	thisUserID, userID uint32,
	dto *people_provider.EmployeeUpdateModel,
) error {
	var (
		ll = a.logger.Named("UpdateEmployee").With("userID", userID)
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
