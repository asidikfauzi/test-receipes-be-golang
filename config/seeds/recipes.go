package seeds

import (
	"fmt"
	"github.com/asidikfauzi/test-recipes-be-golang/models"
	"github.com/asidikfauzi/test-recipes-be-golang/repository/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type RecipeSeed struct {
	db *gorm.DB
}

func NewRecipeSeeder(conn *gorm.DB) domain.RecipeMigration {
	return &RecipeSeed{
		db: conn,
	}
}

func (r *RecipeSeed) UpRecipeSeeder() {

	var recipeIDs []uuid.UUID

	for i := 0; i < 3; i++ {
		recipeID, _ := uuid.NewRandom()
		recipeIDs = append(recipeIDs, recipeID)
	}

	var mainFoodCategory models.Categories
	if err := r.db.Where("category_name = ?", "The Main Food").First(&mainFoodCategory).Error; err != nil {
		fmt.Println("Category 'The Main Food' not found:", err)
		return
	}

	var drinkCategory models.Categories
	if err := r.db.Where("category_name = ?", "Drink").First(&drinkCategory).Error; err != nil {
		fmt.Println("Category 'Drink' not found:", err)
		return
	}

	var dessertCategory models.Categories
	if err := r.db.Where("category_name = ?", "Dessert").First(&dessertCategory).Error; err != nil {
		fmt.Println("Category 'Dessert' not found:", err)
		return
	}

	recipes := []models.Recipes{
		{
			RecipeID:   recipeIDs[0],
			RecipeName: "Fried Rice",
			RecipeDescription: "Fried rice is a delicious and popular dish from Indonesia. " +
				"It is a dish made of rice that is stir-fried together with various " +
				"spices, meat, vegetables and other ingredients. Fried rice is often served " +
				"with eggs, pickles, and crackers for a more delicious taste.",
			RecipeImage:              "nasigoreng.png",
			RecipePreparationTime:    "20-30 minutes",
			RecipeCookingTime:        "5-7 minutes",
			RecipePortionSuggestions: "2 cups of rice",
			RecipeRating:             "4",
			CreatedAt:                time.Now(),
			CategoryId:               mainFoodCategory.CategoryID,
		},
		{
			RecipeID:   recipeIDs[1],
			RecipeName: "Ice Tea",
			RecipeDescription: "Iced tea is a drink that blends with cultural diversity and is always " +
				"a favorite choice for many people. This drink is made by mixing brewed tea with ice cubes. " +
				"The most commonly used teas are black or green tea, although there are also other " +
				"variations such as herbal teas.",
			RecipeImage:              "esteh.png",
			RecipePreparationTime:    "5-10 minutes",
			RecipeCookingTime:        "5-10 minutes",
			RecipePortionSuggestions: "2 liters of water",
			RecipeRating:             "5",
			CreatedAt:                time.Now(),
			CategoryId:               drinkCategory.CategoryID,
		},
		{
			RecipeID:   recipeIDs[2],
			RecipeName: "Woodpecker Coffee",
			RecipeDescription: "Fried rice is a delicious and popular dish from Indonesia. " +
				"It is a dish made of rice that is stir-fried together with various " +
				"spices, meat, vegetables and other ingredients. Fried rice is often served " +
				"with eggs, pickles, and crackers for a more delicious taste.",
			RecipeImage:              "woodpecker-coffee.png",
			RecipePreparationTime:    "20-30 minutes",
			RecipeCookingTime:        "10-15 minutes",
			RecipePortionSuggestions: "2 eggs and 0.5 kg of flour",
			RecipeRating:             "4.5",
			CreatedAt:                time.Now(),
			CategoryId:               dessertCategory.CategoryID,
		},
	}

	r.db.Create(recipes)
}
