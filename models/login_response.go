package models

// User DTO response info
// @Description response information after login
type UserDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Response Login Response info with a struct of user dto
// @Description response information after login and the jwt token
type LoginResponse struct {
	User  UserDto `json:"user"`
	Token string  `json:"token"`
}
