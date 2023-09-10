package people_provider

import (
	"context"
	"github.com/lib/pq"
)

type DeptModel struct {
	id          uint16
	title       string
	description string
	parentID    *uint16
}

func (m *DeptModel) ID() uint16 {
	return m.id
}

func (m *DeptModel) Title() string {
	return m.title
}

func (m *DeptModel) Description() string {
	return m.description
}

func (m *DeptModel) ParentID() uint16 {
	if m.parentID == nil {
		return 0
	}
	return *m.parentID
}

func (p *PeopleProvider) GetDeptMap(ctx context.Context, deptIDs []uint16) (map[uint16]*DeptModel, error) {
	depts, err := p.GetDepts(ctx, deptIDs)
	if err != nil {
		return nil, err
	}

	deptMap := make(map[uint16]*DeptModel)
	for _, dept := range depts {
		deptMap[dept.id] = dept
	}

	return deptMap, nil
}

func (p *PeopleProvider) GetDepts(ctx context.Context, deptIDs []uint16) ([]*DeptModel, error) {
	var (
		sq = `SELECT id, title, description, parent_id
		FROM people.departments
		WHERE id = ANY ($1)`

		ls = make([]*DeptModel, 0)
	)

	rows, err := p.db.Query(ctx, sq, pq.Array(deptIDs))
	if err != nil {
		p.logger.Named("GetDepts.Query").DB(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var m DeptModel

		if err = rows.Scan(
			&m.id,
			&m.title,
			&m.description,
			&m.parentID,
		); err != nil {
			p.logger.Named("GetDepts.Scan").DB(err)
			return nil, err
		}

		ls = append(ls, &m)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ls, err
}

func (p *PeopleProvider) GetBosses(ctx context.Context) ([]*UserModel, error) {
	return nil, nil
}
