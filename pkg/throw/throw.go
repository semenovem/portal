package throw

import (
	"fmt"
	"github.com/semenovem/portal/pkg/throw/throwtrace"
	"runtime"
	"strings"
)

type Throw struct {
	kind        Kind           // Тип лога
	msg         string         // Текст ошибки, как есть в Error()
	desc        string         // Текстовый код, по которому можно получить перевод для локализации
	descArgs    []any          // Аргументы для подстановки в шаблон
	with        map[string]any // Данные ключ-значение для логгера
	tracePoints []*throwtrace.Point
	reasons     []*throwtrace.Reason // Причины ошибки
}

func (t *Throw) Named() string {
	if t == nil {
		return ""
	}
	return t.msg
}

func (t *Throw) Error() string {
	if t == nil {
		return ""
	}
	return t.msg
}

func (t *Throw) Kind() Kind {
	if t == nil {
		return Unknown
	}
	return t.kind
}

// Описание ошибки с аргументами для шаблонизации
func (t *Throw) getDesc() (string, []any) {
	if t == nil {
		return "", nil
	}
	return t.desc, t.descArgs
}

func (t *Throw) SetDesc(desc any, descArgs ...any) *Throw {
	if t == nil {
		return newEmptyThrow().SetDesc(desc, descArgs...)
	}
	t.desc = fmt.Sprintf("%v", desc)
	t.descArgs = descArgs
	return t
}

// SetMsg установить сообщение ошибки
func (t *Throw) SetMsg(msgOrErr any) *Throw {
	if t == nil {
		return newEmptyThrow().SetMsg(msgOrErr)
	}
	t.msg = fmt.Sprintf("%s", msgOrErr)
	return t
}

// With установить сообщение ошибки
func (t *Throw) With(key string, val any) *Throw {
	if t == nil {
		return newEmptyThrow().With(key, val)
	}
	t.addWith(key, val)

	return t
}

func (t *Throw) traceData() ([]*throwtrace.Point, map[string]any) {
	if t == nil {
		return nil, nil
	}
	return t.tracePoints, t.with
}

func (t *Throw) addWith(key string, val any) {
	if len(t.with) == 0 {
		t.with = make(map[string]any)
	}
	t.with[key] = val
}

// AddTrace добавляет точку трейса
// name - имя метода и место, из которого добавляется трейс
// with - данные для логгера, аналог With в zap
func (t *Throw) AddTrace(name string, with map[string]any) *Throw {
	if t == nil {
		th := newEmptyThrow()
		th.addTrace(name, with)
		return th
	}
	t.addTrace(name, with)
	return t
}

func (t *Throw) addTrace(name string, with map[string]any) {
	_, path, line, _ := runtime.Caller(2)
	ps := strings.Split(path, "/")

	if len(ps) > 1 {
		path = strings.Join(ps[len(ps)-2:], "/")
	}

	t.tracePoints = append(t.tracePoints, &throwtrace.Point{
		Name:     name,
		LineCode: fmt.Sprintf("%s:%d", path, line),
	})

	t.unionWith(with)
}

// AddReason добавляет причину проблемы и метаданные к ней.
// reason - UPPER_SNAKE_CASE `[A-Z][A-Z0-9_]+[A-Z0-9]` length <= 63 длины
// metadata - ключ мапы должен быть `[a-zA-Z0-9-_]` length <= 64
func (t *Throw) AddReason(reason string, metadata map[string]string) *Throw {
	if t == nil {
		return newEmptyThrow().AddReason(reason, metadata)
	}

	if s := validateReasonName(reason); s != "" {
		panic(s)
	}

	for k := range metadata {
		if s := validateReasonMetadataKey(k); s != "" {
			panic(s)
		}
	}

	if t == nil {
		return nil
	}
	t.reasons = append(t.reasons, &throwtrace.Reason{Reason: reason, Metadata: metadata})

	return t
}

func (t *Throw) getReasons() []*throwtrace.Reason {
	return t.reasons
}

func (t *Throw) unionWith(m map[string]any) {
	if len(m) == 0 {
		return
	}
	if len(t.with) == 0 {
		t.with = make(map[string]any)
	}

	for k, v := range m {
		t.with[k] = v
	}
}

// NewAuth Создает ошибку авторизации
func newThrow(k Kind, msg any) *Throw {
	var th = &Throw{}

	switch typ := msg.(type) {
	case error:
		if Is(typ) {
			th = Cast(typ)
		}
		th.msg = typ.Error()
	case nil:

	default:
		th.msg = fmt.Sprintf("%v", typ)
	}

	th.kind = k

	return th
}

func newEmptyThrow() *Throw {
	return &Throw{}
}
