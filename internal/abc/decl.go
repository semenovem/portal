package abc

import (
	"github.com/semenovem/portal/internal/abc/auth/action"
	"github.com/semenovem/portal/internal/abc/auth/provider"
	"github.com/semenovem/portal/internal/abc/media/action"
	"github.com/semenovem/portal/internal/abc/media/provider"
	"github.com/semenovem/portal/internal/abc/people/action"
	"github.com/semenovem/portal/internal/abc/people/provider"
	"github.com/semenovem/portal/internal/abc/store/action"
	"github.com/semenovem/portal/internal/abc/store/provider"
)

type Providers struct {
	checker
	Auth   *auth_provider.AuthProvider
	People *people_provider.PeopleProvider
	Store  *store_provider.StoreProvider
	Media  *media_provider.MediaProvider
}

type Actions struct {
	checker
	Auth   *auth_action.AuthAction
	People *people_action.PeopleAction
	Store  *store_action.StoreAction
	Media  *media_action.MediaAction
}

type checker struct {
}

func (c *checker) Valid() error {
	// TODO проверка что все значения установлены и нет nil
	return nil
}
