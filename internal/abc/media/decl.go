package media

import (
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

type ConfigMedia struct {
	AvatarMaxBytes uint32 // Максимальный размер файла аватарки в байтах
	ImageMaxBytes  uint32 // Максимальный размер файла фото в байтах
	VideoMaxBytes  uint32 // Максимальный размер файла видео в байтах
	DocMaxBytes    uint32 // Максимальный размер файла документа в байтах
}

type ObjectType string
type ContentType string

func ObjectByContentType(s string) (ObjectType, error) {
	if o, ok := mediaContentTypes[s]; ok {
		return o, nil
	}

	return "", throw.ErrUnknownContentType
}
