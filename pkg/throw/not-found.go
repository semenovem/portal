package throw

var (
	Err404             = NewNotFoundErr("not found")
	Err404User         = NewNotFoundErr("user not found")
	Err404AuthSession  = NewNotFoundErr("auth session not found")
	Err404OnetimeEntry = NewNotFoundErr("onetime entry not found")
)

// NotFoundErr ошибки в результате отсутствия запрошенной сущности
type NotFoundErr interface {
	Error() string
	isNotFoundErr() bool
}

type notFoundErr struct {
	msg string
}

func NewNotFoundErr(msg string) error {
	return &notFoundErr{msg: msg}
}

func (e notFoundErr) Error() string {
	return e.msg
}

func (e notFoundErr) isNotFoundErr() bool {
	return true
}

func IsNotFoundErr(err error) bool {
	_, ok := err.(*notFoundErr)
	return ok
}
