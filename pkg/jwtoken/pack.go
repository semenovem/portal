package jwtoken

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func (s *Service) NewPairTokens(params *TokenParams) (*PairTokens, error) {
	access := &AccessPayload{
		Expired:   time.Now().Add(s.accessLifetime).Unix(),
		UserID:    params.UserID,
		SessionID: params.SessionID,
	}

	refresh := &RefreshPayload{
		Expired:   time.Now().Add(s.refreshLifetime).Unix(),
		SessionID: params.SessionID,
	}

	return s.newPairTokens(access, refresh)
}

func (s *Service) newPairTokens(aPayload *AccessPayload, rPayload *RefreshPayload) (*PairTokens, error) {
	access, err := jwt.
		NewWithClaims(jwt.SigningMethodHS256, accessClaim(aPayload)).
		SignedString(s.accessSecret)
	if err != nil {
		return nil, err
	}

	refresh, err := jwt.
		NewWithClaims(jwt.SigningMethodHS256, refreshClaim(rPayload)).
		SignedString(s.refreshSecret)
	if err != nil {
		return nil, err
	}

	return &PairTokens{
		Access:  access,
		Refresh: refresh,
	}, nil
}

func accessClaim(t *AccessPayload) jwt.MapClaims {
	return jwt.MapClaims{
		claimExpired:   fmt.Sprintf("%d", t.Expired),
		claimSessionID: fmt.Sprintf("%d", t.SessionID),
		claimUserID:    fmt.Sprintf("%d", t.UserID),
	}
}

func refreshClaim(t *RefreshPayload) jwt.MapClaims {
	return jwt.MapClaims{
		claimExpired:   fmt.Sprintf("%d", t.Expired),
		claimSessionID: fmt.Sprintf("%d", t.SessionID),
		claimRefreshID: t.RefreshID.String(),
	}
}
