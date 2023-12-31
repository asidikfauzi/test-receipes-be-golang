package database

import (
	"errors"
	"github.com/asidikfauzi/test-recipes-be-golang/models"
	"github.com/asidikfauzi/test-recipes-be-golang/repository/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type CategoryDatabase struct {
	db *gorm.DB
}

func NewCategoryDatabase(conn *gorm.DB) domain.CategoryDatabase {
	return &CategoryDatabase{
		db: conn,
	}
}

func (d *CategoryDatabase) GetCategories(offset, limit int) ([]models.GetAllCategories, int64, error) {
	var (
		categories []models.Categories
		totalCount int64
	)

	if err := d.db.Where("deleted_at IS NULL").Order("category_name ASC").
		Offset(offset).
		Limit(limit).
		Find(&categories).Error; err != nil {
		return nil, totalCount, err
	}

	if err := d.db.Model(&categories).Where("deleted_at IS NULL").Count(&totalCount).Error; err != nil {
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

func (d *CategoryDatabase) GetCategoryById(id string) (category models.GetAllCategories, err error) {
	var categories models.Categories

	uuidID, err := uuid.Parse(id)
	if err != nil {
		return category, err
	}

	if err = d.db.Where("category_id = ?", uuidID).Where("deleted_at IS NULL").First(&categories).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("Category not found!")
		}
		return category, err
	}

	response := models.GetAllCategories{
		CategoryID:   categories.CategoryID,
		CategoryName: categories.CategoryName,
		CreatedAt:    categories.CreatedAt,
	}

	return response, nil

}

func (d *CategoryDatabase) CheckExists(name string) error {
	var categories models.Categories

	if err := d.db.Where("category_name = ?", name).Where("deleted_at IS NULL").First(&categories).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	return errors.New("Category name already exists")
}

func (d *CategoryDatabase) CheckExistsById(id, name string) error {
	var categories models.Categories

	if err := d.db.Where("category_name = ?", name).Where("category_id != ? ", id).Where("deleted_at IS NULL").First(&categories).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	return errors.New("Category name already exists")
}

func (d *CategoryDatabase) InsertCategory(category models.CategoryRequest) error {
	var categories models.Categories

	categories.CategoryName = category.CategoryName
	categories.CreatedAt = time.Now()

	err := d.db.Create(&categories).Error
	if err != nil {
		err = errors.New(err.Error())
		return err
	}

	return nil
}

func (d *CategoryDatabase) UpdateCategory(id string, updatedCategory models.CategoryRequest) error {
	var (
		categories models.Categories
		err        error
	)

	uuidID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	if err = d.db.Where("category_id = ?", uuidID).Where("deleted_at IS NULL").First(&categories).Error; err != nil {
		return err
	}

	now := time.Now()
	categories.CategoryName = updatedCategory.CategoryName
	categories.UpdatedAt = &now
	if err = d.db.Save(&categories).Error; err != nil {
		return err
	}

	return nil
}

func (d *CategoryDatabase) DeleteCategory(id string) error {
	var (
		categories models.Categories
		err        error
	)

	uuidID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	if err = d.db.Where("category_id = ?", uuidID).First(&categories).Error; err != nil {
		return err
	}

	now := time.Now()
	categories.DeletedAt = &now
	if err = d.db.Save(&categories).Error; err != nil {
		return err
	}

	return nil
}
