package people

import "errors"

var (
	ErrUserExpired      = errors.New("user expired")
	ErrUserFired        = errors.New("user fired")
	ErrUserNotStartWork = errors.New("user not start work")
	ErrUserNotActive    = errors.New("user have not active status")
)
