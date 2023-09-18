package media_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/internal/abc/controller"
	"github.com/semenovem/portal/pkg/it"
	"github.com/semenovem/portal/pkg/throw"
	"mime/multipart"
	"net/http"

	_ "github.com/semenovem/portal/pkg/fail"
)

// UploadAvatar docs
//
//	@Summary		Загрузка аватара пользователя
//	@Description	file - файл
//	@Description
//	@Produce	json
//	@Accept		multipart/form-data
//	@Success	201	{object}	avatarUpload
//	@Failure	400	{object}	fail.Response
//	@Router		/media/avatar [POST]
//	@Tags		media
//	@Security	ApiKeyAuth
func (cnt *Controller) UploadAvatar(c echo.Context) error {
	var (
		ctx        = c.Request().Context()
		thisUserID = controller.ExtractThisUserID(c)
		ll         = cnt.logger.Func(ctx, "UploadAvatar").With("thisUserID", thisUserID)
		note       string
		fileHeader *multipart.FileHeader
	)

	form, err := c.MultipartForm()
	if err != nil {
		ll.Named("MultipartForm").Errorf(err.Error())
		return cnt.fail.Send(c, "", http.StatusBadRequest, err)
	}

	if files := form.File[fileUploadKey]; len(files) == 0 {
		err = throw.NewBadRequestErr(throw.ErrNoFile)
	} else if len(files) > 1 {
		err = throw.NewBadRequestErr(throw.ErrOverFile)
	} else {
		fileHeader = files[0]
	}
	if err != nil {
		ll.BadRequest(err)
		return cnt.com.Response(c, ll, err)
	}

	mediaObjectID, nested := cnt.processUploadingFile(ctx, thisUserID, fileHeader, note)
	if nested != nil {
		ll.Named("processUploadingFile").Nestedf(nested.Message())
		return cnt.fail.SendNested(c, "", nested)
	}

	// TODO сообщение в аудит о загрузке файла

	ll.With("id", mediaObjectID).Debug("file uploaded")

	return c.JSON(http.StatusOK, newFileUploadResponse(&it.MediaUploadFile{
		ID:          mediaObjectID,
		Typ:         "",
		PreviewLink: "",
		Note:        "",
	}))
}

//func (cnt *Controller)

//// Load docs
////
////	@Summary	Чтение произвольных клиентских данных
////	@Description
////	@Produce	json
////	@Param		store_code	path		string	true	"code store"
////	@Success	200			{object}	loadView
////	@Failure	400			{object}	fail.Response
////	@Router		/store/:store_path [GET]
////	@Tags		store
////	@Security	ApiKeyAuth
//func (cnt *Controller) Load(c echo.Context) error {
//	var (
//		ctx  = c.Request().Context()
//		ll   = cnt.logger.Named("Store")
//		form = new(storePathForm)
//	)
//
//	thisUserID, nested := cnt.com.ExtractUserAndForm(c, form)
//	if nested != nil {
//		ll.Named("ExtractForm").Nestedf(nested.Message())
//		return cnt.fail.SendNested(c, "", nested)
//	}
//
//	ll = ll.With("store_path", form.StorePath).With("thisUserID", thisUserID)
//
//	payload, err := cnt.storeAct.Load(ctx, thisUserID, form.StorePath)
//	if err != nil {
//		ll.Named("Load").Nested(err)
//
//		if errors.Is(err, action.ErrNotFound) {
//			return cnt.fail.Send(c, "", http.StatusNotFound, err)
//		}
//
//		return cnt.fail.SendInternalServerErr(c, "", err)
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
//	@Failure	400			{object}	fail.Response
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
//		return cnt.fail.SendNested(c, "", nested)
//	}
//
//	ll = ll.With("store_path", form.StorePath).With("thisUserID", thisUserID)
//
//	if err := cnt.storeAct.Delete(ctx, thisUserID, form.StorePath); err != nil {
//		ll.Named("Load").Nested(err)
//
//		if errors.Is(err, action.ErrNotFound) {
//			return cnt.fail.Send(c, "", http.StatusNotFound, err)
//		}
//
//		return cnt.fail.SendInternalServerErr(c, "", err)
//	}
//
//	ll.Debug("deleted")
//
//	return c.NoContent(http.StatusOK)
//}
