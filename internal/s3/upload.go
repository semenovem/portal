package s3

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"io"
)

func (s *Service) UploadImage(
	ctx context.Context,
	file io.Reader,
	objectPath string,
	fileSize int64,
) error {
	uploadInfo, err := s.s3Client.PutObject(
		ctx,
		"images",
		objectPath,
		file,
		fileSize,
		minio.PutObjectOptions{},
	)
	if err != nil {
		s.logger.Named("UploadImage").With("objectPath", objectPath).Error(err.Error())
		return err
	}

	fmt.Println(">>>>>>>>> ", uploadInfo.Size)

	return nil
}
