package people_action

import (
	"context"
	"github.com/semenovem/portal/internal/abc/people"
)

func (a *PeopleAction) PublicHandbook(ctx context.Context) (*people.EmployeesSearchResult, error) {
	var (
		ll = a.logger.Named("PublicHandbook")
	)

	opts := &people.EmployeesSearchOpts{
		Limit:  0,
		Offset: 0,
	}

	result, err := a.peoplePvd.SearchEmployees(ctx, opts)
	if err != nil {
		ll.Named("peoplePvd.SearchEmployees").Nested(err)
		return nil, err
	}

	return result, nil
}
