package auth_controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/internal/abc/auth/action"
	"github.com/semenovem/portal/internal/audit"
	"net/http"

	_ "github.com/semenovem/portal/pkg/failing"
)

// CreateOnetimeLink docs
//
//	@Summary	Создание ссылки для одноразовой авторизации
//	@Description
//	@Produce	json
//	@Param		payload	body		onetimeAuthForm	true	"данные для создания сессии"
//	@Success	200		{object}	onetimeAuthResponse
//	@Failure	400		{object}	failing.Response
//	@Router		/auth/onetime [POST]
//	@Tags		auth
//	@Security	ApiKeyAuth
func (cnt *Controller) CreateOnetimeLink(c echo.Context) error {
	var (
		ll   = cnt.logger.Named("LoginOnetimeLink")
		form = new(onetimeAuthForm)
		ctx  = c.Request().Context()
	)

	thisUserID, nested := cnt.com.ExtractUserAndForm(c, form)
	if nested != nil {
		ll.Named("ExtractUserAndForm").Nested(nested.Message())
		return cnt.failing.SendNested(c, "", nested)
	}

	ll = ll.With("userID", form.UserID)

	entryID, err := cnt.authAct.CreateOnetimeEntry(ctx, form.UserID)
	if err != nil {
		ll.Named("CreateOnetimeEntry").Nested(err.Error())

		if auth_action.IsAuthErr(err) {
			return cnt.failing.Send(c, "", http.StatusBadRequest, err)
		}

		return cnt.failing.SendInternalServerErr(c, "", err)
	}

	cnt.audit.Auth(thisUserID, audit.CreateOnetimeEntry, audit.P{
		"user_id":    form.UserID,
		"session_id": entryID,
	})

	return c.JSON(http.StatusOK, onetimeAuthResponse{
		URI:     fmt.Sprintf("https://portal.glazkoff.ru/auth/onetime/%s", entryID.String()),
		EntryID: entryID.String(),
	})
}

// LoginOnetimeLink docs
//
//	@Summary	Логин по одноразовой ссылке
//	@Description
//	@Produce	json
//	@Param		session_id	path		string	true	"id сессии авторизации"
//	@Success	200			{object}	loginResponse
//	@Failure	400			{object}	failing.Response
//	@Router		/auth/onetime/:entry_id [POST]
//	@Tags		auth
//	@Security	ApiKeyAuth
func (cnt *Controller) LoginOnetimeLink(c echo.Context) error {
	var (
		ll   = cnt.logger.Named("LoginOnetimeLink")
		form = new(entryPointForm)
		ctx  = c.Request().Context()
	)

	if nested := cnt.com.ExtractForm(c, form); nested != nil {
		ll.Named("ExtractForm").Nested(nested.Message())
		return cnt.failing.SendNested(c, "", nested)
	}

	ll = ll.With("sessionID", form.EntryID)

	session, err := cnt.authAct.LoginByOnetimeEntryID(ctx, form.EntryID)
	if err != nil {
		ll.Named("LoginByOnetimeEntryID").Nested(err.Error())
		if auth_action.IsAuthErr(err) {
			return cnt.failing.Send(c, "", http.StatusNotFound, err)
		}

		return cnt.failing.SendInternalServerErr(c, "", err)
	}

	cnt.audit.Auth(session.UserID, audit.UserLogin, audit.P{
		"sessionID":        session.ID,
		"byOnetimeEntryID": form.EntryID,
	})

	pair, nested := cnt.pairToken(session)
	if nested != nil {
		ll.Named("pairToken").Nested(nested.Message())
		return cnt.failing.SendNested(c, "", nested)
	}

	for _, cookie := range cnt.refreshTokenCookies(pair.Refresh) {
		c.SetCookie(cookie)
	}

	return c.JSON(http.StatusOK, loginResponse{
		refreshResponse: refreshResponse{
			AccessToken:  pair.Access,
			RefreshToken: pair.Refresh,
		},
		UserID: session.UserID,
	})
}