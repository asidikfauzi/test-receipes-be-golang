package inject

import (
	"github.com/asidikfauzi/test-recipes-be-golang/config"
	"github.com/asidikfauzi/test-recipes-be-golang/config/seeds"
	"github.com/asidikfauzi/test-recipes-be-golang/repository/domain"
)

func DependencyInjection() domain.Config {
	db, err := config.Open()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	categorySeeder := seeds.NewCategorySeeder(db)
	ingredientSeeder := seeds.NewIngredientSeeder(db)
	recipeSeeder := seeds.NewRecipeSeeder(db)

	return config.NewConfig(
		db,
		categorySeeder,
		ingredientSeeder,
		recipeSeeder,
	)
}
