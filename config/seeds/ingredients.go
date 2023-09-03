package seeds

import (
	"github.com/asidikfauzi/test-recipes-be-golang/config/migrations"
	"github.com/asidikfauzi/test-recipes-be-golang/repository/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type IngredientSeed struct {
	db *gorm.DB
}

func NewIngredientSeeder(conn *gorm.DB) domain.IngredientMigration {
	return &IngredientSeed{
		db: conn,
	}
}

func (g *IngredientSeed) UpIngredientSeeder() {
	ingredientNames := []string{
		"Egg",
		"Tea Leaves",
		"Flour",
		"Rice",
		"Red Onion",
		"Garlic",
		"Strawberry",
		"Bread",
	}

	for _, ingredientName := range ingredientNames {
		var existingIngredient migrations.Ingredients
		if err := g.db.Where("ingredient_name = ?", ingredientName).First(&existingIngredient).Error; err == nil {
			continue
		}

		ingredientID, _ := uuid.NewRandom()
		newIngredient := migrations.Ingredients{
			IngredientID:   ingredientID,
			IngredientName: ingredientName,
			CreatedAt:      time.Now(),
		}

		g.db.Create(&newIngredient)
	}
}
