package people_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Handbook docs
//
//	@Summary	Справочник сотрудников
//	@Description Доступен в локальной сети без авторизации
//	@Description
//	@Produce	json
//	@Success	200	{object}	publicHandbookResponse
//	@Failure	400	{object}	fail.Response
//	@Router		/people/handbook [GET]
//	@Tags		people
func (cnt *Controller) Handbook(c echo.Context) error {
	var (
		ll  = cnt.logger.Named("Handbook")
		ctx = c.Request().Context()
	)

	result, err := cnt.peopleAct.PublicHandbook(ctx)
	if err != nil {
		ll.Named("peopleAct.PublicHandbook").Nested(err)
		return cnt.com.Response(c, ll, err)
	}

	ll.Debug("success")

	response := publicHandbookResponse{
		Total:     result.Total,
		Employees: newEmployeePublicProfileViews(result.Employees),
	}

	return c.JSON(http.StatusOK, response)
}
