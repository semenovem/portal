package throw

/* AuthErr ошибки в результате нарушения при авторизации */

const (
	MsgUserCantLogin = "user cant login"
)

var (
	ErrAuthPasswdIncorrect = NewAuthErr("password is incorrect")
	ErrAuthUserNotWorks    = NewAuthErr(MsgUserCantLogin)
	ErrAuthRefreshUnknown  = NewAuthErr("refresh token data does not match")

	ErrAuthCookieEmpty = NewAuthErr("empty header [Authorization] token")
	ErrUserLogouted    = NewAuthErr("user is logouted")
	ErrAccessTokenExp  = NewAuthErr("access token expired")
	ErrInvalidBearer   = NewAuthErr("invalid bearer token")
)

func NewAuthErr(msg string, prevErrMsg ...string) error {
	if len(prevErrMsg) != 0 {
		for _, s := range prevErrMsg {
			msg += ": " + s
		}
	}

	return &authErr{msg: msg}
}

func IsAuthErr(err error) bool {
	_, ok := err.(*authErr)
	return ok
}

type AuthErr interface {
	Error() string
	isAuthErr() bool
}

type authErr struct {
	msg    string
	target error
}

func (e authErr) Error() string {
	return e.msg
}

func (e authErr) Target() error {
	if e.target != nil {
		return e.target
	}

	return e
}

func (e authErr) isAuthErr() bool {
	return true
}
