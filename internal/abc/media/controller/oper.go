package media_controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"

	_ "github.com/semenovem/portal/pkg/failing"
)

// Upload docs
//
//	@Summary		Сохранение произвольных клиентских данных
//	@Description
//	@Produce	json
//	@Param		store_code	path	string		true	"code store"
//	@Param		payload		body	storeForm	true	"Данные для сохранения"
//	@Accept		multipart/form-data
//	@Success	201			"no content"
//	@Failure	400			{object}	failing.Response
//	@Router		/media/upload [POST]
//	@Tags		store
//	@Security	ApiKeyAuth
func (cnt *Controller) Upload(c echo.Context) error {
	var (
		ll  = cnt.logger.Named("Upload")
		ctx = c.Request().Context()
	)

	thisUser, nested := cnt.com.ExtractThisUser(c)
	if nested != nil {
		ll.Named("ExtractThisUser").Nested(nested.Message())
		return cnt.failing.SendNested(c, "", nested)
	}

	// TODO Проверить, может ли пользователь загружать файлы

	form, err := c.MultipartForm()
	if err != nil {
		ll.Named("MultipartForm").Debug(err.Error())
		return cnt.failing.Send(c, "", http.StatusBadRequest, err)
	}

	fmt.Println()

	for k, v := range form.Value {
		fmt.Printf(">>>>>> k = %s  v = %+v\n", k, v)
	}

	for name, files := range form.File {
		fmt.Printf("> files >>>>> k = %s  v = %+v\n", name, files)

		for _, fdata := range files {
			// 1. проверить размер
			// 2. тип

			file, err := fdata.Open()
			if err != nil {
				//ll.Named()
			}

			cnt.mediaAct.Upload(ctx, thisUser, file)
		}
	}

	//thisUserID, nested := cnt.com.ExtractUserAndForm(c, form)
	//if nested != nil {
	//	ll.Named("ExtractForm").Nested(nested.Message())
	//	return cnt.failing.SendNested(c, "", nested)
	//}
	//
	//ll = ll.With("store_path", form.StorePath).With("thisUserID", thisUserID)
	//
	//if err := cnt.storeAct.Store(ctx, thisUserID, form.StorePath, form.Payload); err != nil {
	//	ll.Named("Store").Nested(err.Error())
	//	return cnt.failing.SendInternalServerErr(c, "", err)
	//}
	//
	//ll.Debug("stored")

	return c.NoContent(http.StatusOK)
}

//func (cnt *Controller)

//// Load docs
////
////	@Summary	Чтение произвольных клиентских данных
////	@Description
////	@Produce	json
////	@Param		store_code	path		string	true	"code store"
////	@Success	200			{object}	loadView
////	@Failure	400			{object}	failing.Response
////	@Router		/store/:store_path [GET]
////	@Tags		store
////	@Security	ApiKeyAuth
//func (cnt *Controller) Load(c echo.Context) error {
//	var (
//		ll   = cnt.logger.Named("Store")
//		form = new(storePathForm)
//		ctx  = c.Request().Context()
//	)
//
//	thisUserID, nested := cnt.com.ExtractUserAndForm(c, form)
//	if nested != nil {
//		ll.Named("ExtractForm").Nested(nested.Message())
//		return cnt.failing.SendNested(c, "", nested)
//	}
//
//	ll = ll.With("store_path", form.StorePath).With("thisUserID", thisUserID)
//
//	payload, err := cnt.storeAct.Load(ctx, thisUserID, form.StorePath)
//	if err != nil {
//		ll.Named("Load").Nested(err.Error())
//
//		if errors.Is(err, action.ErrNotFound) {
//			return cnt.failing.Send(c, "", http.StatusNotFound, err)
//		}
//
//		return cnt.failing.SendInternalServerErr(c, "", err)
//	}
//
//	ll.Debug("loaded")
//
//	return c.JSON(http.StatusOK, loadView{payload})
//}

// Delete docs
//
//	@Summary	Удаление
//	@Description
//	@Produce	json
//	@Param		store_code	path	string	true	"code store"
//	@Success	204			"no content"
//	@Failure	400			{object}	failing.Response
//	@Router		/store/:store_path [DELETE]
//	@Tags		store
//	@Security	ApiKeyAuth
//func (cnt *Controller) Delete(c echo.Context) error {
//	var (
//		ll   = cnt.logger.Named("Store")
//		form = new(storePathForm)
//		ctx  = c.Request().Context()
//	)
//
//	thisUserID, nested := cnt.com.ExtractUserAndForm(c, form)
//	if nested != nil {
//		ll.Named("ExtractForm").Nested(nested.Message())
//		return cnt.failing.SendNested(c, "", nested)
//	}
//
//	ll = ll.With("store_path", form.StorePath).With("thisUserID", thisUserID)
//
//	if err := cnt.storeAct.Delete(ctx, thisUserID, form.StorePath); err != nil {
//		ll.Named("Load").Nested(err.Error())
//
//		if errors.Is(err, action.ErrNotFound) {
//			return cnt.failing.Send(c, "", http.StatusNotFound, err)
//		}
//
//		return cnt.failing.SendInternalServerErr(c, "", err)
//	}
//
//	ll.Debug("deleted")
//
//	return c.NoContent(http.StatusOK)
//}
