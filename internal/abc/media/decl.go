package media

import (
	"errors"
	"github.com/semenovem/portal/pkg/throw"
)

const (
	ObjectPNG  ObjectType = "png"
	ObjectJPEG ObjectType = "jpeg"
	ObjectPDF  ObjectType = "pdf"
)

var (
	mediaContentTypes = map[string]ObjectType{
		"image/png":       ObjectPNG,
		"image/jpeg":      ObjectJPEG,
		"application/pdf": ObjectPDF,
	}
)

var (
	ErrContentObjectForbidden = errors.New("content object type forbidden")
)

type ObjectType string
type ContentType string

func ObjectByContentType(s string) (ObjectType, error) {
	if o, ok := mediaContentTypes[s]; ok {
		return o, nil
	}

	return "", throw.ErrUnknownContentType
}
