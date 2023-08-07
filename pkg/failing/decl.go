package failing

const (
	fieldNameErr       = "__error__" // Имя поля в которое добавляется содержимое ошибки в dev режиме

	invalidRequestText = "Invalid request parameters"
)

type ValidationError struct {
	Path    string `json:"path"`
	Message string `json:"message"`
}

type Response struct {
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	ValidationErrors []*ValidationError     `json:"validation_errors,omitempty"`
	AdditionalFields map[string]interface{} `json:"additional_fields,omitempty"`
}

// Args аргументы для шаблона fmt.Sprint(...)
type Args []interface{}

type logger interface {
	Errorf(template string, args ...interface{})
}
