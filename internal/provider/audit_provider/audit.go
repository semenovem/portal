package audit_provider

import (
	"fmt"
	"github.com/semenovem/portal/pkg/audit"
)

// Refusal отказ
func (a *AuditProvider) Refusal(code audit.Code, cause audit.Cause, payload map[string]interface{}) {
	a.input <- &auditPipe{
		code:    code,
		cause:   cause,
		payload: payload,
	}
}

// Approved успешно
func (a *AuditProvider) Approved(code audit.Code, payload map[string]interface{}) {
	a.input <- &auditPipe{
		code:    code,
		payload: payload,
	}
}

// Send Аудит аутентификации
func (a *AuditProvider) Send(code audit.Code, decision audit.Decision, payload map[string]interface{}) {
	fmt.Println("audit >>>>>>>>>> code = ", code)
	fmt.Println("audit >>>>>>>>>> decision = ", decision)
	fmt.Println("audit >>>>>>>>>> payload = ", payload)
}

// Auth Аудит действий пользователя
func (a *AuditProvider) Auth(userID uint32, code audit.Code, payload map[string]interface{}) {
	a.input <- &auditPipe{
		userID:  userID,
		code:    code,
		payload: payload,
	}
}
