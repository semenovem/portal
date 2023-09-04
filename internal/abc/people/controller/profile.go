package people_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/internal/abc/controller"
	"net/http"

	_ "github.com/semenovem/portal/internal/rest/view"
	_ "github.com/semenovem/portal/pkg/fail"
)

// SelfProfile docs
//
//	@Summary	Получить свой профиль
//	@Description
//	@Produce	json
//	@Success	200	{object}	userPublicProfileView
//	@Failure	400	{object}	fail.Response
//	@Router		/people/self/profile [GET]
//	@Tags		people
//	@Security	ApiKeyAuth
func (cnt *Controller) SelfProfile(c echo.Context) error {
	var (
		thisUserID = controller.ExtractThisUserID(c)
		ll         = cnt.logger.Named("SelfProfile").With("thisUserID", thisUserID)
		ctx        = c.Request().Context()
	)

	profile, err := cnt.peopleAct.GetUserProfile(ctx, thisUserID, thisUserID)
	if err != nil {
		ll = ll.Named("GetUserProfile")
		return cnt.com.Response(c, err, ll)
	}

	return c.JSON(http.StatusOK, newUserProfileView(profile))
}

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
		thisUserID = controller.ExtractThisUserID(c)
		ll         = cnt.logger.Named("UserPublicProfile").With("thisUserID", thisUserID)
		ctx        = c.Request().Context()
		form       = new(userForm)
	)

	if err := cnt.com.ExtractForm(c, ll, form); err != nil {
		return err
	}

	profile, err := cnt.peopleAct.GetUserProfile(ctx, thisUserID, form.UserID)
	if err != nil {
		ll = ll.Named("GetUserProfile").With("user", form.UserID)
		return cnt.com.Response(c, err, ll)
	}

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
		thisUserID = controller.ExtractThisUserID(c)
		ll         = cnt.logger.Named("UserPublicProfile").With("thisUserID", thisUserID)
		ctx        = c.Request().Context()
		form       = new(userForm)
	)

	if err := cnt.com.ExtractForm(c, ll, form); err != nil {
		return err
	}

	profile, err := cnt.peopleAct.GetUserProfile(ctx, thisUserID, form.UserID)
	if err != nil {
		ll = ll.Named("GetUserProfile")
		return cnt.com.Response(c, err, ll)
	}

	return c.JSON(http.StatusOK, newUserPublicProfileView(profile))
}
