package dto

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
