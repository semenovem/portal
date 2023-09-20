package media_controller

import (
	"bytes"
	"context"
	"github.com/semenovem/portal/internal/abc/media"
	"github.com/semenovem/portal/pkg/fail"
	"io"
	"mime/multipart"
	"net/http"
)

func (cnt *Controller) processingUploading(
	ctx context.Context,
	fh *multipart.FileHeader,
	allow map[media.ObjectType]struct{},
) (media.ObjectType, io.Reader, fail.Nested) {
	var (
		ll          = cnt.logger.Func(ctx, "processingUploading")
		contentType = fh.Header.Get(contentTypeKey)
	)

	objectType, err := media.ObjectByContentType(contentType)
	if err != nil {
		ll.With("Content-Type", contentType).BadRequest(err)
		return "", nil, fail.NewNested(http.StatusBadRequest, err)
	}

	if _, ok := allow[objectType]; !ok {
		ll.With("objectType", objectType).BadRequest(err)
		return "", nil, fail.NewNested(http.StatusBadRequest, err)
	}

	if _, ok := allowedContentTypes[contentType]; !ok {
		ll.With("content-type", contentType).BadRequest(media.ErrContentObjectForbidden)
		return "", nil, fail.NewNested(http.StatusBadRequest, media.ErrContentObjectForbidden)
	}

	binary, err := readFileHeader(fh)
	if err != nil {
		ll.Named("readFileHeader").ErrorE(err)
		return "", nil, fail.NewNested(http.StatusInternalServerError, err)
	}

	if nested := cnt.detectRealContentType(binary, contentType); nested != nil {
		ll.Named("detectRealContentType").Nestedf(nested.Message())
		return "", nil, nested
	}

	return objectType, bytes.NewReader(binary), nil
}
