package models

import (
	"time"

	"gorm.io/gorm"
)

// Todo model info
// @Description Todo information
// @Description with id, title, completed bool, and time of creation
type Todo struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `gorm:"not null" json:"title"`
	Completed bool           `gorm:"default:false" json:"completed"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
