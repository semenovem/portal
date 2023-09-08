package people_controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/internal/abc/controller"
	people_dto "github.com/semenovem/portal/internal/abc/people/dto"
	"github.com/semenovem/portal/pkg/throw"
	"net/http"
)

// UserUpdate docs
//
//	@Summary	Обновление данных пользователя
//	@Description `expired` в формате `2001-03-24T00:00:00Z`
//	@Description
//	@Description
//	@Description
//	@Description
//	@Produce	json
//	@Param		login	path		string	true	"проверяемый логин"
//	@Success	200		{object}	freeLoginNameResponse
//	@Failure	400		{object}	fail.Response
//	@Router		/people/employee/:user_id [PATCH]
//	@Tags		people
//	@Security	ApiKeyAuth
func (cnt *Controller) UserUpdate(c echo.Context) error {
	var (
		thisUserID = controller.ExtractThisUserID(c)
		ll         = cnt.logger.Named("UserUpdate").With("thisUserID", thisUserID)
		ctx        = c.Request().Context()
		form       = new(employeeUpdateForm)
	)

	if err := cnt.com.ExtractForm(c, ll, form); err != nil {
		return err.Err
	}

	fmt.Printf(">>>>>>>>>>>> %+v \n", form)

	expiredAt, err := controller.ParseTime(form.ExpiredAt)
	if err != nil {
		err = throw.NewBadRequestTimeFieldErr("expired_at")
		ll.Named("controller.ParseTime").With("form", form).BadRequest(err)
		return cnt.com.Response(c, ll, err)
	}

	workedAt, err := controller.ParseTime(form.WorkedAt)
	if err != nil {
		err = throw.NewBadRequestTimeFieldErr("worked_at")
		ll.Named("controller.ParseTime").With("form", form).BadRequest(err)
		return cnt.com.Response(c, ll, err)
	}

	firedAt, err := controller.ParseTime(form.FiredAt)
	if err != nil {
		err = throw.NewBadRequestTimeFieldErr("fired_at")
		ll.Named("controller.ParseTime").With("form", form).BadRequest(err)
		return cnt.com.Response(c, ll, err)
	}

	dto := people_dto.UserDTO{
		Firstname:  form.Firstname,
		Surname:    form.Surname,
		Note:       form.Note,
		Status:     form.Status,
		Roles:      form.Roles,
		AvatarID:   form.AvatarID,
		Login:      form.Login,
		PositionID: form.PositionID,
		DeptID:     form.DeptID,
		ExpiredAt:  expiredAt,
		WorkedAt:   workedAt,
		FiredAt:    firedAt,
	}

	resp, err := cnt.peopleAct.UpdateEmployee(ctx, thisUserID, form.UserID, &dto)
	if err != nil {
		ll.Named("peopleAct.UpdateEmployee").Nested(err)
		return cnt.com.Response(c, ll, err)
	}

	ll.Debug("user updated")

	return c.JSON(http.StatusOK, newEmployeeUpdateResponse(resp))
}
