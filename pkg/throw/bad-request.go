package throw

/* BadRequestErr ошибки в результате проверок */

var (
	Err400DuplicateLogin = NewBadRequestErr("duplicate user login")
	Err400DuplicateEmail = NewBadRequestErr("duplicate user email")

	Err400FiredBeforeStart = NewBadRequestErr("time fired before start work")
)

func NewBadRequestErr(msg string) error {
	return &badRequestErr{msg: msg}
}

func IsBadRequestErr(err error) bool {
	_, ok := err.(*badRequestErr)
	return ok
}

type BadRequestErr interface {
	Error() string
	isBadRequestErr() bool
	Target() error
}

type badRequestErr struct {
	msg    string
	target error
}

func (e badRequestErr) Error() string {
	if e.target != nil {
		return e.target.Error() + ": " + e.msg
	}

	return e.msg
}

func (e badRequestErr) Target() error {
	if e.target != nil {
		return e.target
	}

	return e
}

func (e badRequestErr) isBadRequestErr() bool {
	return true
}
