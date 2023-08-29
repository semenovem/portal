package people_action

import (
	"context"
	"github.com/semenovem/portal/internal/abc/action"
	"github.com/semenovem/portal/internal/abc/provider"
	"github.com/semenovem/portal/pkg/it"
)

func (a *PeopleAction) GetUserProfile(
	ctx context.Context,
	thisUserID, userID uint32,
) (*it.UserProfile, error) {
	ll := a.logger.Named("GetUserProfile")

	// TODO тут делать проверку права на просмотр данных пользователя

	profile, err := a.peoplePvd.GetUserProfile(ctx, userID)
	if err != nil {
		ll.Named("GetUserProfile").Nested(err)

		if provider.IsNoRows(err) {
			return nil, action.ErrNotFound
		}

		return nil, err
	}

	return profile, err
}
