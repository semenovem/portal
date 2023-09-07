package people_controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/internal/abc/controller"
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
		//ctx        = c.Request().Context()
		form = new(userUpdateForm)
	)

	if err := cnt.com.ExtractForm(c, ll, form); err != nil {
		return err
	}

	fmt.Printf(">>>>>>>>>>>> %+v \n", form.Fields)

	ll.Debug("user updated")

	return c.JSON(http.StatusOK, "")
}
