package people_action

import (
	"context"
	"github.com/semenovem/portal/internal/abc/provider"
	"github.com/semenovem/portal/pkg/it"
	"github.com/semenovem/portal/pkg/throw"
	"strings"
)

func (a *PeopleAction) CreateUser(
	ctx context.Context,
	thisUserID uint32,
	dto *CreateUserDTO,
) (*it.UserProfile, error) {
	ll := a.logger.Named("CreateUser")

	// TODO Должна быть проверка на право удаления пользователя
	// ----
	// ----
	// ----

	userID, err := a.peoplePvd.CreateUser(ctx, dto.toProviderUserModel())
	if err != nil {
		ll.Named("peoplePvd.CreateUser").Nested(err)

		if provider.IsDuplicateKeyError(err) {
			if strings.Contains(err.Error(), "users_login_key") {
				return nil, throw.ErrBadRequestDuplicateLogin
			}

			return nil, throw.NewBadRequestErr(err.Error())
		}

		return nil, err
	}

	return a.getUserProfile(ctx, userID)
}

func (a *PeopleAction) DeleteUser(
	ctx context.Context,
	thisUserID, userID uint32,
) error {
	//ll := a.logger.Named("DeleteUser")

	// TODO Должна быть проверка на право удаления пользователя

	return nil
}
