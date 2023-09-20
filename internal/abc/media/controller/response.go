package media_controller

import "github.com/semenovem/portal/pkg/it"

type avatarUploadResponse struct {
	AvatarID             uint32 `json:"avatar_id"`
	PreviewContentBase64 string `json:"preview_content_base64"`
}

type loadResponse struct {
	Payload string `json:"payload"`
}

type fileUploadView struct {
	ID          uint32 `json:"id"`
	PreviewLink string `json:"preview_link"` // uri
}

type fileUploadResponse struct {
	File fileUploadView `json:"file"`
}

func newFileUploadResponse(f *it.MediaUploadFile) *fileUploadResponse {
	return &fileUploadResponse{
		File: fileUploadView{
			ID:          f.ID,
			PreviewLink: f.PreviewLink,
		},
	}
}

type boxUploadResponse struct {
	BatchID uint32

	Files []struct {
		ID      uint32
		Type    string
		Preview string // base64
	}
}
