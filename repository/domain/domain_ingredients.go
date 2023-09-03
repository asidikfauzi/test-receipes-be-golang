package domain

import "github.com/asidikfauzi/test-recipes-be-golang/models"

type IngredientDatabase interface {
	GetIngredients(offset, limit int) ([]models.GetAllIngredients, int64, error)
	GetIngredientById(id string) (models.GetAllIngredients, error)
	CheckExists(name string) error
	CheckExistsById(id, name string) error
	InsertIngredient(category models.IngredientRequest) error
	UpdateIngredient(id string, updatedIngredient models.IngredientRequest) error
	DeleteIngredient(id string) error
}
