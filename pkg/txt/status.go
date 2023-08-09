package txt

import (
	"net/http"
)

var statuses = map[int]string{
	http.StatusBadRequest:          "Неправильный запрос",
	http.StatusUnauthorized:        "Вы не авторизованы",
	http.StatusForbidden:           "Запрещено",
	http.StatusNotFound:            "Объект не найден",
	http.StatusTooManyRequests:     "Слишком много запросов",
	http.StatusInternalServerError: "Внутренняя ошибка сервера",
}
