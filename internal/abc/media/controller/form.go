package media_controller

type storePathForm struct {
	StorePath string `param:"store_path" validate:"required"`
}

type storeForm struct {
	StorePath string `param:"store_path" validate:"required"`
	Payload   string `json:"payload"  validate:"required"`
}
