package audit

type Code string

const (
	UserLogin          Code = "user-login"           // Авторизация пользователя
	UserLogout         Code = "user-logout"          // Выход из системы
	CreateOnetimeEntry Code = "create-onetime-entry" // Создание авторизованной ссылки

	AuthSessionDeny Code = "auth-session-deny" // Создана новая авторизованная сессия
	AuthSession     Code = "auth-session"      // Создана новая авторизованная сессия

	EntityUser Code = "entity-user"

	UserToEmployee Code = "user-to-employee"

	FreeLoginNameCheck Code = "free-login-name-check" // Проверить, свободен ли логин

	EmployeeUpdate Code = "employee-update" // Обновление данных пользователя
)
