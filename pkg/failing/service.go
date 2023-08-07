package failing

import (
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Config struct {
	IsDevMode             bool // Режим разработки
	Logger                logger
	TranslatorDefault     ut.Translator
	Translators           map[string]ut.Translator
	Messages              map[string]*Message
	ValidationMessageMap  map[string]string
	HTTPStatuses          map[int]*Message
	UnknownMessage        Message // Сообщения для неизвестных ошибок
	InvalidRequestMessage Message // Сообщение для ошибок валидации
}

type Service struct {
	translatorDefault     ut.Translator
	isDev                 bool
	messages              map[string]*Message
	validationMessageMap  map[string]string
	logger                logger
	httpStatuses          map[int]*Message
	unknownMessage        *Message
	invalidRequestMessage *Message
}

func New(c *Config) *Service {
	s := &Service{
		translatorDefault:     c.TranslatorDefault,
		isDev:                 c.IsDevMode,
		messages:              c.Messages,
		validationMessageMap:  c.ValidationMessageMap,
		logger:                c.Logger,
		unknownMessage:        &c.UnknownMessage,
		invalidRequestMessage: &c.InvalidRequestMessage,
		httpStatuses:          c.HTTPStatuses,
	}

	return s
}

// NewResponse
// opts, параметры, определяемые по типу:
// > map[string]interface{} - дополнительные поля
// > error - ошибка, если isDev = true (dev режим) добавить информацию в доп поля
// > []ValidationError - ошибки валидации
func (s *Service) NewResponse(language string, opts ...interface{}) *Response {
	return s.newResponse(language, s.parseOpts(opts))
}

func (s *Service) Send(c echo.Context, language string, httpStatus int, opts ...interface{}) error {
	opt := s.parseOpts(opts)

	if opt.message == nil {
		if st, ok := s.httpStatuses[httpStatus]; ok {
			opt.message = st
		} else {
			opt.message = s.unknownMessage
		}
	}

	return c.JSON(httpStatus, s.newResponse(language, opt))
}

func (s *Service) SendInternalServerErr(c echo.Context, language string, err error) error {
	return s.Send(c, language, http.StatusInternalServerError, err)
}

// SendNested обработка ответа из вложенных вызовов
func (s *Service) SendNested(c echo.Context, language string, nestedResp Nested) error {
	var (
		opt              = s.parseOpts(nestedResp.getOpts())
		validationErrors []*ValidationError
		validationErr    = nestedResp.getValidationErr()
	)

	if validationErr != nil {
		if errs, ok := validationErr.(validator.ValidationErrors); ok {
			validationErrors = s.validationErrors(language, errs)
			opt.message = s.invalidRequestMessage
		}
	}

	if opt.message == nil {
		if st, ok := s.httpStatuses[nestedResp.getHTTPStatusCode()]; ok {
			opt.message = st
		} else {
			opt.message = s.unknownMessage
		}
	}

	response := s.newResponse(language, opt)

	if len(validationErrors) != 0 {
		response.ValidationErrors = validationErrors
	}

	return c.JSON(nestedResp.getHTTPStatusCode(), response)
}

func (s *Service) newResponse(lang string, opt *parsedOpt) *Response {
	if opt.message == nil {
		opt.message = s.unknownMessage
	}

	txt := opt.message.Text(lang)
	if len(opt.args) != 0 {
		txt = fmt.Sprintf(txt, opt.args...)
	}

	return &Response{
		Code:             opt.message.Code,
		Message:          txt,
		AdditionalFields: opt.additionalFields,
	}
}

func (s *Service) TextDefault(key string) string {
	return s.msg(key, "")
}

func (s *Service) TextFromMsgKey(key string, language string) string {
	return s.msg(key, language)
}

func (s *Service) msg(key string, lang string) string {
	if msg, ok := s.messages[key]; ok {
		msg.Text(lang)
	}

	return s.unknownMessage.Text(lang)
}
