package models

import (
	"github.com/google/uuid"
	"time"
)

type Ingredients struct {
	IngredientID   uuid.UUID  `gorm:"uuid;default:uuid_generate_v4();primary_key;column:ingredient_id;" json:"ingredient_id"`
	IngredientName string     `gorm:"type:varchar(120);not null" json:"ingredient_name"`
	CreatedAt      time.Time  `gorm:"default:null"`
	UpdatedAt      *time.Time `gorm:"default:null"`
	DeleteAt       *time.Time `gorm:"default:null"`
}
