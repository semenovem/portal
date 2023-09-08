package people_action

import (
	"context"
	"github.com/semenovem/portal/internal/abc/people/dto"
)

func (a *PeopleAction) UpdateEmployee(
	ctx context.Context,
	thisUserID, userID uint32,
	dto *people_dto.UserDTO,
) (*people_dto.UserProcessingErrDTO, error) {

	//people_dto.UserProcessingErrDTO

	return nil, nil
}
