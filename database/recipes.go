package database

import (
	"errors"
	"github.com/asidikfauzi/test-recipes-be-golang/config/migrations"
	"github.com/asidikfauzi/test-recipes-be-golang/models"
	"github.com/asidikfauzi/test-recipes-be-golang/repository/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"time"
)

type RecipeDatabase struct {
	db               *gorm.DB
	categoryDatabase domain.CategoryDatabase
}

func NewRecipeDatabase(
	conn *gorm.DB,
	cd domain.CategoryDatabase) domain.RecipeDatabase {
	return &RecipeDatabase{
		db:               conn,
		categoryDatabase: cd,
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

	recToIng, err := d.GetRecipeToIngredientByRecipeId(id)
	if err != nil {
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
		Ingredients:              recToIng,
	}

	return response, nil

}

func (d *RecipeDatabase) GetRecipeToIngredientByRecipeId(id string) (recipeToIngredient []models.GetAllRecipesToIngredientsWithName, err error) {
	var recipesToIngredients []models.RecipesToIngredients

	uuidID, err := uuid.Parse(id)
	if err != nil {
		return recipeToIngredient, err
	}

	if err = d.db.
		Joins("INNER JOIN ingredients i ON recipes_to_ingredients.ingredient_id = i.ingredient_id").
		Where("recipes_to_ingredients.recipe_id = ?", uuidID).
		Where("recipes_to_ingredients.deleted_at IS NULL").
		Select("recipes_to_ingredients.ingredient_id, " +
			"recipes_to_ingredients.rec_to_ing_amount, " +
			"recipes_to_ingredients.ingredient_id, " +
			"i.ingredient_name").
		Find(&recipesToIngredients).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("Recipe not found!")
		}
		return recipeToIngredient, err
	}

	var response []models.GetAllRecipesToIngredientsWithName
	for _, recToIng := range recipesToIngredients {

		response = append(response, models.GetAllRecipesToIngredientsWithName{
			RecToIngAmount: recToIng.RecToIngAmount,
			IngredientID:   recToIng.IngredientID,
			IngredientName: recToIng.IngredientName,
		})
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

func (d *RecipeDatabase) InsertRecipe(recipe models.RecipeRequest, recipeToIngredients []models.RecipesToIngredientsRequest) error {
	var err error

	_, err = d.categoryDatabase.GetCategoryById(recipe.CategoryId.String())
	if err != nil {
		return err
	}

	tx := d.db.Begin()
	if tx.Error != nil {
		log.Fatal(tx.Error)
	}

	resultRecipe := migrations.Recipes{
		RecipeName:               recipe.RecipeName,
		RecipeDescription:        recipe.RecipeDescription,
		RecipeImage:              recipe.RecipeImage,
		RecipePreparationTime:    recipe.RecipePreparationTime,
		RecipeCookingTime:        recipe.RecipeCookingTime,
		RecipePortionSuggestions: recipe.RecipePortionSuggestions,
		RecipeRating:             recipe.RecipeRating,
		CreatedAt:                time.Now(),
		CategoryId:               recipe.CategoryId,
	}
	result := d.db.Create(&resultRecipe)

	if result.Error != nil {
		tx.Rollback()
		err = errors.New(result.Error.Error())
		return err
	}

	recipeID := resultRecipe.RecipeID

	err = d.InsertRecipesToIngredients(recipeID, recipeToIngredients)
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

func (d *RecipeDatabase) InsertRecipesToIngredients(recipeId uuid.UUID, requestRecToInc []models.RecipesToIngredientsRequest) error {
	var (
		recipesToIngredients migrations.RecipesToIngredients
	)

	for _, ingredient := range requestRecToInc {
		recipeToIngID := uuid.New()
		recipesToIngredients.RecToIngID = recipeToIngID
		recipesToIngredients.RecipeID = recipeId
		recipesToIngredients.IngredientID = ingredient.IngredientID
		recipesToIngredients.RecToIngAmount = ingredient.RecToIngAmount
		recipesToIngredients.CreatedAt = time.Now()

		errCreate := d.db.Create(&recipesToIngredients).Error
		if errCreate != nil {
			errCreate = errors.New(errCreate.Error())
			return errCreate
		}
	}

	return nil
}
