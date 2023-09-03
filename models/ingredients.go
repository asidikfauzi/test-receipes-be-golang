package models

import (
	"github.com/google/uuid"
	"time"
)

type Ingredients struct {
	IngredientID   uuid.UUID  `gorm:"uuid;default:uuid_generate_v4();primary_key;column:ingredient_id;" json:"ingredient_id"`
	IngredientName string     `gorm:"type:varchar(120);not null" json:"ingredient_name"`
	CreatedAt      time.Time  `gorm:"default:null" json:"created_at"`
	UpdatedAt      *time.Time `gorm:"default:null" json:"updated_at"`
	DeletedAt      *time.Time `gorm:"default:null" json:"deleted_at"`
}
