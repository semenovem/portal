package people_action

import (
	"github.com/semenovem/portal/internal/provider/people_provider"
	"github.com/semenovem/portal/pkg"
)

type PeopleAction struct {
	logger    pkg.Logger
	peoplePvd *people_provider.PeopleProvider
}

func New(
	logger pkg.Logger,
	peoplePvd *people_provider.PeopleProvider,
) *PeopleAction {
	return &PeopleAction{
		logger:    logger.Named("PeopleAction"),
		peoplePvd: peoplePvd,
	}
}
