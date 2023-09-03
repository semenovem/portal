package people_controller

import (
	"github.com/labstack/echo/v4"
	people_action "github.com/semenovem/portal/internal/abc/people/action"
	"github.com/semenovem/portal/internal/audit"
	"net/http"
	"time"

	_ "github.com/semenovem/portal/pkg/fail"
)

// CreateUser docs
//
//	@Summary	Создает пользователя
//	@Description
//	@Produce	json
//	@Param		payload	body		createUserForm	true	"данные создаваемого пользователя"
//	@Success	200		{object}	userPublicProfileView
//	@Failure	400		{object}	fail.Response
//	@Router		/people [POST]
//	@Tags		people
//	@Security	ApiKeyAuth
func (cnt *Controller) CreateUser(c echo.Context) error {
	var (
		ll   = cnt.logger.Named("CreateUser")
		ctx  = c.Request().Context()
		form = new(createUserForm)
	)

	thisUserID, nested := cnt.com.ExtractUserAndForm(c, form)
	if nested != nil {
		ll.Named("ExtractUserAndForm").Nestedf(nested.Message())
		return cnt.fail.SendNested(c, "", nested)
	}

	model := &people_action.CreateUserDTO{
		FirstName:  form.FirstName,
		Surname:    form.Surname,
		Note:       form.Note,
		Status:     form.getStatus(),
		Roles:      form.getRoles(),
		ExpiredAt:  time.Time{},
		Login:      form.Login,
		PasswdHash: form.Passwd,
	}

	profile, err := cnt.peopleAct.CreateUser(ctx, thisUserID, model)
	if err != nil {
		ll.Named("CreateUser").Nested(err)
		return cnt.com.Response(c, err)
	}

	cnt.audit.Oper(thisUserID, audit.EntityUser, audit.Create, audit.P{
		"userID": profile.ID,
	})

	ll.With("userID", 0).Debug("user created")

	return c.JSON(http.StatusOK, newUserProfileView(profile))
}

// DeleteUser docs
//
//	@Summary	Удаляет пользователя
//	@Description
//	@Produce	json
//	@Param		user_id	path		string	true	"id пользователя"
//	@Success	200		{object}	userPublicProfileView
//	@Failure	400		{object}	fail.Response
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
		ll.Named("ExtractUserAndForm").Nestedf(nested.Message())
		return cnt.fail.SendNested(c, "", nested)
	}

	//profile, err := cnt.peopleAct.GetUserProfile(ctx, thisUserID, form.UserID)
	//if err != nil {
	//	ll = ll.Named("GetUserProfile").With("thisUserID", thisUserID)
	//
	//	switch err.(type) {
	//	case *action.NotFoundErr:
	//		ll.NotFoundTag().Info(err.Error())
	//		return cnt.fail.Send(c, "", http.StatusNotFound, err)
	//	case *action.ForbiddenErr:
	//		ll.DenyTag().Info(err.Error())
	//		return cnt.fail.Send(c, "", http.StatusForbidden, err)
	//	default:
	//		ll.Nested(err)
	//		return cnt.fail.SendInternalServerErr(c, "", err)
	//	}
	//}

	cnt.audit.Oper(thisUserID, audit.EntityUser, audit.Delete, audit.P{
		"userID": form.UserID,
	})

	ll.With("userID", 0).Debug("user deleted")

	return c.NoContent(http.StatusNoContent)
}
