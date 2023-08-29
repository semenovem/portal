package media_controller

import "github.com/semenovem/portal/pkg/it"

type loadResponse struct {
	Payload string `json:"payload"`
}

type fileUploadView struct {
	ID          uint32 `json:"id"`
	Note        string `json:"note"`
	PreviewLink string `json:"preview_link"` // uri
}

type fileUploadResponse struct {
	File fileUploadView `json:"file"`
}

func newFileUploadResponse(f *it.MediaFile) *fileUploadResponse {
	return &fileUploadResponse{
		File: fileUploadView{
			ID:          f.ID,
			Note:        f.Note,
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
