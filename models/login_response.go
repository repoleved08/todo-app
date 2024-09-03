package models

type UserDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type LoginResponse struct {
	User  UserDto `json:"user"`
	Token string  `json:"token"`
}
