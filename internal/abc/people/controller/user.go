package people_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/internal/abc/controller"
	"github.com/semenovem/portal/internal/abc/people/action"
	"github.com/semenovem/portal/internal/audit"
	_ "github.com/semenovem/portal/pkg/fail"
	"github.com/semenovem/portal/pkg/it"
	"net/http"
)

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
		return err
	}

	exists, err := cnt.peopleAct.CheckLoginName(ctx, thisUserID, form.LoginName)
	if err != nil && !it.IsValidateErr(err) {
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

// CreateUser docs
//
//	@Summary		Создает пользователя
//	@Description	`expired_at` в формате `2001-03-24T00:00:00Z`
//	@Description	введенный login нужно проверить, что он допустим `/people/free-login/:login_name`
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
		thisUserID = controller.ExtractThisUserID(c)
		ll         = cnt.logger.Named("CreateUser").With("thisUserID", thisUserID)
		ctx        = c.Request().Context()
		form       = new(createUserForm)
	)

	if err := cnt.com.ExtractForm(c, ll, form); err != nil {
		return err
	}

	model := &people_action.CreateUserDTO{
		FirstName: form.FirstName,
		Surname:   form.Surname,
		Note:      form.getNote(),
		Status:    form.getStatus(),
		Roles:     form.getRoles(),
		ExpiredAt: form.ExpiredAt,
		Login:     form.getLogin(),
		Passwd:    form.Passwd,
		AvatarID:  form.AvatarID,
	}

	userID, err := cnt.peopleAct.CreateUser(ctx, thisUserID, model)
	if err != nil {
		ll = ll.Named("peopleAct.CreateUser")
		return cnt.com.Response(c, ll, err)
	}

	cnt.audit.Oper(thisUserID, audit.EntityUser, audit.Create, audit.P{
		"userID": userID,
	})

	ll.With("userID", userID).Debug("user created")

	return c.JSON(http.StatusOK, userCreateResponse{UserID: userID})
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
		form       = new(userForm)
	)

	if err := cnt.com.ExtractForm(c, ll, form); err != nil {
		return err
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
