package media_controller

import (
	"github.com/semenovem/portal/internal/abc/controller"
	"github.com/semenovem/portal/internal/abc/media"
	"github.com/semenovem/portal/internal/abc/media/action"
	"github.com/semenovem/portal/internal/audit"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/fail"
	"mime/multipart"
	"net/http"
)

const (
	uploadNoteKey = "note"
	fileUploadKey = "file"

	contentTypeKey = "Content-Type"
)

type Controller struct {
	logger   pkg.Logger
	fail     *fail.Service
	com      *controller.Common
	mediaAct *media_action.MediaAction
	audit    *audit.AuditProvider
	config   *media.ConfigMedia
}

func New(
	arg *controller.InitArgs,
	config *media.ConfigMedia,
	mediaAct *media_action.MediaAction,
) *Controller {
	return &Controller{
		logger:   arg.Logger.Named("media-cnt"),
		fail:     arg.FailureService,
		audit:    arg.Audit,
		com:      arg.Common,
		config:   config,
		mediaAct: mediaAct,
	}
}

var (
	allowedContentTypes = map[string]struct{}{
		"image/png":       {},
		"image/jpeg":      {},
		"application/pdf": {},
	}
)

func readFileHeader(fh *multipart.FileHeader) ([]byte, error) {
	f, err := fh.Open()
	if err != nil {
		return nil, err
	}

	defer f.Close()

	byt := make([]byte, fh.Size)
	_, err = f.Read(byt)

	return byt, err
}

func (cnt *Controller) detectFactContentType(byt []byte, declareContentType string) fail.Nested {
	factContentType := http.DetectContentType(byt)

	if factContentType != declareContentType {
		err := cnt.logger.With("declareContentType", declareContentType).
			With("factContentType", factContentType).
			BadRequestStrRetErr("fake content type")

		return fail.NewNested(http.StatusBadRequest, err)
	}

	return nil
}
