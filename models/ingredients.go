package models

import (
	"github.com/google/uuid"
	"time"
)

type Ingredients struct {
	IngredientID   uuid.UUID  `json:"ingredient_id"`
	IngredientName string     `json:"ingredient_name"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
}

type GetAllIngredients struct {
	IngredientID   uuid.UUID `json:"ingredient_id"`
	IngredientName string    `json:"ingredient_name"`
	CreatedAt      time.Time `json:"created_at"`
}

type IngredientRequest struct {
	IngredientName string `json:"ingredient_name" binding:"required"`
}
