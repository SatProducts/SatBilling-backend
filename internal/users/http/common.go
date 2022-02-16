package http

type NewUserRequest struct {
	Login string `json:"login"`
	Password string `json:"password"`
	Permissions uint8 `json:"permissions"`
}