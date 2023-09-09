package people_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/internal/abc/controller"
	people_dto "github.com/semenovem/portal/internal/abc/people/dto"
	_ "github.com/semenovem/portal/pkg/fail"
	"net/http"
)

// EmployeeUpdate docs
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
//	@Tags		people
//	@Security	ApiKeyAuth
func (cnt *Controller) EmployeeUpdate(c echo.Context) error {
	var (
		thisUserID = controller.ExtractThisUserID(c)
		ll         = cnt.logger.Named("EmployeeUpdate").With("thisUserID", thisUserID)
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

	ll.Debug("user updated")

	return c.NoContent(http.StatusOK)
}
