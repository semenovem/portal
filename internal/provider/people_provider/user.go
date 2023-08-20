package people_provider

import (
	"context"
	"fmt"
	"github.com/semenovem/portal/internal/provider"
	"github.com/semenovem/portal/pkg/it"
)

func (p *PeopleProvider) GetUserProfile(ctx context.Context, userID uint32) (*it.UserProfile, error) {
	var (
		sq = `SELECT u.id, u.status, u.roles, u.firstname, u.surname, u.avatar, u.note, u.position, em.position_id
       		FROM people.users* AS u
       		LEFT JOIN people.employees AS em ON em.id = u.id
       		WHERE u.id = $1 AND u.deleted = false`

		u = it.UserProfile{
			UserCore: it.UserCore{
				ID:     0,
				Status: "",
				Roles:  nil,
			},
			Avatar:       nil,
			FirstName:    "",
			Surname:      "",
			PositionName: "",
			Note:         "",
		}

		positionID uint32
	)

	err := p.db.QueryRow(ctx, sq, userID).
		Scan(&u.ID, &u.Status, &u.Roles, &u.FirstName, &u.Surname, &u.Avatar, &u.Note, &u.PositionName, &positionID)
	if err != nil {
		if !provider.IsNoRows(err) {
			p.logger.Named("GetUserProfile").Error(err.Error())
		}

		return nil, err
	}

	if positionID != 0 {

	}

	fmt.Printf(">>>>>>>>>>>>> %+v\n", u)

	return &u, nil
}
