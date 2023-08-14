package jwtoken

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"strconv"
)

func (s *Service) GetAccessPayload(str string) (*AccessPayload, error) {
	rawToken, err := s.parse(&s.jwtParser, str, s.extractAccess)
	if err != nil {
		return nil, err
	}

	return s.extractAccessPayload(rawToken)
}

func (s *Service) GetRefreshPayload(str string) (*RefreshPayload, error) {
	rawToken, err := s.parse(&s.jwtParser, str, s.extractRefresh)
	if err != nil {
		return nil, err
	}

	return s.extractRefreshPayload(rawToken)
}

func (s *Service) parse(parser *jwt.Parser, str string, fn jwt.Keyfunc) (jwt.MapClaims, error) {
	if str == "" {
		return nil, ErrEmpty
	}

	t, err := parser.Parse(str, fn)
	if err != nil {
		err1, ok := err.(*jwt.ValidationError)
		if !ok {
			return nil, ErrInvalid
		}

		switch err1.Errors {
		case jwt.ValidationErrorExpired:
			return nil, ErrExpired
		case jwt.ValidationErrorClaimsInvalid:
			return nil, ErrClaimInvalid
		case jwt.ValidationErrorSignatureInvalid:
			return nil, ErrSignInvalid
		}

		return nil, ErrInvalid
	}

	if !t.Valid {
		return nil, ErrInvalid
	}

	return extractClaims(t)
}

func (s *Service) extractAccess(token *jwt.Token) (interface{}, error) {
	return getWJTSecret(s.accessSecret, token)
}

func (s *Service) extractRefresh(token *jwt.Token) (interface{}, error) {
	return getWJTSecret(s.refreshSecret, token)
}

func (s *Service) extractAccessPayload(claims jwt.MapClaims) (*AccessPayload, error) {
	sessionID, err := unpackUID32(claims, claimSessionID)
	if err != nil {
		return nil, err
	}

	userID, err := unpackUID32(claims, claimUserID)
	if err != nil {
		return nil, err
	}

	return &AccessPayload{
		Expired:   unpackExpired(claims),
		SessionID: sessionID,
		UserID:    userID,
	}, nil
}

func (s *Service) extractRefreshPayload(claims jwt.MapClaims) (*RefreshPayload, error) {
	sessionID, err := unpackUID32(claims, claimSessionID)
	if err != nil {
		return nil, err
	}

	refreshID, err := unpackUUID(claims, claimRefreshID)
	if err != nil {
		return nil, err
	}

	return &RefreshPayload{
		Expired:   unpackExpired(claims),
		SessionID: sessionID,
		RefreshID: refreshID,
	}, nil
}

func unpackUID32(claims jwt.MapClaims, field string) (uint32, error) {
	v, ok := claims[field]
	if !ok {
		return 0, fmt.Errorf("field [%s] not exists", field)
	}

	s, ok := v.(string)
	if !ok {
		return 0, errors.New("value is not a string")
	}

	i, err := strconv.ParseUint(s, 10, 32)

	return uint32(i), err
}

func unpackExpired(claims jwt.MapClaims) int64 {
	if s, ok := claims[claimExpired].(string); ok {
		exp, _ := strconv.ParseInt(s, 10, 64)
		return exp
	}

	return 0
}

func unpackUUID(claims jwt.MapClaims, field string) (uuid.UUID, error) {
	v, ok := claims[field]
	if !ok {
		return uuid.Nil, fmt.Errorf("field [%s] not exists", field)
	}

	s, ok := v.(string)
	if !ok {
		return uuid.Nil, fmt.Errorf("value [%+v][%T] is not a string", v, v)
	}

	id, err := uuid.Parse(s)
	if err != nil {
		return uuid.Nil, fmt.Errorf("value [%s] is not uuid", s)
	}

	return id, nil
}
