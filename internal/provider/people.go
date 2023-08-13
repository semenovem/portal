package provider

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/semenovem/portal/pkg"
)

type PeoplePvd struct {
	db     *pgxpool.Pool
	logger pkg.Logger
}

func NewPeoplePvd(db *pgxpool.Pool, logger pkg.Logger) *PeoplePvd {
	return &PeoplePvd{
		db:     db,
		logger: logger.Named("peoplePvd"),
	}
}
