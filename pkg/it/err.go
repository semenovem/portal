package it

type ValidateErr interface {
	Error() string
	isValidateErr() bool
}

type validateErr struct {
	msg string
}

func (e validateErr) Error() string {
	return e.msg
}

func (e validateErr) isAccessErr() bool {
	return true
}

func IsValidateErr(err error) bool {
	_, ok := err.(*validateErr)
	return ok
}
