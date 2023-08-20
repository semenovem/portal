package people_provider

import (
	"context"
	"github.com/semenovem/portal/pkg/it"
)

func (p *PeopleProvider) GetPosition(ctx context.Context, positionID uint16) (*it.UserPosition, error) {
	var (
		sq  = `SELECT title, COALESCE(parent_id, 0) AS parent_id FROM people.positions WHERE id = $1`
		pos = it.UserPosition{
			ID: positionID,
		}
	)

	err := p.db.QueryRow(ctx, sq, positionID).Scan(&pos.Title, &pos.ParentID)
	if err != nil {
		p.logger.Named("GetPosition").With("positionID", positionID).DB(err)
		return nil, err
	}

	return &pos, err
}
