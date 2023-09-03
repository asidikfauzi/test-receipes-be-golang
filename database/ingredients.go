package database

import (
	"errors"
	"github.com/asidikfauzi/test-recipes-be-golang/models"
	"github.com/asidikfauzi/test-recipes-be-golang/repository/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type IngredientDatabase struct {
	db *gorm.DB
}

func NewIngredientDatabase(conn *gorm.DB) domain.IngredientDatabase {
	return &IngredientDatabase{
		db: conn,
	}
}

func (d *IngredientDatabase) GetIngredients(offset, limit int) ([]models.GetAllIngredients, int64, error) {
	var (
		ingredients []models.Ingredients
		totalCount  int64
	)

	if err := d.db.Where("deleted_at IS NULL").Order("ingredient_name ASC").
		Offset(offset).
		Limit(limit).
		Find(&ingredients).Error; err != nil {
		return nil, totalCount, err
	}

	if err := d.db.Model(&ingredients).Where("deleted_at IS NULL").Count(&totalCount).Error; err != nil {
		return nil, totalCount, err
	}

	var response []models.GetAllIngredients
	for _, ingredient := range ingredients {
		response = append(response, models.GetAllIngredients{
			IngredientID:   ingredient.IngredientID,
			IngredientName: ingredient.IngredientName,
			CreatedAt:      ingredient.CreatedAt,
		})
	}
	return response, totalCount, nil

}

func (d *IngredientDatabase) GetIngredientById(id string) (ingredient models.GetAllIngredients, err error) {
	var ingredients models.Ingredients

	uuidID, err := uuid.Parse(id)
	if err != nil {
		return ingredient, err
	}

	if err = d.db.Where("ingredient_id = ?", uuidID).Where("deleted_at IS NULL").First(&ingredients).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("Ingredient not found!")
		}
		return ingredient, err
	}

	response := models.GetAllIngredients{
		IngredientID:   ingredients.IngredientID,
		IngredientName: ingredients.IngredientName,
		CreatedAt:      ingredients.CreatedAt,
	}

	return response, nil

}

func (d *IngredientDatabase) CheckExists(name string) error {
	var ingredients models.Ingredients

	if err := d.db.Where("ingredient_name = ?", name).Where("deleted_at IS NULL").First(&ingredients).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	return errors.New("Ingredient name already exists")
}

func (d *IngredientDatabase) CheckExistsById(id, name string) error {
	var ingredients models.Ingredients

	if err := d.db.Where("ingredient_name = ?", name).Where("ingredient_id != ? ", id).Where("deleted_at IS NULL").First(&ingredients).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	return errors.New("Ingredient name already exists")
}

func (d *IngredientDatabase) InsertIngredient(ingredient models.IngredientRequest) error {
	var ingredients models.Ingredients

	ingredients.IngredientName = ingredient.IngredientName
	ingredients.CreatedAt = time.Now()

	err := d.db.Create(&ingredients).Error
	if err != nil {
		err = errors.New(err.Error())
		return err
	}

	return nil
}

func (d *IngredientDatabase) UpdateIngredient(id string, updatedIngredient models.IngredientRequest) error {
	var (
		ingredients models.Ingredients
		err         error
	)

	uuidID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	if err = d.db.Where("ingredient_id = ?", uuidID).Where("deleted_at IS NULL").First(&ingredients).Error; err != nil {
		return err
	}

	now := time.Now()
	ingredients.IngredientName = updatedIngredient.IngredientName
	ingredients.UpdatedAt = &now
	if err = d.db.Save(&ingredients).Error; err != nil {
		return err
	}

	return nil
}

func (d *IngredientDatabase) DeleteIngredient(id string) error {
	var (
		ingredients models.Ingredients
		err         error
	)

	uuidID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	if err = d.db.Where("ingredient_id = ?", uuidID).First(&ingredients).Error; err != nil {
		return err
	}

	now := time.Now()
	ingredients.DeletedAt = &now
	if err = d.db.Save(&ingredients).Error; err != nil {
		return err
	}

	return nil
}
