package failing

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

// SendValidationErr ответ клиенту ошибкой валидации
func (s *Service) SendValidationErr(c echo.Context, language string, err error) error {
	return c.JSON(http.StatusBadRequest, s.newValidationResponse(c, language, err))
}

func (s *Service) newValidationResponse(c echo.Context, language string, err error) *Response {
	var (
		resp = s.newResponse(language, &parsedOpt{
			message: s.invalidRequestMessage,
		})
	)

	if errs, ok := err.(validator.ValidationErrors); ok {
		resp.ValidationErrors = s.validationErrors(language, errs)
	} else {
		if resp.AdditionalFields == nil {
			resp.AdditionalFields = make(map[string]interface{})
		}

		resp.AdditionalFields[fieldNameErr] = err.Error()
	}

	return resp
}

func (s *Service) validationErrors(lang string, errs []validator.FieldError) []*ValidationError {
	validationErrs := make([]*ValidationError, 0)

	for _, fieldError := range errs {
		f := ValidationError{Path: toSnakeCase(fieldError.StructField())}

		if msg := s.findMessageByValidationMessageTag(fieldError.Tag()); msg == nil {
			f.Message = fieldError.Translate(s.translatorDefault)
		} else {
			f.Message = msg.Text(lang)
		}

		validationErrs = append(validationErrs, &f)
	}

	return validationErrs
}

func (s *Service) findMessageByValidationMessageTag(tag string) *Message {
	if msgKey, ok := s.validationMessageMap[tag]; ok {
		return s.messages[msgKey]
	}

	return nil
}
