package people_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/pkg/throw"
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
		ll  = cnt.logger.Named("SelfProfile")
		ctx = c.Request().Context()
	)

	thisUserID, nested := cnt.com.ExtractThisUser(c)
	if nested != nil {
		ll.Named("ExtractThisUser").Nestedf(nested.Message())
		return cnt.fail.SendNested(c, "", nested)
	}

	profile, err := cnt.peopleAct.GetUserProfile(ctx, thisUserID, thisUserID)
	if err != nil {
		ll = ll.Named("GetUserProfile").With("thisUserID", thisUserID)

		switch err.(type) {
		case throw.NotFoundErr:
			ll.NotFound(err)
			return cnt.fail.Send(c, "", http.StatusNotFound, err)
		case *throw.AccessErr:
			ll.Deny(err)
			return cnt.fail.Send(c, "", http.StatusForbidden, err)
		default:
			ll.Nested(err)
			return cnt.fail.SendInternalServerErr(c, "", err)
		}
	}

	return c.JSON(http.StatusOK, newUserPublicProfileView(profile))
}

// Profile docs
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
func (cnt *Controller) Profile(c echo.Context) error {
	var (
		ll   = cnt.logger.Named("Profile")
		ctx  = c.Request().Context()
		form = new(UserForm)
	)

	thisUserID, nested := cnt.com.ExtractUserAndForm(c, form)
	if nested != nil {
		ll.Named("ExtractUserAndForm").Nestedf(nested.Message())
		return cnt.fail.SendNested(c, "", nested)
	}

	profile, err := cnt.peopleAct.GetUserProfile(ctx, thisUserID, form.UserID)
	if err != nil {
		ll = ll.Named("GetUserProfile").With("thisUserID", thisUserID)

		switch err.(type) {
		case throw.NotFoundErr:
			ll.NotFound(err)
			return cnt.fail.Send(c, "", http.StatusNotFound, err)
		case *throw.AccessErr:
			ll.Deny(err)
			return cnt.fail.Send(c, "", http.StatusForbidden, err)
		default:
			ll.Nested(err)
			return cnt.fail.SendInternalServerErr(c, "", err)
		}
	}

	return c.JSON(http.StatusOK, newUserPublicProfileView(profile))
}
