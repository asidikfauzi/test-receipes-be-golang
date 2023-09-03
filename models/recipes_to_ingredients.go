package models

import (
	"github.com/google/uuid"
	"time"
)

type RecipesToIngredients struct {
	RecToIngID     uuid.UUID  ` gorm:"uuid;default:uuid_generate_v4();primary_key;column:rec_to_ing_id;" json:"rec_to_ing_id"`
	RecToIngAmount float64    `json:"rec_to_ing_amount"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
	RecipeID       uuid.UUID  `json:"recipe_id"`
	IngredientID   uuid.UUID  `json:"ingredient_id"`
	IngredientName string     `json:"ingredient_name"`
}

type GetAllRecipesToIngredients struct {
	RecToIngID     string  `json:"rec_to_ing_id"`
	RecToIngAmount float64 `json:"rec_to_ing_amount"`
	Ingredients    Ingredients
}

type GetAllRecipesToIngredientsWithName struct {
	RecToIngAmount float64   `json:"rec_to_ing_amount"`
	IngredientID   uuid.UUID `json:"ingredient_id"`
	IngredientName string    `json:"ingredient_name"`
}

type RecipesToIngredientsRequest struct {
	RecToIngAmount float64    `json:"rec_to_ing_amount"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	IngredientID   uuid.UUID  `json:"ingredient_id"`
}
