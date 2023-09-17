package people_provider

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/semenovem/portal/internal/abc/people"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/throw"
	"html"
	"strings"
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

	return throw.NewBadRequestErrf(err.Error())
}

// EmployeesSearchOpts параметры поиска сотрудников
type EmployeesSearchOpts struct {
	Limit  uint32
	Offset uint32
	//Expired         bool       // Включая у кого истек срок действия
	//ExpiredAtAfter  *time.Time // Включая истекшие записи после
	//ExpiredAtBefore *time.Time // Включая у кого истек срок действия
	Fired    bool // Включая уволенных
	Statuses []people.UserStatus
}

func escapeHTML(s *string) *string {
	if s == nil {
		return nil
	}

	ss := html.EscapeString(strings.TrimSpace(*s))
	return &ss
}
