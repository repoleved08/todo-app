package models

// Login Request model info
// @Description Login required information
// @Description with username and password
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

