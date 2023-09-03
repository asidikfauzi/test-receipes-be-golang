package routes

import (
	"github.com/asidikfauzi/test-recipes-be-golang/config"
	"github.com/asidikfauzi/test-recipes-be-golang/controllers/category"
	"github.com/asidikfauzi/test-recipes-be-golang/controllers/ingredient"
	"github.com/gin-gonic/gin"
)

type IRoutes interface {
	InitRoutes()
}

type RoutesService struct {
	CategoryController   category.CategoryController     `inject:"category_controller"`
	IngredientController ingredient.IngredientController `inject:"ingredient_controller"`
}

func InitPackage() *RoutesService {
	return &RoutesService{
		CategoryController:   &category.MasterCategory{},
		IngredientController: &ingredient.MasterIngredient{},
	}
}

func (r *RoutesService) InitRoutes() {
	g := gin.Default()

	endpoint := g.Group("/v1")
	{
		categories := endpoint.Group("/category")
		{
			categories.GET("/", r.CategoryController.GetAllCategories)
			categories.GET("/:id", r.CategoryController.GetCategoryById)
			categories.POST("/", r.CategoryController.CreateCategory)
			categories.PUT("/:id", r.CategoryController.UpdateCategory)
			categories.DELETE("/:id", r.CategoryController.DeleteCategory)
		}

		ingredients := endpoint.Group("/ingredient")
		{
			ingredients.GET("/", r.IngredientController.GetAllIngredients)
			ingredients.GET("/:id", r.IngredientController.GetIngredientById)
			ingredients.POST("/", r.IngredientController.CreateIngredient)
			ingredients.PUT("/:id", r.IngredientController.UpdateIngredient)
			ingredients.DELETE("/:id", r.IngredientController.DeleteIngredient)
		}
	}

	err := g.Run(":" + config.GetEnv("PORT"))
	if err != nil {
		return
	}

}
