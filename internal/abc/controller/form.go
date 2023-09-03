package controller

import (
	"github.com/semenovem/portal/internal/audit"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/fail"
)

type PaginationForm struct {
	Offset uint32 `query:"offset"`
	Limit  uint32 `query:"limit" validate:"number,max=100"`
}

func (f *PaginationForm) GetOffset() uint32 {
	if f.Offset <= 0 {
		return 0
	}

	return f.Offset
}

func (f *PaginationForm) GetLimit() uint32 {
	if f.Limit <= 0 {
		return 50
	}

	return f.Limit
}

type CntArgs struct {
	Logger         pkg.Logger
	FailureService *fail.Service
	Audit          *audit.AuditProvider
	Common         *Common
}
