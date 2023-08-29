package media_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/pkg/it"
	"mime/multipart"
	"net/http"

	_ "github.com/semenovem/portal/pkg/failing"
)

// FileUpload docs
//
//	@Summary		Сохранение файлов
//	@Description	note - подпись к файлу (опционально)
//	@Description	file - файл
//	@Description
//	@Produce	json
//	@Accept		multipart/form-data
//	@Success	201			{object}	fileUploadResponse
//	@Failure	400			{object}	failing.Response
//	@Router		/media/file [POST]
//	@Tags		media
//	@Security	ApiKeyAuth
func (cnt *Controller) FileUpload(c echo.Context) error {
	var (
		ll         = cnt.logger.Named("FileUpload")
		ctx        = c.Request().Context()
		note       string
		fileHeader *multipart.FileHeader
	)

	thisUserID, nested := cnt.com.ExtractThisUser(c)
	if nested != nil {
		ll.Named("ExtractThisUser").Nestedf(nested.Message())
		return cnt.failing.SendNested(c, "", nested)
	}

	form, err := c.MultipartForm()
	if err != nil {
		ll.Named("MultipartForm").Errorf(err.Error())
		return cnt.failing.Send(c, "", http.StatusBadRequest, err)
	}

	if notes := form.Value[uploadNoteKey]; len(notes) != 0 {
		switch len(notes) {
		case 1:
			note = notes[0]
		default:
			ll.With("notes", notes).BadRequest(it.ErrOverNote)
			return cnt.failing.Send(c, "", http.StatusBadRequest, it.ErrOverNote)
		}
	}

	if files := form.File[fileUploadKey]; len(files) == 0 {
		ll.BadRequest(it.ErrNoFile)
		return cnt.failing.Send(c, "", http.StatusBadRequest, it.ErrNoFile)
	} else if len(files) > 1 {
		ll.BadRequest(it.ErrOverFile)
		return cnt.failing.Send(c, "", http.StatusBadRequest, it.ErrOverFile)
	} else {
		fileHeader = files[0]
	}

	mediaFile, nested := cnt.processUploadingFile(ctx, thisUserID, fileHeader, note)
	if nested != nil {
		ll.Named("processUploadingFile").Nested(err)
		return cnt.failing.SendNested(c, "", nested)
	}

	return c.JSON(http.StatusOK, newFileUploadResponse(mediaFile))
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
//		ll.Named("ExtractForm").Nestedf(nested.Message())
//		return cnt.failing.SendNested(c, "", nested)
//	}
//
//	ll = ll.With("store_path", form.StorePath).With("thisUserID", thisUserID)
//
//	payload, err := cnt.storeAct.Load(ctx, thisUserID, form.StorePath)
//	if err != nil {
//		ll.Named("Load").Nested(err)
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
//		ll.Named("ExtractForm").Nestedf(nested.Message())
//		return cnt.failing.SendNested(c, "", nested)
//	}
//
//	ll = ll.With("store_path", form.StorePath).With("thisUserID", thisUserID)
//
//	if err := cnt.storeAct.Delete(ctx, thisUserID, form.StorePath); err != nil {
//		ll.Named("Load").Nested(err)
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
