package apiapp

import (
	auth_action "github.com/semenovem/portal/internal/abc/auth/action"
	"github.com/semenovem/portal/internal/abc/auth/provider"
	media_action "github.com/semenovem/portal/internal/abc/media/action"
	"github.com/semenovem/portal/internal/abc/media/provider"
	people_action "github.com/semenovem/portal/internal/abc/people/action"
	"github.com/semenovem/portal/internal/abc/people/provider"
	store_action "github.com/semenovem/portal/internal/abc/store/action"
	"github.com/semenovem/portal/internal/abc/store/provider"
	"time"
)

func (app *appAPI) initDomains() error {
	ll := app.logger.Named("initDomains")

	// --------------------- store ---------------------
	app.providers.Store = store_provider.New(
		app.logger,
		app.db,
		app.redis,
		time.Minute,
		time.Minute,
	)

	app.actions.Store = store_action.New(
		app.logger,
		app.providers.Store,
	)

	// --------------------- media ---------------------
	app.providers.Media = media_provider.New(
		app.logger,
		app.db,
		app.redis,
	)

	app.actions.Media = media_action.New(
		app.logger,
		app.s3,
		app.providers.Media,
	)

	// --------------------- people ---------------------
	app.providers.People = people_provider.New(
		app.db,
		app.logger,
	)

	app.actions.People = people_action.New(
		app.logger,
		app.loginPasswdAuth,
		app.providers.People,
	)

	// --------------------- auth ---------------------
	app.providers.Auth = auth_provider.New(
		app.logger,
		app.db,
		app.redis,
		app.config,
	)

	app.actions.Auth = auth_action.New(
		app.logger,
		app.providers.Auth,
		app.providers.People,
	)

	if err := app.providers.Valid(); err != nil {
		ll.Named("providers.Valid").Errorf("providers check failed: %s", err.Error())
		return err
	}

	if err := app.actions.Valid(); err != nil {
		ll.Named("actions.Valid").Errorf("actions check failed: %s", err.Error())
		return err
	}

	return nil
}
