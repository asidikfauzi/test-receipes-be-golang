package seeds

import (
	"github.com/asidikfauzi/test-recipes-be-golang/models"
	"github.com/asidikfauzi/test-recipes-be-golang/repository/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

	var ingredientIDs []uuid.UUID

	for i := 0; i < 8; i++ {
		ingredientID, _ := uuid.NewRandom()
		ingredientIDs = append(ingredientIDs, ingredientID)
	}

	ingredients := []models.Ingredients{
		{
			IngredientID:   ingredientIDs[0],
			IngredientName: "Egg",
		},
		{
			IngredientID:   ingredientIDs[1],
			IngredientName: "Tea Leaves",
		},
		{
			IngredientID:   ingredientIDs[2],
			IngredientName: "Flour",
		},
		{
			IngredientID:   ingredientIDs[3],
			IngredientName: "Rice",
		},
		{
			IngredientID:   ingredientIDs[4],
			IngredientName: "Red Onion",
		},
		{
			IngredientID:   ingredientIDs[5],
			IngredientName: "Garlic",
		},
		{
			IngredientID:   ingredientIDs[6],
			IngredientName: "Strawberry",
		},
		{
			IngredientID:   ingredientIDs[7],
			IngredientName: "Bread",
		},
	}

	g.db.Create(ingredients)
}
