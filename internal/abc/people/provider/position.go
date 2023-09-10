package people_provider

import (
	"context"
	"github.com/lib/pq"
)

type PositionModel struct {
	id          uint16
	title       string
	description string
}

func (m *PositionModel) ID() uint16 {
	return m.id
}

func (m *PositionModel) Title() string {
	return m.title
}

func (m *PositionModel) Description() string {
	return m.description
}

func (p *PeopleProvider) GetPositionModel(ctx context.Context, positionID uint16) (*PositionModel, error) {
	var (
		sq  = `SELECT id, title, description FROM people.positions WHERE id = $1`
		pos PositionModel
	)

	if err := p.db.QueryRow(ctx, sq, positionID).Scan(
		&pos.id,
		&pos.title,
		&pos.description,
	); err != nil {
		p.logger.Named("GetPositionModel").With("positionID", positionID).DB(err)
		return nil, err
	}

	return &pos, nil
}

func (p *PeopleProvider) GetPositionMap(
	ctx context.Context,
	positionIDs []uint16,
) (map[uint16]*PositionModel, error) {
	positions, err := p.GetPositions(ctx, positionIDs)
	if err != nil {
		return nil, err
	}

	posMap := make(map[uint16]*PositionModel)
	for _, m := range positions {
		posMap[m.id] = m
	}

	return posMap, nil
}

func (p *PeopleProvider) GetPositions(ctx context.Context, positionIDs []uint16) ([]*PositionModel, error) {
	var (
		sq = `SELECT id, title, description FROM people.positions WHERE id = ANY ($1)`
		ls = make([]*PositionModel, 0)
	)

	rows, err := p.db.Query(ctx, sq, pq.Array(positionIDs))
	if err != nil {
		p.logger.Named("GetPositions").DB(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var m PositionModel

		if err = rows.Scan(
			&m.id,
			&m.title,
			&m.description,
		); err != nil {
			p.logger.Named("GetPositions.Scan").DB(err)
			return nil, err
		}

		ls = append(ls, &m)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ls, err
}
