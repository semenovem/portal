package people_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/internal/abc/controller"
	"github.com/semenovem/portal/internal/audit"
	_ "github.com/semenovem/portal/pkg/fail"
	"github.com/semenovem/portal/pkg/throw"
	"net/http"
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

	profile, err := cnt.peopleAct.GetUserModel(ctx, thisUserID, thisUserID)
	if err != nil {
		ll = ll.Named("GetUserProfile")
		return cnt.com.Response(c, ll, err)
	}

	return c.JSON(http.StatusOK, newUserProfileView(profile))
}

// CheckLogin docs
//
//	@Summary	Проверяет, свободен ли указанный логин
//	@Description
//	@Produce	json
//	@Param		login	path		string	true	"проверяемый логин"
//	@Success	200		{object}	freeLoginNameResponse
//	@Failure	400		{object}	fail.Response
//	@Router		/people/free-login/:login_name [GET]
//	@Tags		people
//	@Security	ApiKeyAuth
func (cnt *Controller) CheckLogin(c echo.Context) error {
	var (
		thisUserID = controller.ExtractThisUserID(c)
		ll         = cnt.logger.Named("CheckLogin").With("thisUserID", thisUserID)
		ctx        = c.Request().Context()
		form       = new(freeLoginNameForm)
	)

	if err := cnt.com.ExtractForm(c, ll, form); err != nil {
		return err.Err
	}

	exists, err := cnt.peopleAct.CheckLoginName(ctx, thisUserID, form.LoginName)
	if err != nil && !throw.IsBadRequestErr(err) {
		ll = ll.Named("peopleAct.CheckLoginName")
		return cnt.com.Response(c, ll, err)
	}

	var validateErr string
	if err != nil {
		validateErr = err.Error()
	}

	ll.With("free", exists).With("validateErr", validateErr).Debug("checked")

	return c.JSON(http.StatusOK, freeLoginNameResponse{
		Free:        exists,
		ValidateErr: validateErr,
	})
}

// DeleteUser docs
//
//	@Summary	Удаляет пользователя
//	@Description
//	@Produce	json
//	@Param		user_id	path	string	true	"id пользователя"
//	@Success	204		"no-content"
//	@Failure	400		{object}	fail.Response
//	@Router		/people/:user_id [DELETE]
//	@Tags		people
//	@Security	ApiKeyAuth
func (cnt *Controller) DeleteUser(c echo.Context) error {
	var (
		thisUserID = controller.ExtractThisUserID(c)
		ll         = cnt.logger.Named("DeleteUser").With("thisUserID", thisUserID)
		ctx        = c.Request().Context()
		form       = new(userPathForm)
	)

	if err := cnt.com.ExtractForm(c, ll, form); err != nil {
		return err.Err
	}

	ll = ll.With("userID", form.UserID)

	if err := cnt.peopleAct.DeleteUser(ctx, thisUserID, form.UserID); err != nil {
		ll = ll.Named("peopleAct.DeleteUser")
		return cnt.com.Response(c, ll, err)
	}

	cnt.audit.Del(thisUserID, audit.EntityUser, audit.P{
		"userID": form.UserID,
	})

	ll.With("userID", 0).Debug("user is deleted")

	return c.NoContent(http.StatusNoContent)
}

// UndeleteUser docs
//
//	@Summary	Восстанавливает удаленного пользователя
//	@Description
//	@Produce	json
//	@Param		user_id	path	string	true	"id пользователя"
//	@Success	204		"no-content"
//	@Failure	400		{object}	fail.Response
//	@Router		/people/:user_id/undelete [POST]
//	@Tags		people
//	@Security	ApiKeyAuth
func (cnt *Controller) UndeleteUser(c echo.Context) error {
	var (
		thisUserID = controller.ExtractThisUserID(c)
		ll         = cnt.logger.Named("UndeleteUser").With("thisUserID", thisUserID)
		ctx        = c.Request().Context()
		form       = new(userPathForm)
	)

	if err := cnt.com.ExtractForm(c, ll, form); err != nil {
		return err.Err
	}

	ll = ll.With("userID", form.UserID)

	if err := cnt.peopleAct.UndeleteUser(ctx, thisUserID, form.UserID); err != nil {
		ll = ll.Named("peopleAct.UndeleteUser")
		return cnt.com.Response(c, ll, err)
	}

	cnt.audit.Undel(thisUserID, audit.EntityUser, audit.P{
		"userID": form.UserID,
	})

	ll.With("userID", 0).Debug("user is undeleted")

	return c.NoContent(http.StatusNoContent)
}
