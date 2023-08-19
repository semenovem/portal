package auth_provider

import (
	"context"
	"fmt"
)

// IsSessionCanceled проверяет, закончилась ли сессия
func (p *AuthProvider) IsSessionCanceled(ctx context.Context, sessionID uint32) (bool, error) {
	i, err := p.redis.Exists(ctx, p.getSessionCancelKeyName(sessionID)).Result()
	if err != nil {
		p.logger.Named("IsSessionCanceled").RedisTag().Error(err.Error())
		return false, err
	}

	return i != 0, nil
}

// IsSessionCanceled отозвать сессию
func (p *AuthProvider) sessionCanceled(ctx context.Context, sessionID uint32) error {
	err := p.redis.Set(ctx, p.getSessionCancelKeyName(sessionID), "", p.jwtAccessTokenLifetimeMin).Err()
	if err != nil {
		p.logger.Named("sessionCanceled").RedisTag().Error(err.Error())
		return err
	}

	return nil
}

func (p *AuthProvider) getSessionCancelKeyName(sessionID uint32) string {
	return fmt.Sprintf("session_cancel_%d", sessionID)
}
