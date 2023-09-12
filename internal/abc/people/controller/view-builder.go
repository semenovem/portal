package people_controller

import (
	"fmt"
	"github.com/semenovem/portal/internal/abc/controller"
	"github.com/semenovem/portal/internal/abc/people/provider"
	"github.com/semenovem/portal/pkg/it"
)

func newUserPublicProfileView(u *people_provider.UserModel) *userPublicProfileView {
	r := &userPublicProfileView{
		ID:         u.ID(),
		Firstname:  u.Firstname(),
		Surname:    u.Surname(),
		Patronymic: u.Patronymic(),
	}
	if avatarID := u.AvatarID(); avatarID != 0 {
		// TODO сформировать реальную ссылку на скачивание аватара пользователя
		r.Avatar = fmt.Sprintf("https://asdas/asdasd/%d", avatarID)
	}

	return r
}

func newUserProfileView(u *people_provider.UserModel) *userProfileView {
	r := &userProfileView{
		userPublicProfileView: *newUserPublicProfileView(u),
		Note:                  u.Note(),
		ExpiredAt:             controller.TimeToString(u.ExpiredAt()),
		Status:                string(u.Status()),
		Roles:                 it.StringifyUserRoles(u.Roles()),
	}

	return r
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
		userPublicProfileView: *newUserPublicProfileView(&u.UserModel),
		WorkedAt:              u.WorkedAt(),
		FiredAt:               u.FiredAt(),
		Note:                  u.Note(),
		BossID:                0,
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
