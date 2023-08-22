package people_action

import (
	"context"
)

func (a *PeopleAction) DeleteUser(
	ctx context.Context,
	thisUserID, userID uint32,
) error {
	//ll := a.logger.Named("DeleteUser")

	// TODO Должна быть проверка на право удаления пользователя

	return nil
}
