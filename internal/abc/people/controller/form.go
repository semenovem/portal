package people_controller

type UserForm struct {
	UserID uint32 `param:"user_id" validate:"required"`
}
