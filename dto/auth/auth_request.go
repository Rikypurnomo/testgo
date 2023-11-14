package authdto

type AuthRequest struct {
	Name     string `json:"Name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	IsAdmin  bool   `json:"is_admin"`
	Password string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	IsAdmin  bool   `json:"is_admin"`
	Password string `json:"password" validate:"required"`
}
