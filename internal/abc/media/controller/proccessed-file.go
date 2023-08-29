package media_controller

import (
	"context"
	"github.com/semenovem/portal/pkg/failing"
	"github.com/semenovem/portal/pkg/it"
	"mime/multipart"
	"net/http"
)

func (cnt *Controller) processUploadingFile(
	ctx context.Context,
	thisUserID uint32,
	fh *multipart.FileHeader,
	note string,
) (*it.MediaFile, failing.Nested) {
	var (
		ll = cnt.logger.Named("FileUpload")
	)

	if fh.Size > cnt.maxUpload {
		ll.With("size", fh.Size).BadRequest(it.ErrFileTooBig)
		return nil, failing.NewNested(http.StatusBadRequest, it.ErrFileTooBig)
	}

	if fh.Size == 0 {
		ll.With("size", fh.Size).BadRequest(it.ErrFileEmpty)
		return nil, failing.NewNested(http.StatusBadRequest, it.ErrFileEmpty)
	}

	typ := fh.Header.Get(contentTypeKey)

	if _, ok := allowedContentTypes[typ]; !ok {
		ll.With("content-type", typ).BadRequest(it.ErrUnsupportedContent)
		return nil, failing.NewNested(http.StatusBadRequest, it.ErrUnsupportedContent)
	}

	f, err := fh.Open()
	if err != nil {
		ll.Named("fileOpen").ErrorE(err)
		return nil, failing.NewNested(http.StatusInternalServerError, err)
	}

	defer f.Close()

	// Проверка фактического типа файла
	imgBuff := make([]byte, 512)
	if _, err = f.Read(imgBuff); err != nil {
		ll.Named("fileRead").ErrorE(err)
		return nil, failing.NewNested(http.StatusInternalServerError, err)
	}

	typ = http.DetectContentType(imgBuff)

	if _, ok := allowedContentTypes[typ]; !ok {
		ll.Named("DetectContentType").BadRequest(it.ErrUnsupportedContent)
		return nil, failing.NewNested(http.StatusBadRequest, it.ErrUnsupportedContent)
	}

	_, err = f.Seek(0, 0)
	if err != nil {
		ll.Error(err.Error())
		ll.Named("Seek").ErrorE(err)
		return nil, failing.NewNested(http.StatusInternalServerError, err)
	}

	mf, err := cnt.mediaAct.Upload(ctx, thisUserID, f, note)
	if err != nil {
		ll.Named("Upload")
		switch err.(type) {
		case it.AccessErr:
			return nil, failing.NewNested(http.StatusForbidden, err)
		}

		return nil, failing.NewNested(http.StatusInternalServerError, err)
	}

	return mf, nil
}
