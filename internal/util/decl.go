package util

func ErrToPointStr(err error) *string {
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
