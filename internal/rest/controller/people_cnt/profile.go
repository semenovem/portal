package people_cnt

import (
	"fmt"
	"github.com/labstack/echo/v4"

	_ "github.com/semenovem/portal/internal/rest/view"
	_ "github.com/semenovem/portal/pkg/failing"
)

// SelfProfile docs
//
//	@Summary	Получить свой профиль
//	@Description
//	@Produce	json
//	@Success	200		{object}	ProfileFull
//	@Failure	400		{object}	failing.Response
//	@Router		/people/self/profile [GET]
//	@Tags		people
//	@Security	ApiKeyAuth
func (ct *Controller) SelfProfile(c echo.Context) error {
	fmt.Println("!!!!!!!!!!!!!")

	return nil
}

// Profile docs
//
//	@Summary	Получить профиль пользователя по его ID
//	@Description Проверяет действующие права на просмотр расширенных данных пользователя
//	@Description
//	@Produce	json
//	@Param		user_id	path		string	true	"id пользователя"
//	@Success	200		{object}	ProfileFull
//	@Failure	400		{object}	failing.Response
//	@Router		/people/:user_id/profile [GET]
//	@Tags		people
//	@Security	ApiKeyAuth
func (ct *Controller) Profile(c echo.Context) error {
	fmt.Println("!!!!!!!!!!!!!")

	return nil
}
