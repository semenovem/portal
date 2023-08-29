package media_action

import (
	"context"
	"github.com/semenovem/portal/pkg/it"
	"io"
)

func (a *MediaAction) Upload(
	ctx context.Context,
	thisUserID uint32,
	reader io.Reader,
	note string,
) (*it.MediaFile, error) {

	// TODO Проверить, может ли пользователь загружать файлы

	f := it.MediaFile{
		ID:          300,
		PreviewLink: "asdasdasd",
		Note:        "asdasfasdfsdf",
	}

	return &f, nil
}
