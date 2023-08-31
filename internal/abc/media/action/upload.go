package media_action

import (
	"context"
	"github.com/semenovem/portal/pkg/it"
	"io"
)

func (a *MediaAction) Upload(
	ctx context.Context,
	thisUserID uint32,
	mediaObj it.MediaObjectType,
	reader io.Reader,
	size int64,
	note string,
) (*it.MediaObject, error) {

	// TODO Проверить, может ли пользователь загружать файлы

	f := it.MediaObject{
		ID:          300,
		PreviewLink: "asdasdasd",
		Note:        "asdasfasdfsdf",
	}

	//err := throw.NewAccessErr("sdfsdfsf")

	a.s3.UploadObject(ctx, reader, "images", "/11/1", size)

	return &f, nil
}
