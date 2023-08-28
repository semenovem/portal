package s3

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/semenovem/portal/config"
	"github.com/semenovem/portal/pkg"
	"net/http"
	"time"

	"github.com/minio/minio-go/v7"
)

type Props struct {
	Ctx    context.Context
	Logger pkg.Logger
	S3Conn *config.S3Conn
}

type Service struct {
	logger   pkg.Logger
	s3Client *minio.Client
}

func New(config *Props) (*Service, error) {
	var (
		o = &Service{
			logger: config.Logger.Named("Service"),
		}

		ll = o.logger.Named("New")

		conn        = config.S3Conn
		err         error
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*10)

		transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: conn.InsecureSkipVerify,
			},
		}

		options = &minio.Options{
			Creds:     credentials.NewStaticV4(conn.AccessKey, conn.SecretKey, ""),
			Transport: transport,
			Secure:    conn.UseSSL,
		}
	)

	defer cancel()

	if o.s3Client, err = minio.New(conn.URL, options); err != nil {
		ll.Named("minio.New").Error(err.Error())
		return nil, err
	}

	exists, err := o.s3Client.BucketExists(ctx, conn.BucketName)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, fmt.Errorf("bucket %s does not exists", conn.BucketName)
	}

	if err = o.createBuckets(ctx); err != nil {
		ll.Named("createBuckets").Error(err.Error())
		return nil, err
	}

	return o, nil
}

func (s *Service) createBuckets(ctx context.Context) error {
	var (
		ll  = s.logger.Named("createBuckets")
		opt = minio.MakeBucketOptions{
			Region:        "",
			ObjectLocking: false,
		}
	)

	if err := s.s3Client.MakeBucket(ctx, "images", opt); err != nil {
		ll.Named("MakeBucket").With("name", "images")
		return err
	}

	if err := s.s3Client.MakeBucket(ctx, "videos", opt); err != nil {
		ll.Named("MakeBucket").With("name", "videos")
		return err
	}

	if err := s.s3Client.MakeBucket(ctx, "docs", opt); err != nil {
		ll.Named("MakeBucket").With("name", "docs")
		return err
	}

	if err := s.s3Client.MakeBucket(ctx, "preload", opt); err != nil {
		ll.Named("MakeBucket").With("name", "preload")
		return err
	}

	return nil
}
