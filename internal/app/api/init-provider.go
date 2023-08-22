package apiapp

import (
	"github.com/semenovem/portal/internal/abc/auth/provider"
	"github.com/semenovem/portal/internal/abc/media/provider"
	"github.com/semenovem/portal/internal/abc/people/provider"
	"github.com/semenovem/portal/internal/abc/store/provider"
	"time"
)

func (app *appAPI) initProviders() {
	app.providers.auth = auth_provider.New(
		app.logger,
		app.db,
		app.redis,
		time.Minute*time.Duration(app.config.JWT.AccessTokenLifetimeMin),
		time.Minute*time.Duration(app.config.Auth.OnetimeEntryLifetimeMin),
	)

	app.providers.people = people_provider.New(
		app.db,
		app.logger,
	)

	app.providers.store = store_provider.New(
		app.logger,
		app.db,
		app.redis,
		time.Minute,
		time.Minute,
	)

	app.providers.media = media_provider.New(
		app.logger,
		app.db,
		app.redis,
	)
}
