package people_entity

const (
	EmployeeFieldFirstname  EmployeeField = "firstname"
	EmployeeFieldSurname    EmployeeField = "surname"
	EmployeeFieldNote       EmployeeField = "note"
	EmployeeFieldStatus     EmployeeField = "status"
	EmployeeFieldRoles      EmployeeField = "roles"
	EmployeeFieldAvatarID   EmployeeField = "avatar_id"
	EmployeeFieldExpiredAt  EmployeeField = "expired_at"
	EmployeeFieldPositionID EmployeeField = "position_id"
	EmployeeFieldDeptID     EmployeeField = "dept_id"
	EmployeeFieldWorkedAt   EmployeeField = "worked_at"
	EmployeeFieldFiredAt    EmployeeField = "fired_at"
)

type EmployeeField string
