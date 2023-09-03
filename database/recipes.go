package database

import (
	"github.com/asidikfauzi/test-recipes-be-golang/models"
	"github.com/asidikfauzi/test-recipes-be-golang/repository/domain"
	"gorm.io/gorm"
)

type RecipeDatabase struct {
	db *gorm.DB
}

func NewRecipeDatabase(conn *gorm.DB) domain.RecipeDatabase {
	return &RecipeDatabase{
		db: conn,
	}
}

func (d *RecipeDatabase) GetRecipes(offset, limit int) ([]models.GetAllRecipes, int64, error) {
	var (
		recipes    []models.Recipes
		totalCount int64
	)

	if err := d.db.
		Joins("INNER JOIN categories c ON recipes.category_id::uuid = c.category_id").
		Where("recipes.deleted_at IS NULL").
		Order("recipes.recipe_name ASC").
		Select("recipes.recipe_id, " +
			"recipes.recipe_name, " +
			"recipes.recipe_description, " +
			"recipes.recipe_image, " +
			"recipes.recipe_preparation_time, " +
			"recipes.recipe_cooking_time, " +
			"recipes.recipe_portion_suggestions, " +
			"recipes.recipe_rating, " +
			"recipes.created_at, " +
			"recipes.category_id, " +
			"c.category_name").
		Offset(offset).
		Limit(limit).
		Find(&recipes).Error; err != nil {
		return nil, totalCount, err
	}
	if err := d.db.Model(&recipes).Where("recipes.deleted_at IS NULL").Count(&totalCount).Error; err != nil {
		return nil, totalCount, err
	}

	var response []models.GetAllRecipes
	for _, recipe := range recipes {

		response = append(response, models.GetAllRecipes{
			RecipeID:                 recipe.RecipeID,
			RecipeName:               recipe.RecipeName,
			RecipeDescription:        recipe.RecipeDescription,
			RecipeImage:              recipe.RecipeImage,
			RecipePreparationTime:    recipe.RecipePreparationTime,
			RecipeCookingTime:        recipe.RecipeCookingTime,
			RecipePortionSuggestions: recipe.RecipePortionSuggestions,
			RecipeRating:             recipe.RecipeRating,
			CategoryId:               recipe.CategoryId,
			CategoryName:             recipe.CategoryName,
			CreatedAt:                recipe.CreatedAt,
		})
	}
	return response, totalCount, nil

}
