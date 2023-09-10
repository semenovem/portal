package people_controller

import (
	"fmt"
	"github.com/semenovem/portal/internal/abc/people/provider"
	"github.com/semenovem/portal/pkg/it"
)

func newEmployeePublicProfileViews(ls []*it.EmployeeProfile) []*employeeProfileView {
	a := make([]*publicEmployeeView, len(ls))
	for i := range ls {
		a[i] = newEmployeePublicProfileView(ls[i])
	}

	return nil
}

func newEmployeeProfileViews(
	employees []*people_provider.EmployeeModel,
	deptMap map[uint16]*people_provider.DeptModel,
	positionMap map[uint16]*people_provider.PositionModel,
	// добавить руководителя
) []*employeeProfileView {
	ls := make([]*employeeProfileView, 0, len(employees))

	for _, u := range employees {
		var (
			dept = deptMap[u.DeptID()]
			post = positionMap[u.PositionID()]
			view = newEmployeeProfileView(u, dept, post)
		)

		ls = append(ls, view)
	}

	return ls
}

func newEmployeeProfileView(
	u *people_provider.EmployeeModel,
	dept *people_provider.DeptModel,
	position *people_provider.PositionModel,
) *employeeProfileView {
	obj := employeeProfileView{
		userPublicProfileView: userPublicProfileView{
			ID:        u.ID(),
			Firstname: u.Firstname(),
			Surname:   u.Surname(),
			Avatar:    fmt.Sprintf("https://asdas/asdasd/%d", u.AvatarID()),
		},
		WorkedAt: u.WorkedAt(),
		FiredAt:  u.FiredAt(),
		Note:     u.Note(),
		BossID:   0,
	}

	if dept != nil {
		obj.DeptName = dept.Title()
	} else {
		obj.DeptName = "unknown"
	}

	if position != nil {
		obj.PositionName = position.Title()
	} else {
		obj.PositionName = "unknown"
	}

	return &obj
}
