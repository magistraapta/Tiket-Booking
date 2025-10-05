package model

type UserModel struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Data    *UserModel `json:"data,omitempty"`
	Error   string     `json:"error,omitempty"`
}
