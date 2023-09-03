package models

import (
	"github.com/google/uuid"
	"time"
)

type RecipesToIngredients struct {
	RecToIngID     uuid.UUID  `json:"rec_to_ing_id"`
	RecToIngAmount float64    `json:"rec_to_ing_amount"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
	RecipeID       uuid.UUID  `json:"recipe_id"`
	IngredientID   uuid.UUID  `json:"ingredient_id"`
}

type GetAllRecipesToIngredients struct {
	RecToIngID     string  `json:"rec_to_ing_id"`
	RecToIngAmount float64 `json:"rec_to_ing_amount"`
	Ingredients    Ingredients
}

type RecipesToIngredientsRequest struct {
	RecToIngID     uuid.UUID   `json:"rec_to_ing_id"`
	RecToIngAmount float64     `json:"rec_to_ing_amount"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      *time.Time  `json:"updated_at"`
	RecipeID       uuid.UUID   `json:"recipe_id"`
	IngredientID   []uuid.UUID `json:"ingredient_id"`
}
