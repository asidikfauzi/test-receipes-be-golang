package models

import (
	"github.com/google/uuid"
	"time"
)

type Categories struct {
	CategoryID   uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primary_key;column:category_id;" json:"category_id"`
	CategoryName string     `json:"category_name"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

type GetAllCategories struct {
	CategoryID   uuid.UUID `json:"category_id"`
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"created_at"`
}

type CategoryRequest struct {
	CategoryName string `json:"category_name" binding:"required"`
}
