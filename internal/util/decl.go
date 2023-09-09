package util

import (
	"strings"
	"time"
)

func ErrToPointerStr(err error) *string {
	if err == nil {
		return nil
	}

	s := err.Error()
	return &s
}

func ErrToStr(err error) string {
	if err == nil {
		return ""
	}

	return err.Error()
}

func ZeroStrNil(s *string) *string {
	if s == nil || *s == "" {
		return nil
	}

	return s
}

func ZeroArrStrNil(s *[]string) *[]string {
	if s == nil || len(*s) == 0 {
		return nil
	}

	return s
}

func ZeroTimeNil(s *time.Time) *time.Time {
	if s == nil || s.IsZero() {
		return nil
	}

	return s
}

func ZeroUint32Nil(s *uint32) *uint32 {
	if s == nil || *s == 0 {
		return nil
	}

	return s
}

func NormLowerStrToPointer(s string) *string {
	ss := strings.ToLower(strings.TrimSpace(s))
	return &ss
}
