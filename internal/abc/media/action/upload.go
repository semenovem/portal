package media_action

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"github.com/semenovem/portal/internal/abc/media"
	"github.com/semenovem/portal/internal/s3"
)

func (a *MediaAction) Upload(
	ctx context.Context,
	thisUserID uint32,
	mediaObj media.ObjectType,
	binary []byte,
	note string,
) (uint32, error) {
	var (
		ll = a.logger.Func(ctx, "Upload")
	)

	// TODO Проверить, может ли пользователь загружать файлы

	var (
		hash       = sha1.Sum(binary)
		objectPath = hex.EncodeToString(hash[:]) + "." + string(mediaObj)
	)

	if err := a.s3.UploadFile(ctx, binary, s3.UploadedBucketName, objectPath); err != nil {
		ll.Named("UploadFile").With("bucket", s3.UploadedBucketName).With("objectPath", objectPath)
		return 0, err
	}

	//
	uploadedFileID, err := a.mediaPvd.CreateUploadedFile(ctx, objectPath, note, mediaObj)
	if err != nil {
		ll.Named("CreateUploadedFile").Nested(err)
		return 0, err
	}

	return uploadedFileID, nil
}
