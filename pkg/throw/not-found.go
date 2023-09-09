package throw

/* NotFoundErr ошибки в результате отсутствия запрошенной сущности */

var (
	Err404             = NewNotFoundErr("not found")
	Err404User         = NewNotFoundErr("user not found")
	Err404AuthSession  = NewNotFoundErr("auth session not found")
	Err404OnetimeEntry = NewNotFoundErr("onetime entry not found")
)

func NewNotFoundErr(msg string) error {
	return &notFoundErr{msg: msg}
}

func IsNotFoundErr(err error) bool {
	_, ok := err.(*notFoundErr)
	return ok
}

type NotFoundErr interface {
	Error() string
	isNotFoundErr() bool
}

type notFoundErr struct {
	msg    string
	target error
}

func (e notFoundErr) Error() string {
	return e.msg
}

func (e notFoundErr) Target() error {
	if e.target != nil {
		return e.target
	}

	return e
}

func (e notFoundErr) isNotFoundErr() bool {
	return true
}
