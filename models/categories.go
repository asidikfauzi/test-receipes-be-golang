package models

import (
	"github.com/google/uuid"
	"time"
)

type Categories struct {
	CategoryID   uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primary_key;column:category_id;" json:"category_id"`
	CategoryName string     `gorm:"type:varchar(120);not null" json:"category_name"`
	CreatedAt    time.Time  `gorm:"default:null"`
	UpdatedAt    *time.Time `gorm:"default:null"`
	DeleteAt     *time.Time `gorm:"default:null"`
}

type GetAllCategories struct {
	CategoryID   uuid.UUID `json:"category_id"`
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `gorm:"default:null"`
}

type CategoryRequest struct {
	CategoryName string `json:"category_name" binding:"required"`
}
