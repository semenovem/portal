package controller

import (
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/failing"
)

type PaginationForm struct {
	Offset int `query:"offset"`
	Limit  int `query:"limit" validate:"number,max=100"`
}

func (f *PaginationForm) GetOffset() int {
	if f.Offset <= 0 {
		return 0
	}

	return f.Offset
}

func (f *PaginationForm) GetLimit() int {
	if f.Limit <= 0 {
		return 50
	}

	return f.Limit
}

type CntArgs struct {
	Logger  pkg.Logger
	Failing *failing.Service
	Act     *Action
}
