package domain

type CategoryMigration interface {
	UpCategorySeeder()
}

type IngredientMigration interface {
	UpIngredientSeeder()
}

type RecipeMigration interface {
	UpRecipeSeeder()
}
