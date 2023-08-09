package router

import (
	"github.com/semenovem/portal/pkg/failing"
	"github.com/semenovem/portal/pkg/txt"
	"github.com/semenovem/portal/pkg/vld"
)

var validators = map[string]string{
	vld.UserLoginTag:    txt.ValidUserLogin,
	vld.UserPasswordTag: txt.ValidatePasswordWeak,
}

var (
	unknownFailing = failing.Message{
		Code:        10000,
		DefaultText: "Неизвестная ошибка",
		Translations: map[string]string{
			txt.EN: "Unknown error",
		},
	}

	invalidFailing = failing.Message{
		Code:        10001,
		DefaultText: "Невалидные параметры запроса",
		Translations: map[string]string{
			txt.EN: "Invalid request parameters",
		},
	}
)
