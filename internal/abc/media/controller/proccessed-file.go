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
) (uint32, fail.Nested) {
	var (
		ll = cnt.logger.Named("FileUpload")
	)

	if fh.Size > cnt.maxUpload {
		ll.With("size", fh.Size).BadRequest(throw.ErrFileTooBig)
		return 0, fail.NewNested(http.StatusBadRequest, throw.ErrFileTooBig)
	}

	if fh.Size == 0 {
		ll.With("size", fh.Size).BadRequest(throw.ErrFileEmpty)
		return 0, fail.NewNested(http.StatusBadRequest, throw.ErrFileEmpty)
	}

	contentType := fh.Header.Get(contentTypeKey)

	if _, ok := allowedContentTypes[contentType]; !ok {
		ll.With("content-type", contentType).BadRequest(throw.ErrUnsupportedContent)
		return 0, fail.NewNested(http.StatusBadRequest, throw.ErrUnsupportedContent)
	}

	binary, err := readFileHeader(fh)
	if err != nil {
		ll.Named("readFileHeader").ErrorE(err)
		return 0, fail.NewNested(http.StatusInternalServerError, err)
	}

	if nested := cnt.detectFactContentType(binary, contentType); nested != nil {
		ll.Named("detectFactContentType").Nestedf(nested.Message())
		return 0, nested
	}

	objType, err := it.MediaObjectByContentType(contentType)
	if err != nil {
		ll.Named("MediaObjectByContentType").BadRequest(err)
		return 0, fail.NewNested(http.StatusBadRequest, err)
	}

	uploadedFileID, err := cnt.mediaAct.Upload(ctx, thisUserID, objType, binary, note)
	if err != nil {
		ll.Named("Upload").Nested(err)

		switch err.(type) {
		case *throw.AccessErr:
			return 0, fail.NewNested(http.StatusForbidden, err)
		case *throw.BadRequestErr:
			return 0, fail.NewNested(http.StatusBadRequest, err)
		}

		return 0, fail.NewNested(http.StatusInternalServerError, err)
	}

	return uploadedFileID, nil
}
