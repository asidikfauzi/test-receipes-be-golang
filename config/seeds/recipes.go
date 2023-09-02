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

	recipeData := []struct {
		Name              string
		Description       string
		Image             string
		PreparationTime   string
		CookingTime       string
		PortionSuggestion string
		Rating            string
		CreatedAt         time.Time
		CategoryId        uuid.UUID
	}{
		{
			Name: "Fried Rice",
			Description: "Fried rice is a delicious and popular dish from Indonesia. " +
				"It is a dish made of rice that is stir-fried together with various " +
				"spices, meat, vegetables and other ingredients. Fried rice is often served " +
				"with eggs, pickles, and crackers for a more delicious taste.",
			Image:             "nasigoreng.png",
			PreparationTime:   "20-30 minutes",
			CookingTime:       "5-7 minutes",
			PortionSuggestion: "2 cups of rice",
			Rating:            "4",
			CreatedAt:         time.Now(),
			CategoryId:        mainFoodCategory.CategoryID,
		},
		{
			Name: "Ice Tea",
			Description: "Iced tea is a drink that blends with cultural diversity and is always " +
				"a favorite choice for many people. This drink is made by mixing brewed tea with ice cubes. " +
				"The most commonly used teas are black or green tea, although there are also other " +
				"variations such as herbal teas.",
			Image:             "esteh.png",
			PreparationTime:   "5-10 minutes",
			CookingTime:       "5-10 minutes",
			PortionSuggestion: "2 liters of water",
			Rating:            "5",
			CreatedAt:         time.Now(),
			CategoryId:        drinkCategory.CategoryID,
		},
		{
			Name: "Woodpecker Coffee",
			Description: "Fried rice is a delicious and popular dish from Indonesia. " +
				"It is a dish made of rice that is stir-fried together with various " +
				"spices, meat, vegetables and other ingredients. Fried rice is often served " +
				"with eggs, pickles, and crackers for a more delicious taste.",
			Image:             "woodpecker-coffee.png",
			PreparationTime:   "20-30 minutes",
			CookingTime:       "10-15 minutes",
			PortionSuggestion: "2 eggs and 0.5 kg of flour",
			Rating:            "4.5",
			CreatedAt:         time.Now(),
			CategoryId:        dessertCategory.CategoryID,
		},
	}

	for _, data := range recipeData {
		var existingRecipe models.Recipes
		if err := r.db.Where("recipe_name = ?", data.Name).First(&existingRecipe).Error; err == nil {
			continue
		}

		recipeID, _ := uuid.NewRandom()
		newRecipe := models.Recipes{
			RecipeID:                 recipeID,
			RecipeName:               data.Name,
			RecipeDescription:        data.Description,
			RecipeImage:              data.Image,
			RecipePreparationTime:    data.PreparationTime,
			RecipeCookingTime:        data.CookingTime,
			RecipePortionSuggestions: data.PortionSuggestion,
			RecipeRating:             data.Rating,
			CreatedAt:                data.CreatedAt,
			CategoryId:               data.CategoryId,
		}

		r.db.Create(&newRecipe)
	}

}
