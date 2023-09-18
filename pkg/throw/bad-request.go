package throw

/* BadRequestErr ошибки в результате проверок */

var (
	Err400DuplicateLogin = NewBadRequestErrf("duplicate user login")
	Err400DuplicateEmail = NewBadRequestErrf("duplicate user email")

	Err400FiredBeforeStart = NewBadRequestErrf("time fired before start work")
	Err400FakeContentType  = NewBadRequestErrf("fake content type")
)

type BadRequestErr interface {
	Error() string
	isBadRequestErr() bool
	Target() error
}

type badRequestErr struct {
	msg    string
	target error
}

func NewBadRequestErr(err error) error {
	return &badRequestErr{target: err}
}

func NewBadRequestErrf(msg string) error {
	return &badRequestErr{msg: msg}
}

func IsBadRequestErr(err error) bool {
	_, ok := err.(*badRequestErr)
	return ok
}

func (e badRequestErr) Error() string {
	if e.target != nil {
		if e.msg == "" {
			return e.target.Error()
		}
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
