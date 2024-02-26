package throw

import (
	"errors"
	"github.com/semenovem/portal/pkg/throw/throwtrace"
	"regexp"
)

type Kind int

const (
	Unknown    = iota
	Deny       // Нет прав у пользователя
	Duplicate  // Ограничение уникальности (например в БД)
	NotFound   // Сущность не найдена
	BadRequest // Не корректный запрос
	Auth       // Ошибка авторизации
)

const (
	reasonNameLenMax        = 63
	reasonMetadataKeyLenMax = 64
	regStrReasonName        = "^[A-Z][A-Z0-9_]+[A-Z0-9]$"
	regStrReasonMetadataKey = "^[a-zA-Z0-9-_]+$"
)

var _ = func() struct{} {
	throwtrace.XDoNotUse{}.SetExtractTraceData(func(err error) ([]*throwtrace.Point, map[string]any) {
		return Cast(err).traceData()
	})
	throwtrace.XDoNotUse{}.SetExtractDesc(func(err error) (string, []any) { return Cast(err).getDesc() })
	throwtrace.XDoNotUse{}.SetExtractReasons(func(err error) []*throwtrace.Reason { return Cast(err).getReasons() })

	return struct{}{}
}()

var (
	reasonNameReg        = regexp.MustCompile(regStrReasonName)
	reasonMetadataKeyReg = regexp.MustCompile(regStrReasonMetadataKey)
)

// New Создает ошибку
func New(msgOrErr any) *Throw {
	return newThrow(Unknown, msgOrErr)
}

// NewDeny Создает ошибку прав доступа
func NewDeny(msgOrErr any) *Throw {
	return newThrow(Deny, msgOrErr)
}

// NewDuplicate Создает ошибку ограничения уникальности
func NewDuplicate(msgOrErr any) *Throw {
	return newThrow(Duplicate, msgOrErr)
}

// NewNotFound Создает ошибку не найденной сущности
func NewNotFound(msgOrErr any) *Throw {
	return newThrow(NotFound, msgOrErr)
}

// NewBadRequest Создает ошибку на корректного запроса
func NewBadRequest(msgOrErr any) *Throw {
	return newThrow(BadRequest, msgOrErr)
}

// NewAuth Создает ошибку авторизации
func NewAuth(msgOrErr any) *Throw {
	return newThrow(Auth, msgOrErr)
}

// Is проверяет что ошибка является типом Throw
func Is(err error) bool {
	return err != nil && Cast(err).kind != Unknown
}

// Cast приводит объект к типу Throw
func Cast(err error) *Throw {
	if err == nil {
		return &Throw{}
	}

	var e *Throw
	if errors.As(err, &e) {
		return e
	}

	return &Throw{
		msg: err.Error(),
	}
}

func Trace(err error, name string, with map[string]any) *Throw {
	th := Cast(err)
	th.addTrace(name, with)
	return th
}

func With(err error, key string, val any) *Throw {
	th := Cast(err)
	th.addWith(key, val)
	return th
}

func IsDuplicate(err error) bool {
	return err != nil && Cast(err).kind == Duplicate
}

func IsNotFound(err error) bool {
	return err != nil && Cast(err).kind == NotFound
}

func IsDeny(err error) bool {
	return err != nil && Cast(err).kind == Deny
}

func IsAuth(err error) bool {
	return err != nil && Cast(err).kind == Auth
}

func IsBadRequest(err error) bool {
	return err != nil && Cast(err).kind == BadRequest
}

func validateReasonName(n string) string {
	if n == "" {
		return "empty reason name"
	}

	if len(n) > reasonNameLenMax {
		return "reason name exceeds maximum length"
	}

	if !reasonNameReg.MatchString(n) {
		return "reason name contains prohibited characters [" + regStrReasonName + "]"
	}

	return ""
}

func validateReasonMetadataKey(n string) string {
	if n == "" {
		return "empty metadata key"
	}

	if len(n) > reasonMetadataKeyLenMax {
		return "metadata key exceeds maximum length"
	}

	if !reasonMetadataKeyReg.MatchString(n) {
		return "metadata key contains prohibited characters: allow[" +
			regStrReasonMetadataKey + "], passed:[" + n + "]"
	}

	return ""
}
