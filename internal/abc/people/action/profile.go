package people_action

import (
	"context"
	people_provider "github.com/semenovem/portal/internal/abc/people/provider"
	"github.com/semenovem/portal/internal/abc/provider"
	"github.com/semenovem/portal/pkg/throw"
)

func (a *PeopleAction) GetUserModel(
	ctx context.Context,
	thisUserID, userID uint32,
) (*people_provider.UserModel, error) {
	ll := a.logger.Func(ctx, "getUserProfile")

	// TODO проверка права на просмотр данных пользователя

	profile, err := a.peoplePvd.GetUserModel(ctx, userID)
	if err != nil {
		ll.Named("peoplePvd.GetUserModel").Nested(err)
		return nil, err
	}

	return profile, err
}

func (a *PeopleAction) GetEmployeeModel(ctx context.Context, userID uint32) (*people_provider.EmployeeModel, error) {
	ll := a.logger.Func(ctx, "GetEmployeeModel")

	profile, err := a.peoplePvd.GetEmployeeModel(ctx, userID)
	if err != nil {
		ll.Named("GetUserProfile").Nested(err)

		if provider.IsNoRow(err) {
			return nil, throw.Err404User
		}

		return nil, err
	}

	return profile, err
}
