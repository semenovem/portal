package util

import "time"

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
