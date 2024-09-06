package models

import "gorm.io/gorm"

// Product model info
// @Description product information
// @Description with name, description, price, image_url and user_id
type Product struct {
	gorm.Model  `swaggerignore:"true"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url"`
	UserID      uint    `json:"user_id"`
}
