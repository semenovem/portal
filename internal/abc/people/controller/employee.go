package people_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/internal/abc/controller"
	people_provider "github.com/semenovem/portal/internal/abc/people/provider"
	"github.com/semenovem/portal/internal/audit"
	"github.com/semenovem/portal/internal/util"
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
		ctx = c.Request().Context()
		ll  = cnt.logger.Func(ctx, "EmployeeHandbook")
	)

	opts := &people_provider.EmployeesSearchOpts{
		Limit:  100,
		Offset: 0,
	}

	result, err := cnt.peopleAct.EmployeeHandbook(ctx, opts)
	if err != nil {
		ll.Named("peopleAct.EmployeeHandbook").Nested(err)
		return cnt.com.Response(c, ll, err)
	}

	// TODO добавить руководителя для УЗ
	employeeViews := newEmployeeProfileViews(result.Employees, result.DeptMap, result.PositionMap)

	response := employeeHandbookResponse{
		Total:     result.Total,
		Employees: employeeViews,
	}

	ll.Debug("received")

	return c.JSON(http.StatusOK, response)
}

// CreateEmployee docs
//
//	@Summary		Создает нового сотрудника
//	@Description	`expired_at, worked_at, fired_at` в формате `2001-03-24T00:00:00Z`
//	@Description	введенный login нужно проверить, что он допустим `/people/free-login/:login_name`
//	@Description
//	@Produce	json
//	@Param		payload	body		employeeCreateForm	true	"данные создаваемого пользователя"
//	@Success	200		{object}	userCreateResponse
//	@Failure	400		{object}	fail.Response
//	@Router		/people/employee [POST]
//	@Tags		people/employee
//	@Security	ApiKeyAuth
func (cnt *Controller) CreateEmployee(c echo.Context) error {
	var (
		ctx        = c.Request().Context()
		thisUserID = controller.ExtractThisUserID(c)
		ll         = cnt.logger.Func(ctx, "CreateEmployee").With("thisUserID", thisUserID)
		form       = new(employeeCreateForm)
	)

	if err := cnt.com.ExtractForm(c, ll, form); err != nil {
		return err.Err
	}

	passwdHash := cnt.userPasswdAuth.Hashing(form.Passwd)

	dto := &people_provider.EmployeeCreateModel{
		UserCreateModel: people_provider.UserCreateModel{
			Firstname:  form.Firstname,
			Surname:    form.Surname,
			Patronymic: form.Patronymic,
			Status:     form.Status,
			Groups:     form.Groups,
			Note:       form.Note,
			AvatarID:   form.AvatarID,
			ExpiredAt:  util.MustParseToTime(form.ExpiredAt),
			Login:      form.Login,
			PasswdHash: &passwdHash,
		},
		PositionID: form.PositionID,
		DeptID:     form.DeptID,
		WorkedAt:   util.MustParseToTime(form.WorkedAt),
		FiredAt:    util.MustParseToTime(form.FiredAt),
	}

	userID, err := cnt.peopleAct.CreateEmployee(ctx, thisUserID, dto)
	if err != nil {
		ll = ll.Named("peopleAct.CreateEmployee")
		return cnt.com.Response(c, ll, err)
	}

	cnt.audit.Oper(thisUserID, audit.EntityEmployee, audit.Create, audit.P{
		"userID": userID,
	})

	ll.With("userID", userID).Info("user created")

	return c.JSON(http.StatusOK, userCreateResponse{UserID: userID})
}

// UpdateEmployee docs
//
//	@Summary		Обновление данных пользователя
//	@Description	`expired_at, worked_at, fired_at` в формате `2001-03-24T00:00:00Z`
//	@Description
//	@Description	json объект должен содержать только те поля, которые отправляются на редактирование
//	@Description	все поля опциональны
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
		ctx        = c.Request().Context()
		thisUserID = controller.ExtractThisUserID(c)
		ll         = cnt.logger.Func(ctx, "UpdateEmployee").With("thisUserID", thisUserID)
		form       = new(employeeUpdateForm)
	)

	if err := cnt.com.ExtractForm(c, ll, form); err != nil {
		return err.Err
	}

	ll.With("userID", form.UserID)

	dto := people_provider.EmployeeCreateModel{
		UserCreateModel: people_provider.UserCreateModel{
			Firstname:  form.Firstname,
			Surname:    form.Surname,
			Patronymic: form.Patronymic,
			Status:     form.Status,
			Groups:     form.Groups,
			Note:       form.Note,
			AvatarID:   form.AvatarID,
			ExpiredAt:  util.MustParseToTime(form.ExpiredAt),
			Login:      form.Login,
		},
		PositionID: form.PositionID,
		DeptID:     form.DeptID,
		WorkedAt:   util.MustParseToTime(form.WorkedAt),
		FiredAt:    util.MustParseToTime(form.FiredAt),
	}

	err := cnt.peopleAct.UpdateEmployee(ctx, thisUserID, form.UserID, &dto)
	if err != nil {
		ll.Named("peopleAct.UpdateEmployee").Nested(err)
		return cnt.com.Response(c, ll, err)
	}

	cnt.audit.Oper(thisUserID, audit.EntityEmployee, audit.Update, audit.P{
		"form": form,
	})

	ll.Info("user updated")

	return c.NoContent(http.StatusOK)
}
