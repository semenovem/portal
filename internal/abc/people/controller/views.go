package people_controller

import "github.com/semenovem/portal/pkg/it"

func newEmployeePublicProfileViews(ls []*it.EmployeeProfile) []*employeeProfileView {
	a := make([]*publicEmployeeView, len(ls))
	for i := range ls {
		a[i] = newEmployeePublicProfileView(ls[i])
	}

	return nil
}
