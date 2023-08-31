package media_controller

import (
	"context"
	"github.com/semenovem/portal/pkg/fail"
	"github.com/semenovem/portal/pkg/it"
	"github.com/semenovem/portal/pkg/throw"
	"mime/multipart"
	"net/http"
)

func (cnt *Controller) processUploadingFile(
	ctx context.Context,
	thisUserID uint32,
	fh *multipart.FileHeader,
	note string,
) (*it.MediaObject, fail.Nested) {
	var (
		ll = cnt.logger.Named("FileUpload")
	)

	if fh.Size > cnt.maxUpload {
		ll.With("size", fh.Size).BadRequest(throw.ErrFileTooBig)
		return nil, fail.NewNested(http.StatusBadRequest, throw.ErrFileTooBig)
	}

	if fh.Size == 0 {
		ll.With("size", fh.Size).BadRequest(throw.ErrFileEmpty)
		return nil, fail.NewNested(http.StatusBadRequest, throw.ErrFileEmpty)
	}

	contentType := fh.Header.Get(contentTypeKey)

	if _, ok := allowedContentTypes[contentType]; !ok {
		ll.With("content-type", contentType).BadRequest(throw.ErrUnsupportedContent)
		return nil, fail.NewNested(http.StatusBadRequest, throw.ErrUnsupportedContent)
	}

	f, err := fh.Open()
	if err != nil {
		ll.Named("fileOpen").ErrorE(err)
		return nil, fail.NewNested(http.StatusInternalServerError, err)
	}

	defer f.Close()

	// Проверка фактического типа файла
	imgBuff := make([]byte, 512)
	if _, err = f.Read(imgBuff); err != nil {
		ll.Named("fileRead").ErrorE(err)
		return nil, fail.NewNested(http.StatusInternalServerError, err)
	}

	insideContentType := http.DetectContentType(imgBuff)

	if _, ok := allowedContentTypes[insideContentType]; !ok {
		ll.Named("DetectContentType").BadRequest(throw.ErrUnsupportedContent)
		return nil, fail.NewNested(http.StatusBadRequest, throw.ErrUnsupportedContent)
	}

	if contentType != insideContentType {
		err = ll.With("http_content_type", contentTypeKey).
			With("insideContentType", insideContentType).
			BadRequestStrRetErr("fake content type")

		return nil, fail.NewNested(http.StatusBadRequest, err)
	}

	_, err = f.Seek(0, 0)
	if err != nil {
		ll.Named("Seek").ErrorE(err)
		return nil, fail.NewNested(http.StatusInternalServerError, err)
	}

	objType, err := it.MediaObjectByContentType(contentType)
	if err != nil {
		ll.Named("MediaObjectByContentType").BadRequest(err)
		return nil, fail.NewNested(http.StatusBadRequest, err)
	}

	mf, err := cnt.mediaAct.Upload(ctx, thisUserID, objType, f, fh.Size, note)
	if err != nil {
		ll.Named("Upload")
		switch err.(type) {
		case throw.AccessErr:
			return nil, fail.NewNested(http.StatusForbidden, err)
		}

		return nil, fail.NewNested(http.StatusInternalServerError, err)
	}

	return mf, nil
}
