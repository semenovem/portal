package action

const (
	msgErrForbidden = "access denied"
)

var (
	ErrForbidden = newForbiddenErr("access denied")

	ErrNotFound = newNotFoundErr("not found")
)

type ForbiddenErr struct {
	msg string
}

func (e ForbiddenErr) Error() string {
	return e.msg
}

func IsForbiddenErr(err error) bool {
	_, ok := err.(*ForbiddenErr)
	return ok
}

func newForbiddenErr(msg string) *ForbiddenErr {
	return &ForbiddenErr{
		msg: msg,
	}
}

type NotFoundErr struct {
	msg string
}

func (e NotFoundErr) Error() string {
	return e.msg
}

func IsNotFoundErr(err error) bool {
	_, ok := err.(*NotFoundErr)
	return ok
}

func newNotFoundErr(msg string) *NotFoundErr {
	return &NotFoundErr{
		msg: msg,
	}
}
