package throw

/* AuthErr ошибки в результате нарушения при авторизации */

var (
	ErrAuthPasswdIncorrect = NewAuthErr("password is incorrect")
	ErrAuthUserCannotLogin = NewAuthErr("user cannot login")
	ErrAuthRefreshUnknown  = NewAuthErr("refresh token data does not match")

	ErrAuthCookieEmpty = NewAuthErr("empty header [Authorization] token")
	ErrUserLogouted    = NewAuthErr("user is logouted")
	ErrAccessTokenExp  = NewAuthErr("access token expired")
	ErrInvalidBearer   = NewAuthErr("invalid bearer token")
)

type AuthErr interface {
	Error() string
	isAuthErr() bool
}

type authErr struct {
	msg    string
	target error
}

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

func (e authErr) Error() string {
	if e.target != nil {
		if e.msg == "" {
			return e.target.Error()
		}
		return e.target.Error() + ": " + e.msg
	}

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
