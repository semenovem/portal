package auth_action

import (
	"context"
	"github.com/google/uuid"
	"github.com/semenovem/portal/internal/provider"
	"github.com/semenovem/portal/pkg/it"
)

// CreateOnetimeEntry создание одноразовой точки авторизации
func (a *AuthAction) CreateOnetimeEntry(ctx context.Context, userID uint32) (uuid.UUID, error) {
	ll := a.logger.Named("CreateOnetimeEntry")

	userAuth, err := a.peoplePvd.GetUserAuth(ctx, userID)
	if err != nil {
		ll.Named("CreateOnetimeEntry").Nested(err.Error())

		if provider.IsNoRows(err) {
			return uuid.Nil, errUserNoFound
		}
	}

	if err = a.canLogin(userAuth); err != nil {
		ll.Named("canLogin").Nested(err.Error())
		return uuid.Nil, err
	}

	entryID, err := a.authPvd.NewOnetimeEntry(ctx, userAuth.ID)
	if err != nil {
		ll.Named("NewOnetimeEntry").Nested(err.Error())
	}

	return entryID, nil
}

// LoginByOnetimeEntryID логин по одноразовой точке входа
func (a *AuthAction) LoginByOnetimeEntryID(ctx context.Context, entryID uuid.UUID) (*it.AuthSession, error) {
	ll := a.logger.Named("LoginByOnetimeEntryID")

	userID, err := a.authPvd.GetDelOnetimeEntry(ctx, entryID)
	if err != nil {
		ll.Named("GetDelOnetimeEntry").Nested(err.Error())

		if provider.IsNoRec(err) {
			return nil, errOnetimeEntryNotFound
		}

		return nil, err
	}

	user, err := a.peoplePvd.GetUserAuth(ctx, userID)
	if err != nil {
		ll.Named("GetUserAuth").Nested(err.Error())

		if provider.IsNoRows(err) {
			return nil, errUserNoFound
		}
	}

	session, err := a.newSession(ctx, user, uuid.Nil)
	if err != nil {
		ll.Named("newSession").Nested(err.Error())
		return nil, err
	}

	return session, nil
}
