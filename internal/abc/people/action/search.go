package people_action

import (
	"context"
	people_dto "github.com/semenovem/portal/internal/abc/people/dto"
)

func (a *PeopleAction) EmployeeHandbook(
	ctx context.Context,
	opts *people_dto.EmployeesSearchOpts,
) (*people_dto.EmployeesSearchResult, error) {
	var (
		ll = a.logger.Named("EmployeeHandbook")
	)

	result, err := a.peoplePvd.SearchEmployees(ctx, opts)
	if err != nil {
		ll.Named("peoplePvd.SearchEmployees").Nested(err)
		return nil, err
	}

	return result, nil
}
