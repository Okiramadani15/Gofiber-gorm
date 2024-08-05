package request

type UserCreateRequest struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type UserUpdateRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type UserUpdateCredentialRequest struct {
	Email string `json:"email" validate:"required"`
}
