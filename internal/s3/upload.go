package s3

import (
	"bytes"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
)

func (s *Service) ExistsFile(
	ctx context.Context,
	bucket string,
	object string,
) (bool, error) {
	_, err := s.s3Client.StatObject(ctx, bucket, object, minio.StatObjectOptions{})

	if err != nil {
		errResponse := minio.ToErrorResponse(err)

		switch errResponse.Code {
		//case accessDenied, noSuchBucket, invalidBucketName:
		//	return false, fmt.Errorf("no exists: code=%s, message=%s", errResponse.Code, errResponse.Message)
		case noSuchKey:
			return false, nil
		}

		return false, fmt.Errorf("no exists: code=%s, message=%s", errResponse.Code, errResponse.Message)
	}

	return true, nil
}

func (s *Service) UploadFile(
	ctx context.Context,
	byt []byte,
	bucket string,
	object string,
) error {
	if exists, err := s.ExistsFile(ctx, bucket, object); err != nil {
		return err
	} else if exists {
		return nil
	}

	return s.uploadFile(ctx, byt, bucket, object)
}

func (s *Service) uploadFile(
	ctx context.Context,
	byt []byte,
	bucket string,
	object string,
) error {
	_, err := s.s3Client.PutObject(
		ctx,
		bucket,
		object,
		bytes.NewReader(byt),
		int64(len(byt)),
		minio.PutObjectOptions{},
	)
	if err != nil {
		s.logger.Func(ctx, "UploadFile").With("objectPath", object).Error(err.Error())
		return err
	}

	return nil
}
