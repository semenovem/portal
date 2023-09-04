package apiapp

import (
	"github.com/semenovem/portal/internal/abc/auth/action"
	"github.com/semenovem/portal/internal/abc/media/action"
	"github.com/semenovem/portal/internal/abc/people/action"
	"github.com/semenovem/portal/internal/abc/store/action"
)

func (app *appAPI) initActions() {
	app.actions.auth = auth_action.New(
		app.logger,
		app.userPasswdAuth,
		app.providers.auth,
		app.providers.people,
	)

	app.actions.people = people_action.New(
		app.logger,
		app.userPasswdAuth,
		app.providers.people,
	)

	app.actions.store = store_action.New(
		app.logger,
		app.providers.store,
	)

	app.actions.media = media_action.New(
		app.logger,
		app.s3,
		app.providers.media,
	)
}
