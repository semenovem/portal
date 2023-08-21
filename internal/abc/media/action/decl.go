package media_action

import (
	"github.com/semenovem/portal/internal/abc/media/provider"
	"github.com/semenovem/portal/pkg"
)

type MediaAction struct {
	logger   pkg.Logger
	mediaPvd *media_provider.MediaProvider
}

func New(
	logger pkg.Logger,
	mediaPvd *media_provider.MediaProvider,
) *MediaAction {
	return &MediaAction{
		logger:   logger.Named("MediaAction"),
		mediaPvd: mediaPvd,
	}
}
