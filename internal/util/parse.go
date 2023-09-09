package util

import "time"

func ParsePointerStringToTime(s *string) (*time.Time, error) {
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
