package throw

import "errors"

var (
	e = errors.New

	ErrOverNote           = e("more than one note value passed")
	ErrNoFile             = e("file not sent")
	ErrOverFile           = e("more than one file sent")
	ErrFileTooBig         = e("file too big")
	ErrFileEmpty          = e("file empty")
	ErrUnsupportedContent = e("unsupported content type")
)
