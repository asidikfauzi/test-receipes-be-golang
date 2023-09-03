package seeds

import (
	"github.com/asidikfauzi/test-recipes-be-golang/config/migrations"
	"github.com/asidikfauzi/test-recipes-be-golang/repository/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type CategorySeed struct {
	db *gorm.DB
}

func NewCategorySeeder(conn *gorm.DB) domain.CategoryMigration {
	return &CategorySeed{
		db: conn,
	}
}

func (c *CategorySeed) UpCategorySeeder() {
	categoryNames := []string{"The Main Food", "Drink", "Dessert"}

	for _, categoryName := range categoryNames {
		var existingCategory migrations.Categories
		if err := c.db.Where("category_name = ?", categoryName).First(&existingCategory).Error; err == nil {
			continue
		}

		categoryID, _ := uuid.NewRandom()
		newCategory := migrations.Categories{
			CategoryID:   categoryID,
			CategoryName: categoryName,
			CreatedAt:    time.Now(),
		}

		c.db.Create(&newCategory)
	}
}
