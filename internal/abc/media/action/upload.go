package media_action

import (
	"context"
	"io"
)

func (a *MediaAction) Upload(
	ctx context.Context,
	thisUserID uint32,
	reader io.Reader,
) error {

	// TODO Проверить, может ли пользователь загружать файлы

	return nil
}
