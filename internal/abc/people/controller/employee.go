package people_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// PublicHandbook docs
//
//	@Summary	Справочник сотрудников
//	@Description
//	@Produce	json
//	@Success	200	{object}	publicHandbookResponse
//	@Failure	400	{object}	fail.Response
//	@Router		/people/handbook [GET]
//	@Tags		people
//	@Security	ApiKeyAuth
func (cnt *Controller) PublicHandbook(c echo.Context) error {
	var (
		ll = cnt.logger.Named("PublicHandbook")
		//ctx        = c.Request().Context()
	)

	ll.Debug("success")

	return c.JSON(http.StatusOK, publicHandbookResponse{})
}
