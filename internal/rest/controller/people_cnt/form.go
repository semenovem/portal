package people_cnt

import (
	"github.com/google/uuid"
	"github.com/semenovem/portal/internal/rest/controller"
	"time"
)

type UserForm struct {
	UserID uint32 `param:"user_id" validate:"required"`
}

type SearchForm struct {
	StartTime time.Time   `json:"start_time" query:"start_time" binding:"datetime=2006-01-02T15:04:05Z07:00" validate:"omitempty"`
	EndTime   time.Time   `json:"end_time" query:"end_time" binding:"datetime=2006-01-02T15:04:05Z07:00" validate:"omitempty"`
	UserIDs   []uuid.UUID `json:"user_id[]" query:"user_id[]" validate:"omitempty"`
	Slugs     []string    `json:"slug[]" query:"slug[]" validate:"omitempty"`
	controller.PaginationForm
}
