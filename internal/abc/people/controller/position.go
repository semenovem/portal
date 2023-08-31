package people_controller

import (
	"fmt"
	"github.com/labstack/echo/v4"

	_ "github.com/semenovem/portal/pkg/fail"
)

// Position docs
//
//	@Summary		Получить профиль пользователя по его ID
//	@Description	Проверяет действующие права на просмотр расширенных данных пользователя
//	@Description
//	@Produce	json
//	@Param		user_id	path		string	true	"id пользователя"
//	@Success	200		{object}	string
//	@Failure	400		{object}	fail.Response
//	@Router		/people/positions [GET]
//	@Tags		people/position
//	@Security	ApiKeyAuth
func (cnt *Controller) Position(c echo.Context) error {
	fmt.Println("!!!!!!!!!!!!!")

	return nil
}
