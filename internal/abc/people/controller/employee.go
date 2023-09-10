package people_controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/internal/abc/controller"
	people_dto "github.com/semenovem/portal/internal/abc/people/dto"
	"github.com/semenovem/portal/internal/audit"
	_ "github.com/semenovem/portal/pkg/fail"
	"net/http"
)

// EmployeeHandbook docs
//
//	@Summary		Справочник сотрудников
//	@Description	Доступен в локальной сети без авторизации
//	@Description
//	@Produce	json
//	@Success	200	{object}	employeeHandbookResponse
//	@Failure	400	{object}	fail.Response
//	@Router		/people/employee/handbook [GET]
//	@Tags		people
func (cnt *Controller) EmployeeHandbook(c echo.Context) error {
	var (
		ll  = cnt.logger.Named("EmployeeHandbook")
		ctx = c.Request().Context()
	)

	opts := &people_dto.EmployeesSearchOpts{
		Limit:  100,
		Offset: 0,
	}

	result, err := cnt.peopleAct.EmployeeHandbook(ctx, opts)
	if err != nil {
		ll.Named("peopleAct.EmployeeHandbook").Nested(err)
		return cnt.com.Response(c, ll, err)
	}

	fmt.Println(">>>>>>> Employees   >> ", result.Employees)
	fmt.Println(">>>>>>> DeptMap     >> ", result.DeptMap)
	fmt.Println(">>>>>>> PositionMap >> ", result.PositionMap)
	fmt.Println(">>>>>>> UserBossMap >> ", result.UserBossMap)

	ll.Debug("success")

	response := employeeHandbookResponse{
		Total:     result.Total,
		Employees: newEmployeePublicProfileViews(nil),
	}

	return c.JSON(http.StatusOK, response)
}

// CreateEmployee docs
//
//	@Summary		Создает нового сотрудника
//	@Description	`expired_at, worked_at, fired_at` в формате `2001-03-24T00:00:00Z`
//	@Description	введенный login нужно проверить, что он допустим `/people/free-login/:login_name`
//	@Description
//	@Produce	json
//	@Param		payload	body		createUserForm	true	"данные создаваемого пользователя"
//	@Success	200		{object}	userCreateResponse
//	@Failure	400		{object}	fail.Response
//	@Router		/people/employee [POST]
//	@Tags		people/employee
//	@Security	ApiKeyAuth
func (cnt *Controller) CreateEmployee(c echo.Context) error {
	var (
		thisUserID = controller.ExtractThisUserID(c)
		ll         = cnt.logger.Named("CreateEmployee").With("thisUserID", thisUserID)
		ctx        = c.Request().Context()
		form       = new(createUserForm)
	)

	if err := cnt.com.ExtractForm(c, ll, form); err != nil {
		return err.Err
	}

	passwdHash := cnt.userPasswdAuth.Hashing(form.Passwd)

	dto := &people_dto.EmployeeDTO{
		UserDTO: people_dto.UserDTO{
			Firstname:  form.getFirstname(),
			Surname:    form.getSurname(),
			Note:       form.getNote(),
			Status:     &form.Status,
			Roles:      &form.Roles,
			AvatarID:   &form.AvatarID,
			ExpiredAt:  form.getExpiredAt(),
			Login:      form.getLogin(),
			PasswdHash: &passwdHash,
		},
		PositionID: &form.PositionID,
		DeptID:     &form.DeptID,
		WorkedAt:   form.getWorkedAt(),
		FiredAt:    form.getFiredAt(),
	}

	userID, err := cnt.peopleAct.CreateEmployee(ctx, thisUserID, dto)
	if err != nil {
		ll = ll.Named("peopleAct.CreateEmployee")
		return cnt.com.Response(c, ll, err)
	}

	cnt.audit.Oper(thisUserID, audit.EntityEmployee, audit.Create, audit.P{
		"userID": userID,
	})

	ll.With("userID", userID).Debug("user created")

	return c.JSON(http.StatusOK, userCreateResponse{UserID: userID})
}

// UpdateEmployee docs
//
//	@Summary		Обновление данных пользователя
//	@Description	`expired_at, worked_at, fired_at` в формате `2001-03-24T00:00:00Z`
//	@Description
//	@Description
//	@Description
//	@Description
//	@Produce	json
//	@Param		user_id	path		string	true	"id пользователя"
//	@Success	200		{object}	employeeUpdateForm
//	@Failure	400		{object}	fail.Response
//	@Router		/people/employee/:user_id [PATCH]
//	@Tags		people/employee
//	@Security	ApiKeyAuth
func (cnt *Controller) UpdateEmployee(c echo.Context) error {
	var (
		thisUserID = controller.ExtractThisUserID(c)
		ll         = cnt.logger.Named("UpdateEmployee").With("thisUserID", thisUserID)
		ctx        = c.Request().Context()
		form       = new(employeeUpdateForm)
	)

	if err := cnt.com.ExtractForm(c, ll, form); err != nil {
		return err.Err
	}

	dto := people_dto.EmployeeDTO{
		UserDTO: people_dto.UserDTO{
			Firstname: form.Firstname,
			Surname:   form.Surname,
			Note:      form.Note,
			Status:    form.Status,
			Roles:     form.Roles,
			AvatarID:  form.AvatarID,
			ExpiredAt: form.getExpiredAt(),
			Login:     form.Login,
		},
		PositionID: form.PositionID,
		DeptID:     form.DeptID,
		WorkedAt:   form.getWorkedAt(),
		FiredAt:    form.getFiredAt(),
	}

	err := cnt.peopleAct.UpdateEmployee(ctx, thisUserID, form.UserID, &dto)
	if err != nil {
		ll.Named("peopleAct.UpdateEmployee").Nested(err)
		return cnt.com.Response(c, ll, err)
	}

	cnt.audit.Oper(thisUserID, audit.EntityEmployee, audit.Update, audit.P{
		"form": form,
	})

	ll.Debug("user updated")

	return c.NoContent(http.StatusOK)
}
