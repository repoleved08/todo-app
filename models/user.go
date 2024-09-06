package models

import "gorm.io/gorm"

// Account model info
// @Description User account information
// @Description with username, email and password
type User struct {
	gorm.Model `swaggerignore:"true"`
	Username   string `gorm:"unique"`
	Email      string `gorm:"unique"`
	Password   string
}
