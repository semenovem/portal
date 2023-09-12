package util

import (
	"github.com/semenovem/portal/pkg/throw"
	"time"
)

// ParsePointerStrToTime пустая строка - пустой объект time.Time
func ParsePointerStrToTime(s *string) (*time.Time, error) {
	if s == nil {
		return nil, nil
	}

	if *s == "" {
		return &time.Time{}, nil
	}

	t, err := time.Parse(time.RFC3339, *s)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

// ParseNotEmptyStrToTime пустая строка - nil
func ParseNotEmptyStrToTime(s *string) (*time.Time, error) {
	if s == nil || *s == "" {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, *s)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

// MustParseToTime пустая строка - пустой объект
func MustParseToTime(s *string) *time.Time {
	if s == nil {
		return nil
	}

	if *s == "" {
		return &time.Time{}
	}

	t, err := time.Parse(time.RFC3339, *s)
	if err != nil {
		panic(throw.NewInvalidTimeErr(err))
	}

	return &t
}
