package audit

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

// Del Аудит действий пользователя - удаление
func (a *AuditProvider) Del(userID uint32, code Code, payload map[string]interface{}) {
	a.input <- &auditPipe{
		userID:  userID,
		code:    code,
		action:  Delete,
		payload: payload,
	}
}

// Get Аудит действий пользователя
func (a *AuditProvider) Get(userID uint32, code Code, payload map[string]interface{}) {
	a.input <- &auditPipe{
		userID:  userID,
		code:    code,
		action:  actionGet,
		payload: payload,
	}
}
