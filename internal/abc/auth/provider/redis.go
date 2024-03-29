package auth_provider

import (
	"context"
	"fmt"
)

// IsSessionCanceled проверяет, закончилась ли сессия
func (p *AuthProvider) IsSessionCanceled(ctx context.Context, sessionID uint32) (bool, error) {
	i, err := p.redis.Exists(ctx, p.getSessionCancelKeyName(sessionID)).Result()
	if err != nil {
		p.logger.Func(ctx, "IsSessionCanceled").Redis(err)
		return false, err
	}

	return i != 0, nil
}

// IsSessionCanceled отозвать авторизованную сессию
func (p *AuthProvider) sessionCanceled(ctx context.Context, sessionID uint32) error {
	err := p.redis.Set(
		ctx,
		p.getSessionCancelKeyName(sessionID),
		"",
		p.config.Auth.JWT.AccessTokenLifetime.Val,
	).Err()
	if err != nil {
		p.logger.Func(ctx, "sessionCanceled").Redis(err)
		return err
	}

	return nil
}

func (p *AuthProvider) getSessionCancelKeyName(sessionID uint32) string {
	return fmt.Sprintf("session_cancel_%d", sessionID)
}
