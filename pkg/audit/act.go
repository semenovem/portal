package audit

const (
	Create Action = "create"
	Delete Action = "delete"
	Update Action = "update"
	Deny   Action = "deny" // Запрет на действие
)

// Action Типы действий в аудите
type Action string
