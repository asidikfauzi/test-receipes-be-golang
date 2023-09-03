package routes

import (
	"github.com/asidikfauzi/test-recipes-be-golang/config"
	"github.com/asidikfauzi/test-recipes-be-golang/controllers/category"
	"github.com/gin-gonic/gin"
	"log"
)

type IRoutes interface {
	InitRoutes()
}

type RoutesService struct {
	CategoryController category.CategoryController `inject:"category_controller"`
}

func InitPackage() *RoutesService {
	return &RoutesService{
		CategoryController: &category.MasterCategory{},
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
		}
	}

	log.Println("Server started at http://localhost:9090")
	err := g.Run(":" + config.GetEnv("PORT"))
	if err != nil {
		return
	}

}
