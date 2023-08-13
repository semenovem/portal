package audit

const (
	Create Action = "create"
	Delete Action = "delete"
	Update Action = "update"
)

const (
	Allow Decision = "allow" // Действие выполнено
	Deny  Decision = "deny"  // Запрет на действие
)

const (
	Vehicle Domain = "vehicle"
	User    Domain = "user"
	Auth    Domain = "auth"
)

type Action string   // Типы действий в аудите
type Decision string // Решение на запрос
type Cause string    // Причина если было Deny
type Domain string

type P map[string]interface{}

type AuthArg struct {
	UserID uint32
}

type AuthArg2 struct {
	UserID uint32
}
