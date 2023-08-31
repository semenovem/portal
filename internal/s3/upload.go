package s3

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"io"
)

func (s *Service) UploadObject(
	ctx context.Context,
	file io.Reader,
	bucket string,
	objectPath string,
	fileSize int64,
) error {
	uploadInfo, err := s.s3Client.PutObject(
		ctx,
		bucket,
		objectPath,
		file,
		fileSize,
		minio.PutObjectOptions{},
	)
	if err != nil {
		s.logger.Named("UploadObject").With("objectPath", objectPath).Error(err.Error())
		return err
	}

	fmt.Println(">>>>>>>>> ", objectPath)
	fmt.Println(">>>>>>>>> ", uploadInfo.ChecksumSHA256)
	fmt.Println(">>>>>>>>> ", uploadInfo.ChecksumSHA1)

	return nil
}
