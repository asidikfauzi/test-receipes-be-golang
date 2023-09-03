package recipe

import (
	"github.com/asidikfauzi/test-recipes-be-golang/repository/domain"
	"github.com/gin-gonic/gin"
)

type RecipeController interface {
	GetAllRecipes(c *gin.Context)
	GetIngredientById(c *gin.Context)
}

type MasterRecipe struct {
	RecipeDatabase domain.RecipeDatabase `inject:"recipe_database"`
}
