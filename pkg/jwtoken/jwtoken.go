package jwtoken

import (
	"github.com/golang-jwt/jwt"
)

func New(cfg *Config) *Service {
	return &Service{
		accessSecret:    []byte(cfg.AccessTokenSecret),
		refreshSecret:   []byte(cfg.RefreshTokenSecret),
		accessLifetime:  cfg.AccessTokenLifetime,
		refreshLifetime: cfg.RefreshTokenLifetime,
		jwtParser:       jwt.Parser{UseJSONNumber: false, SkipClaimsValidation: true},
	}
}

func getWJTSecret(b []byte, t *jwt.Token) (interface{}, error) {
	if _, matchHMAC := t.Method.(*jwt.SigningMethodHMAC); !matchHMAC {
		return nil, ErrSignInvalid
	}

	return b, nil
}

func extractClaims(token *jwt.Token) (jwt.MapClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrClaimInvalid
	}

	return claims, nil
}
