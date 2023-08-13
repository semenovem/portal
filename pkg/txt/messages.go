package txt

const (
	ValidUserLogin       = "10004 Логин пользователя может содержать символы и цифры"
	ValidatePasswordWeak = "10005 Пароль должен содержать хотя бы одну букву в нижнем регистре, одну букву в верхнем регистре, одну цифру, один спецсимвол, и не содержать пробелов. Длинной 8-20 символов"

	ValidateEmailNotSpecified  = "10012 Не указан электронный адрес"
	ValidateEmailInvalid       = "10008 Необходимо указать валидный электронный адрес, длиной 4-64 символов"
	ValidateUserNameLength     = "10005 Имя пользователя должно быть 3-64 символов"
	ValidateUserPositionLength = "10030 Позиция пользователя должно быть 3-64 символов"
	ValidateUserNameInvalid    = "10004 Имя пользователя содержит недопустимые символы"
	ValidatePINInvalid         = "10007 ПИН должен состоять из 4-х цифр"
)

const (
	NotFoundErrEntity = "10100 Запрашиваемый объект не найден"
	NotFoundMethod    = "10101 Метод не найден"
	NotFoundUser      = "10102 Пользователь не найден"
)

const (
	AuthInvalidLogoPasswd = "10200 Не верный логин или пароль"
	AuthWrongPIN          = "10212 Неверный ПИН"
)

//nolint:lll
var messages = map[string]*struct {
	en string
}{
	ValidUserLogin:             {},
	ValidateUserPositionLength: {},
	ValidateEmailNotSpecified:  {},
	NotFoundMethod:             {},
	NotFoundErrEntity:          {},
	NotFoundUser:               {},
	AuthInvalidLogoPasswd:      {},
	ValidateUserNameInvalid:    {},
	ValidateUserNameLength:     {},
	ValidatePasswordWeak:       {},
	ValidatePINInvalid:         {},
	ValidateEmailInvalid:       {},
	AuthWrongPIN:               {},
}
