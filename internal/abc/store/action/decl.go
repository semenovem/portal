package store_action

import (
	"context"
	"github.com/semenovem/portal/internal/abc/action"
	"github.com/semenovem/portal/internal/abc/provider"
	"github.com/semenovem/portal/internal/abc/store/provider"
	"github.com/semenovem/portal/pkg"
)

type StoreAction struct {
	logger   pkg.Logger
	storePvd *store_provider.StoreProvider
}

func New(logger pkg.Logger, storePvd *store_provider.StoreProvider) *StoreAction {
	return &StoreAction{
		logger:   logger.Named("StoreAction"),
		storePvd: storePvd,
	}
}

func (a *StoreAction) Load(ctx context.Context, thisUserID uint32, path string) (string, error) {
	ll := a.logger.Named("Load").With("thisUserID", thisUserID).With("path", path)

	payload, err := a.storePvd.LoadArbitraryData(ctx, thisUserID, path)
	if err != nil {
		ll.Named("LoadArbitraryData").Nested(err.Error())

		if provider.IsNoRec(err) {
			return "", action.ErrNotFound
		}

		return "", err
	}

	return payload, nil
}

func (a *StoreAction) Store(ctx context.Context, thisUserID uint32, path, payload string) error {
	ll := a.logger.Named("Load").With("thisUserID", thisUserID).With("path", path)

	err := a.storePvd.StoreArbitraryData(ctx, thisUserID, path, payload)
	if err != nil {
		ll.Named("StoreArbitraryData").Nested(err.Error())
	}

	return err
}

func (a *StoreAction) Delete(ctx context.Context, thisUserID uint32, path string) error {
	ll := a.logger.Named("Load").With("thisUserID", thisUserID).With("path", path)

	err := a.storePvd.DeleteArbitraryData(ctx, thisUserID, path)
	if err != nil {
		ll.Named("DeleteArbitraryData").Nested(err.Error())

		if provider.IsNoRec(err) {
			return action.ErrNotFound
		}
	}

	return err
}
