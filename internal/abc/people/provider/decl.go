package people_provider

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/throw"
)

const (
	userIDIsEmpty    = "userID is empty"
	userLoginIsEmpty = "login is empty"
)

const (
	firedBeforeStartConstraint = "users_fired_before_work_constraint"
	usersLoginUniqueConstraint = "users_login_unique_constraint"
)

type PeopleProvider struct {
	db     *pgxpool.Pool
	logger pkg.Logger
}

func New(db *pgxpool.Pool, logger pkg.Logger) *PeopleProvider {
	return &PeopleProvider{
		db:     db,
		logger: logger.Named("PeopleProvider"),
	}
}

func constraintErr(name string, err error) error {
	switch name {
	case firedBeforeStartConstraint:
		return throw.NewWithTargetErr(throw.Err400FiredBeforeStart, err)
	case usersLoginUniqueConstraint:
		return throw.NewWithTargetErr(throw.Err400DuplicateLogin, err)
	}

	return throw.NewBadRequestErr(err.Error())
}
