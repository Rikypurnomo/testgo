package authdto

type LoginResponse struct {
	IsAdmin bool   `json:"is_admin"`
	Token   string `json:"token"`
}
