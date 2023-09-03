package category

import (
	"github.com/asidikfauzi/test-recipes-be-golang/repository/domain"
	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	GetAllCategories(c *gin.Context)
	GetCategoryById(c *gin.Context)
	CreateCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
}

type MasterCategory struct {
	CategoryDatabase domain.CategoryDatabase `inject:"category_database"`
}
