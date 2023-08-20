package people_provider

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/semenovem/portal/pkg"
)

const (
	userIDIsEmpty    = "userID is empty"
	userLoginIsEmpty = "login is empty"
)

type PeopleProvider struct {
	db     *pgxpool.Pool
	logger pkg.Logger
}

func NewPeoplePvd(db *pgxpool.Pool, logger pkg.Logger) *PeopleProvider {
	return &PeopleProvider{
		db:     db,
		logger: logger.Named("peoplePvd"),
	}
}
