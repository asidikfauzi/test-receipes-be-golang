package models

import (
	"github.com/google/uuid"
	"time"
)

type Recipes struct {
	RecipeID                 uuid.UUID  `json:"recipe_id"`
	RecipeName               string     `json:"recipe_name"`
	RecipeDescription        string     `json:"recipe_description"`
	RecipeImage              string     `json:"recipe_image"`
	RecipePreparationTime    string     `json:"recipe_preparation_time"`
	RecipeCookingTime        string     `json:"recipe_cooking_time"`
	RecipePortionSuggestions string     `json:"recipe_portion_suggestions"`
	RecipeRating             string     `json:"recipe_rating"`
	CreatedAt                time.Time  `json:"created_at"`
	UpdatedAt                *time.Time `json:"updated_at"`
	DeletedAt                *time.Time `json:"deleted_at"`
	CategoryId               uuid.UUID  `json:"category_id"`
	CategoryName             string     `json:"category_name"`
}

type GetAllRecipes struct {
	RecipeID                 uuid.UUID `json:"recipe_id"`
	RecipeName               string    `json:"recipe_name"`
	RecipeDescription        string    `json:"recipe_description"`
	RecipeImage              string    `json:"recipe_image"`
	RecipePreparationTime    string    `json:"recipe_preparation_time"`
	RecipeCookingTime        string    `json:"recipe_cooking_time"`
	RecipePortionSuggestions string    `json:"recipe_portion_suggestions"`
	RecipeRating             string    `json:"recipe_rating"`
	CreatedAt                time.Time `json:"created_at"`
	CategoryId               uuid.UUID `json:"category_id"`
	CategoryName             string    `json:"category_name"`
}
