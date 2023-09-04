package throw

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

// AuthErr ошибки в результате нарушения при авторизации
type AuthErr interface {
	Error() string
	isAuthErr() bool
}

type authErr struct {
	msg string
}

func NewAuthErr(msg string, prevErrMsg ...string) error {
	if len(prevErrMsg) != 0 {
		for _, s := range prevErrMsg {
			msg += ": " + s
		}
	}

	return &authErr{msg: msg}
}

func (e authErr) Error() string {
	return e.msg
}

func (e authErr) isAuthErr() bool {
	return true
}

func IsAuthErr(err error) bool {
	_, ok := err.(*authErr)
	return ok
}
