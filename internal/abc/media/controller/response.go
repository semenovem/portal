package media_controller

type loadResponse struct {
	Payload string `json:"payload"`
}

type fileUploadView struct {
	ID      uint32 `json:"id"`
	Type    string `json:"type"`
	Preview string `json:"preview"` // base64
}

type uploadResponse struct {
	Files []fileUploadView
}

type boxUploadResponse struct {
	BatchID uint32

	Files []struct {
		ID      uint32
		Type    string
		Preview string // base64
	}
}
