package domain

import "github.com/asidikfauzi/test-recipes-be-golang/models"

type RecipeDatabase interface {
	GetRecipes(offset, limit int) ([]models.GetAllRecipes, int64, error)
}
