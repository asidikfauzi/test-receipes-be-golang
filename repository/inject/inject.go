package inject

import (
	"github.com/asidikfauzi/test-recipes-be-golang/config"
	"github.com/asidikfauzi/test-recipes-be-golang/config/seeds"
	"github.com/asidikfauzi/test-recipes-be-golang/database"
	"github.com/asidikfauzi/test-recipes-be-golang/repository/domain"
	"github.com/asidikfauzi/test-recipes-be-golang/routes"
	"github.com/facebookgo/inject"
	"log"
)

type InjectData struct {
	Routes *routes.RoutesService
}

func DependencyInjection(liq InjectData) domain.Config {
	db, err := config.Open()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	categorySeeder := seeds.NewCategorySeeder(db)
	ingredientSeeder := seeds.NewIngredientSeeder(db)
	recipeSeeder := seeds.NewRecipeSeeder(db)

	categoryDatabase := database.NewCategoryDatabase(db)
	ingredientDatabase := database.NewIngredientDatabase(db)
	recipeDatabase := database.NewRecipeDatabase(db, categoryDatabase)

	dependencies := []*inject.Object{
		{Value: categoryDatabase, Name: "category_database"},
		{Value: ingredientDatabase, Name: "ingredient_database"},
		{Value: recipeDatabase, Name: "recipe_database"},
	}

	if liq.Routes != nil {
		dependencies = append(dependencies,
			&inject.Object{Value: liq.Routes, Name: "routes"},
			&inject.Object{Value: liq.Routes.CategoryController, Name: "category_controller"},
			&inject.Object{Value: liq.Routes.IngredientController, Name: "ingredient_controller"},
			&inject.Object{Value: liq.Routes.RecipeController, Name: "recipe_controller"},
		)
	}

	// dependency injection
	var g inject.Graph
	if err := g.Provide(dependencies...); err != nil {
		log.Fatal("Failed Inject Dependencies", err)
	}

	if err := g.Populate(); err != nil {
		log.Fatal("Failed Populate Inject Dependencies", err)
	}

	return config.NewConfig(
		db,
		categorySeeder,
		ingredientSeeder,
		recipeSeeder,
	)
}
