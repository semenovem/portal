package people_provider

import (
	"context"
	"github.com/lib/pq"
	"github.com/semenovem/portal/pkg/it"
)

type PositionModel struct {
	id          uint16
	deptID      uint16
	title       string
	description string
	parentID    *uint16
	deleted     bool
}

func (m *PositionModel) ID() uint16 {
	return m.id
}

func (m *PositionModel) DeptID() uint16 {
	return m.deptID
}

func (m *PositionModel) Title() string {
	return m.title
}

func (m *PositionModel) Description() string {
	return m.description
}

func (m *PositionModel) ParentID() uint16 {
	if m.parentID == nil {
		return 0
	}
	return *m.parentID
}

func (p *PeopleProvider) GetPositionModel(ctx context.Context, positionID uint16) (*it.UserPosition, error) {
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

func (p *PeopleProvider) GetPositionModels(
	ctx context.Context,
	positionIDs []uint16,
) ([]*PositionModel, error) {
	var (
		sq = `SELECT id, dept_id, title, description, parent_id
		FROM people.positions
		WHERE id = ANY ($1)`

		ls = make([]*PositionModel, 0)
	)

	rows, err := p.db.Query(ctx, sq, pq.Array(positionIDs))
	if err != nil {
		p.logger.Named("GetPositionModels").DB(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var m PositionModel

		if err = rows.Scan(
			&m.id,
			&m.deptID,
			&m.title,
			&m.description,
			&m.parentID,
		); err != nil {
			p.logger.Named("GetPositionModels.Scan").DB(err)
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ls, err
}
