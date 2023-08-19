package audit

type Code string

const (
	UserLoginDeny Code = "user-login-deny" // Авторизация по логопасс - запрет
	UserLogin     Code = "user-login"      // Авторизация по логопасс - успешно
	UserLogout    Code = "user-logout"     // Выход из системы

	AuthSessionDeny Code = "auth-session-deny" // Создана новая авторизованная сессия
	AuthSession     Code = "auth-session"      // Создана новая авторизованная сессия
)
