package s3

import (
	"context"
	"crypto/tls"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/semenovem/portal/config"
	"github.com/semenovem/portal/pkg"
	"net/http"
	"time"

	"github.com/minio/minio-go/v7"
)

const (
	docsBucketName    = "docs"
	imagesBucketName  = "images"
	videosBucketName  = "videos"
	preloadBucketName = "preload"
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
		return nil, ll.NestedWith(err, "can't create s3 minio-client")
	}

	buckets := []string{docsBucketName, imagesBucketName, videosBucketName, preloadBucketName}
	for _, n := range buckets {
		if err = o.createBucket(ctx, n); err != nil {
			return nil, ll.NestedWith(err, "can't create s3 buckets")
		}
	}

	return o, nil
}

func (s *Service) createBucket(ctx context.Context, name string) error {
	ll := s.logger.Named("createBucket").With("bucketName", name)

	if exists, err := s.s3Client.BucketExists(ctx, name); err != nil {
		ll.Named("BucketExists").Error(err.Error())
		return err
	} else if exists {
		return nil
	}

	err := s.s3Client.MakeBucket(ctx, name, minio.MakeBucketOptions{})
	if err != nil {
		ll.Named("MakeBucket").Error(err.Error())
		return err
	}

	return nil
}
