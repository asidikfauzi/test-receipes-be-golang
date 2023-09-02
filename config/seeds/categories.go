package seeds

import (
	"github.com/asidikfauzi/test-recipes-be-golang/models"
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

	var categoryIDs []uuid.UUID

	for i := 0; i < 3; i++ {
		categoryID, _ := uuid.NewRandom()
		categoryIDs = append(categoryIDs, categoryID)
	}

	categories := []models.Categories{
		{
			CategoryID:   categoryIDs[0],
			CategoryName: "The Main Food",
			CreatedAt:    time.Now(),
		},
		{
			CategoryID:   categoryIDs[1],
			CategoryName: "Drink",
			CreatedAt:    time.Now(),
		},
		{
			CategoryID:   categoryIDs[2],
			CategoryName: "Dessert",
			CreatedAt:    time.Now(),
		},
	}

	c.db.Create(categories)
}
