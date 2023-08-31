package it

import "errors"

const (
	MediaObjectPNG  = "png"
	MediaObjectJPEG = "jpeg"
	MediaObjectPDF  = "pdf"
)

var (
	ErrUnknownContentType = errors.New("unknown content type")
)

var (
	mediaContentTypes = map[string]MediaObjectType{
		"image/png":       MediaObjectPNG,
		"image/jpeg":      MediaObjectJPEG,
		"application/pdf": MediaObjectPDF,
	}
)

type MediaObjectType string
type MediaContentType string

type MediaObject struct {
	ID          uint32
	Typ         MediaObjectType
	PreviewLink string
	Note        string
}

func MediaObjectByContentType(s string) (MediaObjectType, error) {
	if o, ok := mediaContentTypes[s]; ok {
		return o, nil
	}

	return "", ErrUnknownContentType
}
