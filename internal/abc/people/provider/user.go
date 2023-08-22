package people_provider

import (
	"context"
	"github.com/semenovem/portal/internal/abc/provider"
	"github.com/semenovem/portal/pkg/it"
)

func (p *PeopleProvider) GetUserProfile(ctx context.Context, userID uint32) (*it.UserProfile, error) {
	var (
		sq = `SELECT u.id,
					u.status,
					u.roles,
					u.firstname,
					u.surname,
					u.avatar,
					u.note,
					u.position,
					COALESCE(em.position_id, 0)
       		FROM people.users* AS u
       		LEFT JOIN people.employees AS em ON em.user_id = u.id
       		WHERE u.id = $1 AND u.deleted = false`

		u          = it.UserProfile{}
		positionID uint16
	)

	err := p.db.QueryRow(ctx, sq, userID).
		Scan(
			&u.ID,
			&u.Status,
			&u.Roles,
			&u.FirstName,
			&u.Surname,
			&u.Avatar,
			&u.Note,
			&u.PositionTitle,
			&positionID,
		)
	if err != nil {
		if !provider.IsNoRows(err) {
			p.logger.Named("GetUserProfile").DB(err)
		}

		return nil, err
	}

	if positionID != 0 {
		pos, err := p.GetPosition(ctx, positionID)
		if err != nil {
			if !provider.IsNoRows(err) {
				p.logger.Named("GetUserProfile.GetPosition").DB(err)
			}
			return nil, err
		}

		u.PositionTitle = pos.Title
	}

	return &u, nil
}
