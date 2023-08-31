package throw

// CorrectErr ошибки в результате проверок и валидации
type CorrectErr interface {
	Error() string
}

type correctErr struct {
	msg string
}

func (e correctErr) Error() string {
	return e.msg
}

func NewCorrectErr(msg string) error {
	return &correctErr{msg: msg}
}

func IsCorrectErr(err error) bool {
	_, ok := err.(*correctErr)
	return ok
}
