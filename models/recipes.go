package models

import (
	"github.com/google/uuid"
	"time"
)

type Recipes struct {
	RecipeID                 uuid.UUID  `gorm:"uuid;default:uuid_generate_v4();primary_key;column:recipe_id;" json:"recipe_id"`
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
	RecipeID     uuid.UUID `json:"recipe_id"`
	RecipeName   string    `json:"recipe_name"`
	RecipeImage  string    `json:"recipe_image"`
	RecipeRating string    `json:"recipe_rating"`
	CreatedAt    time.Time `json:"created_at"`
	CategoryId   uuid.UUID `json:"category_id"`
	CategoryName string    `json:"category_name"`
}

type GetRecipesById struct {
	RecipeID                 uuid.UUID                            `json:"recipe_id"`
	RecipeName               string                               `json:"recipe_name"`
	RecipeDescription        string                               `json:"recipe_description"`
	RecipeImage              string                               `json:"recipe_image"`
	RecipePreparationTime    string                               `json:"recipe_preparation_time"`
	RecipeCookingTime        string                               `json:"recipe_cooking_time"`
	RecipePortionSuggestions string                               `json:"recipe_portion_suggestions"`
	RecipeRating             string                               `json:"recipe_rating"`
	CreatedAt                time.Time                            `json:"created_at"`
	CategoryId               uuid.UUID                            `json:"category_id"`
	CategoryName             string                               `json:"category_name"`
	Ingredients              []GetAllRecipesToIngredientsWithName `json:"ingredients"`
}

type RecipeRequest struct {
	RecipeID                 uuid.UUID
	RecipeName               string                        `json:"recipe_name"  binding:"required"`
	RecipeDescription        string                        `json:"recipe_description"  binding:"required"`
	RecipeImage              string                        `json:"recipe_image"  binding:"required"`
	RecipePreparationTime    string                        `json:"recipe_preparation_time"  binding:"required"`
	RecipeCookingTime        string                        `json:"recipe_cooking_time"  binding:"required"`
	RecipePortionSuggestions string                        `json:"recipe_portion_suggestions"  binding:"required"`
	RecipeRating             string                        `json:"recipe_rating"  binding:"required"`
	CreatedAt                time.Time                     `json:"created_at"  binding:"required"`
	CategoryId               uuid.UUID                     `json:"category_id" validate:"uuid"`
	Ingredients              []RecipesToIngredientsRequest `json:"ingredients" binding:"required"`
}
