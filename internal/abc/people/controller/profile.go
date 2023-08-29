package people_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/internal/abc/action"
	"net/http"

	_ "github.com/semenovem/portal/internal/rest/view"
	_ "github.com/semenovem/portal/pkg/failing"
)

// SelfProfile docs
//
//	@Summary	Получить свой профиль
//	@Description
//	@Produce	json
//	@Success	200	{object}	userProfileView
//	@Failure	400	{object}	failing.Response
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
		return cnt.failing.SendNested(c, "", nested)
	}

	profile, err := cnt.peopleAct.GetUserProfile(ctx, thisUserID, thisUserID)
	if err != nil {
		ll = ll.Named("GetUserProfile").With("thisUserID", thisUserID)

		switch err.(type) {
		case action.NotFoundErr:
			ll.NotFound(err)
			return cnt.failing.Send(c, "", http.StatusNotFound, err)
		case action.ForbiddenErr:
			ll.Deny(err)
			return cnt.failing.Send(c, "", http.StatusForbidden, err)
		default:
			ll.Nested(err)
			return cnt.failing.SendInternalServerErr(c, "", err)
		}
	}

	return c.JSON(http.StatusOK, newUserProfileView(profile))
}

// Profile docs
//
//	@Summary		Получить профиль пользователя по его ID
//	@Description	Проверяет действующие права на просмотр расширенных данных пользователя
//	@Description
//	@Produce	json
//	@Param		user_id	path		string	true	"id пользователя"
//	@Success	200		{object}	userProfileView
//	@Failure	400		{object}	failing.Response
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
		return cnt.failing.Send(c, "", http.StatusBadRequest)
	}

	profile, err := cnt.peopleAct.GetUserProfile(ctx, thisUserID, form.UserID)
	if err != nil {
		ll = ll.Named("GetUserProfile").With("thisUserID", thisUserID)

		switch err.(type) {
		case *action.NotFoundErr:
			ll.NotFound(err)
			return cnt.failing.Send(c, "", http.StatusNotFound, err)
		case *action.ForbiddenErr:
			ll.Deny(err)
			return cnt.failing.Send(c, "", http.StatusForbidden, err)
		default:
			ll.Nested(err)
			return cnt.failing.SendInternalServerErr(c, "", err)
		}
	}

	return c.JSON(http.StatusOK, newUserProfileView(profile))
}
