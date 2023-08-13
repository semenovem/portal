package provider

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/semenovem/portal/pkg"
	"github.com/semenovem/portal/pkg/audit"
)

type Audit struct {
	db     *pgxpool.Pool
	logger pkg.Logger
}

func NewAudit(db *pgxpool.Pool, logger pkg.Logger) *Audit {
	return &Audit{
		db:     db,
		logger: logger.Named("authPvd"),
	}
}

// Refusal отказ
func (a *Audit) Refusal(code audit.Code, cause audit.Cause, payload map[string]interface{}) {
	fmt.Printf("audit Refusal >>>>>>>>>> code=%s  cause=%s   payload:[%+v] \n", code, cause, payload)
}

// Approved успешно
func (a *Audit) Approved(code audit.Code, payload map[string]interface{}) {
	fmt.Printf("audit Approved >>>>>>>>>> code=%s  payload:[%+v] \n", code, payload)
}

// Send Аудит аутентификации
func (a *Audit) Send(code audit.Code, decision audit.Decision, payload map[string]interface{}) {
	fmt.Println("audit >>>>>>>>>> code = ", code)
	fmt.Println("audit >>>>>>>>>> decision = ", decision)
	fmt.Println("audit >>>>>>>>>> payload = ", payload)
}

// User Аудит действий пользователя
func (a *Audit) User(userID uint32, action audit.Action, payload map[string]interface{}) {
	fmt.Println("audit >>>>>>>>>> userID = ", userID)
	fmt.Println("audit >>>>>>>>>> decision = ", action)
	fmt.Println("audit >>>>>>>>>> payload = ", payload)
}

// Аудит действий пользователя
