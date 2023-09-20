package media_action

import (
	"github.com/semenovem/portal/internal/abc/media/provider"
	"github.com/semenovem/portal/internal/s3"
	"github.com/semenovem/portal/pkg"
)

type MediaAction struct {
	logger            pkg.Logger
	s3                *s3.Service
	mediaPvd          *media_provider.MediaProvider
	avatarMaxSizeByte uint32
}

func New(
	logger pkg.Logger,
	s3 *s3.Service,
	mediaPvd *media_provider.MediaProvider,
) *MediaAction {
	return &MediaAction{
		logger:   logger.Named("MediaAction"),
		s3:       s3,
		mediaPvd: mediaPvd,
	}
}
