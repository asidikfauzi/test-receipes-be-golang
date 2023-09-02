package domain

import (
	"github.com/asidikfauzi/test-recipes-be-golang/models"
)

type CategoryDatabase interface {
	GetCategories(offset, limit int) ([]models.GetAllCategories, int64, error)
}
