package people_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/internal/audit"
	"github.com/semenovem/portal/pkg/it"
	"net/http"

	_ "github.com/semenovem/portal/pkg/failing"
)

// CreateUser docs
//
//	@Summary	Создает пользователя
//	@Description
//	@Produce	json
//	@Param		user_id	path		string	true	"id пользователя"
//	@Success	200		{object}	userProfileView
//	@Failure	400		{object}	failing.Response
//	@Router		/people [POST]
//	@Tags		people
//	@Security	ApiKeyAuth
func (cnt *Controller) CreateUser(c echo.Context) error {
	var (
		ll = cnt.logger.Named("CreateUser")
		//ctx  = c.Request().Context()
		form = new(CreateUserForm)
	)

	thisUserID, nested := cnt.com.ExtractUserAndForm(c, form)
	if nested != nil {
		ll.Named("ExtractUserAndForm").Nested(nested.Message())
		return cnt.failing.Send(c, "", http.StatusBadRequest)
	}

	//profile, err := cnt.peopleAct.GetUserProfile(ctx, thisUserID, form.UserID)
	//if err != nil {
	//	ll = ll.Named("GetUserProfile").With("thisUserID", thisUserID)
	//
	//	switch err.(type) {
	//	case *action.NotFoundErr:
	//		ll.NotFoundTag().Info(err.Error())
	//		return cnt.failing.Send(c, "", http.StatusNotFound, err)
	//	case *action.ForbiddenErr:
	//		ll.DenyTag().Info(err.Error())
	//		return cnt.failing.Send(c, "", http.StatusForbidden, err)
	//	default:
	//		ll.Nested(err.Error())
	//		return cnt.failing.SendInternalServerErr(c, "", err)
	//	}
	//}

	userProfile := it.UserProfile{}

	cnt.audit.Oper(thisUserID, audit.EntityUser, audit.Create, audit.P{
		"userID": userProfile.ID,
	})

	ll.With("userID", 0).Debug("user created")

	return c.JSON(http.StatusOK, newUserProfileView(&userProfile))
}

// DeleteUser docs
//
//	@Summary	Удаляет пользователя
//	@Description
//	@Produce	json
//	@Param		user_id	path		string	true	"id пользователя"
//	@Success	200		{object}	userProfileView
//	@Failure	400		{object}	failing.Response
//	@Router		/people/:user_id [DELETE]
//	@Tags		people
//	@Security	ApiKeyAuth
func (cnt *Controller) DeleteUser(c echo.Context) error {
	var (
		ll = cnt.logger.Named("DeleteUser")
		//ctx  = c.Request().Context()
		form = new(UserForm)
	)

	thisUserID, nested := cnt.com.ExtractUserAndForm(c, form)
	if nested != nil {
		ll.Named("ExtractUserAndForm").Nested(nested.Message())
		return cnt.failing.Send(c, "", http.StatusBadRequest)
	}

	//profile, err := cnt.peopleAct.GetUserProfile(ctx, thisUserID, form.UserID)
	//if err != nil {
	//	ll = ll.Named("GetUserProfile").With("thisUserID", thisUserID)
	//
	//	switch err.(type) {
	//	case *action.NotFoundErr:
	//		ll.NotFoundTag().Info(err.Error())
	//		return cnt.failing.Send(c, "", http.StatusNotFound, err)
	//	case *action.ForbiddenErr:
	//		ll.DenyTag().Info(err.Error())
	//		return cnt.failing.Send(c, "", http.StatusForbidden, err)
	//	default:
	//		ll.Nested(err.Error())
	//		return cnt.failing.SendInternalServerErr(c, "", err)
	//	}
	//}

	cnt.audit.Oper(thisUserID, audit.EntityUser, audit.Delete, audit.P{
		"userID": form.UserID,
	})

	ll.With("userID", 0).Debug("user deleted")

	return c.NoContent(http.StatusNoContent)
}
