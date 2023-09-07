package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/semenovem/portal/internal/abc/controller"
	"github.com/semenovem/portal/pkg/it"
)

type customValidator struct {
	validate *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.validate.Struct(i)
}

var strValidators = map[string]func(string) error{
	controller.UserLoginVldTag:  it.ValidateUserLogin,
	controller.UserPasswordTag:  it.ValidateUserPassword,
	controller.UserNameVldTag:   it.ValidateUserName,
	controller.UserStatusVldTag: it.ValidateUserStatus,
	controller.UserRoleVldTag:   it.ValidateUserRole,
}

var arrStrValidators = map[string]func([]string) error{
	controller.UserRolesVldTag: it.ValidateUserRoles,
}

func newValidation() (echo.Validator, error) {
	val := validator.New()

	// Строковые валидаторы
	for k, v := range strValidators {
		if err := val.RegisterValidation(k, buildStrValidator(v)); err != nil {
			return nil, err
		}
	}

	//  Валидаторы массива строк
	for k, v := range arrStrValidators {
		if err := val.RegisterValidation(k, buildArrStrValidator(v)); err != nil {
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

func buildArrStrValidator(h func([]string) error) func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		if s, ok := fl.Field().Interface().([]string); ok {
			return h(s) == nil
		}

		return false
	}
}