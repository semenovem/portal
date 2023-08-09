package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/pkg/entity"
	"github.com/semenovem/portal/pkg/vld"
)

type customValidator struct {
	validate *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.validate.Struct(i)
}

var strValidators = map[string]func(string) error{
	vld.UserLoginTag:    entity.ValidateUserLogin,
	vld.UserPasswordTag: entity.ValidateUserPassword,
}

func newValidation() (echo.Validator, error) {
	val := validator.New()

	// Строковые валидаторы
	for k, v := range strValidators {
		err := val.RegisterValidation(k, buildStrValidator(v))
		if err != nil {
			return nil, err
		}
	}

	return &customValidator{validate: val}, nil
}

func buildStrValidator(h func(string) error) func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		if s, ok := fl.Field().Interface().(string); ok {
			return h(s) == nil
		}

		return false
	}
}
