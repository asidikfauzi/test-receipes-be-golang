package domain

import "github.com/asidikfauzi/test-recipes-be-golang/models"

type RecipeDatabase interface {
	GetRecipes(offset, limit int) ([]models.GetAllRecipes, int64, error)
	GetRecipeById(id string) (recipe models.GetRecipesById, err error)
	CheckExists(name string) error
	CheckExistsById(id, name string) error
	InsertRecipe(recipe models.RecipeRequest, recipeToIngredients []models.RecipesToIngredientsRequest) error
	UpdateRecipe(id string, recipe models.RecipeRequest, recipeToIngredients []models.RecipesToIngredientsRequest) error
}
