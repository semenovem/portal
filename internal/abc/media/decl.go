package media

import "errors"

const (
	MediaObjectPNG  MediaObjectType = "png"
	MediaObjectJPEG MediaObjectType = "jpeg"
	MediaObjectPDF  MediaObjectType = "pdf"
)

var (
	ErrUnknownContentType = errors.New("unknown content type")

	mediaContentTypes = map[string]MediaObjectType{
		"image/png":       MediaObjectPNG,
		"image/jpeg":      MediaObjectJPEG,
		"application/pdf": MediaObjectPDF,
	}
)

type ConfigMedia struct {
	AvatarMaxBytes uint32 // Максимальный размер файла аватарки в байтах
	ImageMaxBytes  uint32 // Максимальный размер файла фото в байтах
	VideoMaxBytes  uint32 // Максимальный размер файла видео в байтах
	DocMaxBytes    uint32 // Максимальный размер файла документа в байтах
}

type MediaObjectType string
type MediaContentType string

func MediaObjectByContentType(s string) (MediaObjectType, error) {
	if o, ok := mediaContentTypes[s]; ok {
		return o, nil
	}

	return "", ErrUnknownContentType
}
