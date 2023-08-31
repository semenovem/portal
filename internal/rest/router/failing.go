package router

import (
	"github.com/semenovem/portal/pkg/fail"
	"github.com/semenovem/portal/pkg/txt"
)

var validators = map[string]string{
	userLoginTag:    txt.ValidUserLogin,
	userPasswordTag: txt.ValidatePasswordWeak,
}

var (
	unknownFail = fail.Message{
		Code:        10000,
		DefaultText: "Неизвестная ошибка",
		Translations: map[string]string{
			txt.EN: "Unknown error",
		},
	}

	invalidFail = fail.Message{
		Code:        10001,
		DefaultText: "Невалидные параметры запроса",
		Translations: map[string]string{
			txt.EN: "Invalid request parameters",
		},
	}
)
