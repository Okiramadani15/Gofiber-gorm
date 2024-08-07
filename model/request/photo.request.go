package request

type PhotoCreateRequest struct {
	Categoryid uint `json:"category_id" form:"category_id" validate:"required"`
}
