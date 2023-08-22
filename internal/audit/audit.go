package audit

import (
	"fmt"
)

// Refusal отказ
func (a *AuditProvider) Refusal(code Code, cause Cause, payload map[string]interface{}) {
	a.input <- &auditPipe{
		code:    code,
		cause:   cause,
		payload: payload,
	}
}

// Approved успешно
func (a *AuditProvider) Approved(code Code, payload map[string]interface{}) {
	a.input <- &auditPipe{
		code:    code,
		payload: payload,
	}
}

// Send Аудит аутентификации
func (a *AuditProvider) Send(code Code, decision Decision, payload map[string]interface{}) {
	fmt.Println("audit >>>>>>>>>> code = ", code)
	fmt.Println("audit >>>>>>>>>> decision = ", decision)
	fmt.Println("audit >>>>>>>>>> payload = ", payload)
}

// Auth Аудит авторизаций пользователя
func (a *AuditProvider) Auth(userID uint32, code Code, payload map[string]interface{}) {
	a.input <- &auditPipe{
		userID:  userID,
		code:    code,
		payload: payload,
	}
}

// Oper Аудит действий пользователя
func (a *AuditProvider) Oper(userID uint32, code Code, action Action, payload map[string]interface{}) {
	a.input <- &auditPipe{
		userID:  userID,
		code:    code,
		action:  action,
		payload: payload,
	}
}
