package media_controller

import (
	"github.com/semenovem/portal/config"
	"github.com/semenovem/portal/internal/abc/controller"
	"github.com/semenovem/portal/internal/abc/media"
	"github.com/semenovem/portal/internal/abc/media/action"
	"github.com/semenovem/portal/internal/audit"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/fail"
	"github.com/semenovem/portal/pkg/throw"
	"mime/multipart"
	"net/http"
)

const (
	uploadNoteKey = "note"
	fileUploadKey = "file"

	contentTypeKey = "Content-Type"
)

type Controller struct {
	logger     pkg.Logger
	mainConfig *config.Platform
	fail       *fail.Service
	com        *controller.Common
	mediaAct   *media_action.MediaAction
	audit      *audit.AuditProvider
}

func New(
	arg *controller.InitArgs,
	mediaAct *media_action.MediaAction,
) *Controller {
	return &Controller{
		logger:     arg.Logger.Named("media-cnt"),
		mainConfig: arg.MainConfig,
		fail:       arg.FailureService,
		audit:      arg.Audit,
		com:        arg.Common,
		mediaAct:   mediaAct,
	}
}

var (
	allowedAvatarContentTypes = map[media.ObjectType]struct{}{
		media.ObjectPNG:  {},
		media.ObjectJPEG: {},
	}
	allowedDocContentTypes = map[media.ObjectType]struct{}{
		media.ObjectPNG:  {},
		media.ObjectJPEG: {},
		media.ObjectPDF:  {},
	}

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

func (cnt *Controller) detectRealContentType(byt []byte, declareContentType string) fail.Nested {
	realContentType := http.DetectContentType(byt)

	if realContentType != declareContentType {
		err := cnt.logger.With("declareContentType", declareContentType).
			With("realContentType", realContentType).
			BadRequestStrRetErr("fake content type")

		throw.NewBadRequestErr(throw.Err400FakeContentType)

		return fail.NewNested(http.StatusBadRequest, err)
	}

	return nil
}
