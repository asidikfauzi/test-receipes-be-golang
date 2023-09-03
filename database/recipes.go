package database

import (
	"errors"
	"github.com/asidikfauzi/test-recipes-be-golang/models"
	"github.com/asidikfauzi/test-recipes-be-golang/repository/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"time"
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
			RecipeID:     recipe.RecipeID,
			RecipeName:   recipe.RecipeName,
			RecipeImage:  recipe.RecipeImage,
			RecipeRating: recipe.RecipeRating,
			CategoryId:   recipe.CategoryId,
			CategoryName: recipe.CategoryName,
			CreatedAt:    recipe.CreatedAt,
		})
	}
	return response, totalCount, nil

}

func (d *RecipeDatabase) GetRecipeById(id string) (recipe models.GetRecipesById, err error) {
	var recipes models.Recipes

	uuidID, err := uuid.Parse(id)
	if err != nil {
		return recipe, err
	}

	if err = d.db.
		Joins("INNER JOIN categories c ON recipes.category_id::uuid = c.category_id").
		Where("recipe_id = ?", uuidID).
		Where("recipes.deleted_at IS NULL").
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
		First(&recipes).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("Recipe not found!")
		}
		return recipe, err
	}

	response := models.GetRecipesById{
		RecipeID:                 recipes.RecipeID,
		RecipeName:               recipes.RecipeName,
		RecipeDescription:        recipes.RecipeDescription,
		RecipeImage:              recipes.RecipeImage,
		RecipePreparationTime:    recipes.RecipePreparationTime,
		RecipeCookingTime:        recipes.RecipeCookingTime,
		RecipePortionSuggestions: recipes.RecipePortionSuggestions,
		RecipeRating:             recipes.RecipeRating,
		CategoryId:               recipes.CategoryId,
		CategoryName:             recipes.CategoryName,
		CreatedAt:                recipes.CreatedAt,
	}

	return response, nil

}

func (d *RecipeDatabase) CheckExists(name string) error {
	var recipes models.Recipes

	if err := d.db.Where("recipe_name = ?", name).Where("deleted_at IS NULL").First(&recipes).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	return errors.New("Recipe name already exists")
}

func (d *RecipeDatabase) CheckExistsById(id, name string) error {
	var recipes models.Recipes

	if err := d.db.Where("recipe_name = ?", name).Where("recipe_id != ? ", id).Where("deleted_at IS NULL").First(&recipes).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	return errors.New("Recipe name already exists")
}

func (d *RecipeDatabase) InsertRecipe(recipe models.RecipeRequest, recipeToIngredients models.RecipesToIngredientsRequest) error {

	tx := d.db.Begin()
	if tx.Error != nil {
		log.Fatal(tx.Error)
	}

	var (
		recipes models.Recipes
		err     error
	)

	recipeToIngredients.RecipeID = recipes.RecipeID

	err = d.InsertRecipesToIngredients(recipes.RecipeID, recipe.Ingredients)
	if err != nil {
		tx.Rollback()
		err = errors.New(err.Error())
		return err
	}

	err = d.db.Table("recipes").Create(map[string]interface{}{
		"recipe_name":                recipe.RecipeName,
		"recipe_description":         recipe.RecipeDescription,
		"recipe_image":               recipe.RecipeImage,
		"recipe_preparation_time":    recipe.RecipePreparationTime,
		"recipe_cooking_time":        recipe.RecipeCookingTime,
		"recipe_portion_suggestions": recipe.RecipePortionSuggestions,
		"recipe_rating":              recipe.RecipeRating,
		"created_at":                 time.Now(),
		"category_id":                recipe.CategoryId,
	}).Error
	if err != nil {
		tx.Rollback()
		err = errors.New(err.Error())
		return err
	}

	if err = tx.Commit().Error; err != nil {
		log.Fatal(err)
	}

	return nil
}

func (d *RecipeDatabase) InsertRecipesToIngredients(recipeId uuid.UUID, ingredients []uuid.UUID) error {
	var (
		recipesToIngredients models.RecipesToIngredients
	)

	for _, ingredientID := range ingredients {

		recipesToIngredients.RecipeID = recipeId
		recipesToIngredients.IngredientID = ingredientID
		recipesToIngredients.CreatedAt = time.Now()

		errCreate := d.db.Create(&recipesToIngredients).Error
		if errCreate != nil {
			errCreate = errors.New(errCreate.Error())
			return errCreate
		}
	}

	return nil
}
