package dtos

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type LoginResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	ApiKey   string `json:"api_key"`
}
