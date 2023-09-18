package people_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/internal/abc/controller"
	"net/http"

	_ "github.com/semenovem/portal/pkg/fail"
)

// UserProfile docs
//
//	@Summary		Получить профиль пользователя по его ID
//	@Description	Проверяет действующие права на просмотр расширенных данных пользователя
//	@Description
//	@Produce	json
//	@Param		user_id	path		string	true	"id пользователя"
//	@Success	200		{object}	userPublicProfileView
//	@Failure	400		{object}	fail.Response
//	@Router		/people/:user_id/profile [GET]
//	@Tags		people
//	@Security	ApiKeyAuth
func (cnt *Controller) UserProfile(c echo.Context) error {
	var (
		ctx        = c.Request().Context()
		thisUserID = controller.ExtractThisUserID(c)
		ll         = cnt.logger.Func(ctx, "UserPublicProfile").With("thisUserID", thisUserID)
		form       = new(userPathForm)
	)

	if err := cnt.com.ExtractForm(c, ll, form); err != nil {
		return err.Err
	}

	profile, err := cnt.peopleAct.GetUserModel(ctx, thisUserID, form.UserID)
	if err != nil {
		ll = ll.Named("GetUserProfile").With("user", form.UserID)
		return cnt.com.Response(c, ll, err)
	}

	ll.Debug("received")

	return c.JSON(http.StatusOK, newUserProfileView(profile))
}

// UserPublicProfile docs
//
//	@Summary	Получить публичный профиль пользователя по его ID
//	@Description
//	@Produce	json
//	@Param		user_id	path		string	true	"id пользователя"
//	@Success	200		{object}	userPublicProfileView
//	@Failure	400		{object}	fail.Response
//	@Router		/people/:user_id/profile/public [GET]
//	@Tags		people
//	@Security	ApiKeyAuth
func (cnt *Controller) UserPublicProfile(c echo.Context) error {
	var (
		ctx        = c.Request().Context()
		thisUserID = controller.ExtractThisUserID(c)
		ll         = cnt.logger.Func(ctx, "UserPublicProfile").With("thisUserID", thisUserID)
		form       = new(userPathForm)
	)

	if err := cnt.com.ExtractForm(c, ll, form); err != nil {
		return err.Err
	}

	profile, err := cnt.peopleAct.GetUserModel(ctx, thisUserID, form.UserID)
	if err != nil {
		ll = ll.Named("GetUserProfile")
		return cnt.com.Response(c, ll, err)
	}

	ll.Debug("received")

	return c.JSON(http.StatusOK, newUserPublicProfileView(profile))
}
