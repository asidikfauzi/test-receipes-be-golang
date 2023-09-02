package models

import (
	"github.com/google/uuid"
	"time"
)

type RecipesToIngredients struct {
	RecToIngID     uuid.UUID   `gorm:"uuid;default:uuid_generate_v4();primary_key;column:rec_to_ing_id;" json:"rec_to_ing_id"`
	RecToIngAmount float64     `gorm:"type:double precision;" json:"rec_to_ing_amount"`
	CreatedAt      time.Time   `gorm:"default:null"`
	UpdatedAt      *time.Time  `gorm:"default:null"`
	DeleteAt       *time.Time  `gorm:"default:null"`
	RecipeID       uuid.UUID   `gorm:"type:char(36);"`
	IngredientID   uuid.UUID   `gorm:"type:char(36);"`
	Recipes        Recipes     `gorm:"foreignKey:RecipeID;references:recipe_id"`
	Ingredients    Ingredients `gorm:"foreignKey:IngredientID;references:ingredient_id"`
}
