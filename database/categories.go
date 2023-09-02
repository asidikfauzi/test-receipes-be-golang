package database

import (
	"github.com/asidikfauzi/test-recipes-be-golang/models"
	"github.com/asidikfauzi/test-recipes-be-golang/repository/domain"
	"gorm.io/gorm"
)

type CategoryDatabase struct {
	db *gorm.DB
}

func NewRCategoryDatabase(conn *gorm.DB) domain.CategoryDatabase {
	return &CategoryDatabase{
		db: conn,
	}
}

func (d *CategoryDatabase) GetCategories(offset, limit int) ([]models.GetAllCategories, int64, error) {
	var (
		categories []models.Categories
		totalCount int64
	)

	if err := d.db.Order("category_name ASC").
		Offset(offset).
		Limit(limit).
		Find(&categories).Error; err != nil {
		return nil, totalCount, err
	}

	if err := d.db.Model(&categories).Count(&totalCount).Error; err != nil {
		return nil, totalCount, err
	}

	var response []models.GetAllCategories
	for _, category := range categories {
		response = append(response, models.GetAllCategories{
			CategoryID:   category.CategoryID,
			CategoryName: category.CategoryName,
			CreatedAt:    category.CreatedAt,
		})
	}
	return response, totalCount, nil

}
