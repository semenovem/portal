package jwtoken

import (
	"errors"
	"github.com/google/uuid"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	claimSessionID = "sid"
	claimExpired   = "exp"
	claimUserID    = "uid"
	claimRefreshID = "rid"
)

var (
	ErrInvalid      = errors.New("jwtoken: is invalid")
	ErrClaimInvalid = errors.New("jwtoken: claim is invalid")
	ErrSignInvalid  = errors.New("jwtoken: signature is invalid")
	ErrEmpty        = errors.New("jwtoken: empty")
	ErrExpired      = errors.New("jwtoken: expired")
)

type Service struct {
	accessSecret            []byte
	refreshSecret           []byte
	jwtParser               jwt.Parser
	jwtParserSkipValidation jwt.Parser
	accessLifetime          time.Duration
	refreshLifetime         time.Duration
}

type Config struct {
	AccessTokenSecret    string
	RefreshTokenSecret   string
	AccessTokenLifetime  time.Duration
	RefreshTokenLifetime time.Duration
}

type TokenParams struct {
	SessionID uint32
	UserID    uint32
	RefreshID uuid.UUID
}

type AccessPayload struct {
	SessionID uint32
	UserID    uint32
	Expired   int64
}

type RefreshPayload struct {
	SessionID uint32
	Expired   int64
	RefreshID uuid.UUID
}

type PairTokens struct {
	Access  string
	Refresh string
}
