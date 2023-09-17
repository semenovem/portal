package it

type MediaObjectType string

type MediaUploadFile struct {
	ID          uint32
	Typ         MediaObjectType
	PreviewLink string
	Note        string
}
