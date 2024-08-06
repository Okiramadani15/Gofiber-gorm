package request

type BookCreateRequest struct {
	Title  string `json:"title" validate:"required"`
	Author string `json:"  r" validate:"required,email"`
	Cover  string `json:"cover"`
}
