package ingredient

import (
	"github.com/asidikfauzi/test-recipes-be-golang/repository/domain"
	"github.com/gin-gonic/gin"
)

type IngredientController interface {
	GetAllIngredients(c *gin.Context)
	GetIngredientById(c *gin.Context)
	CreateIngredient(c *gin.Context)
	UpdateIngredient(c *gin.Context)
	DeleteIngredient(c *gin.Context)
}

type MasterIngredient struct {
	IngredientDatabase domain.IngredientDatabase `inject:"ingredient_database"`
}
