// internal/repository/category_repository.go
package repository

import (
	"errors"

	"gorm.io/gorm"

	"phone-accessories/internal/models"
)

type CategoryRepository interface {
	Create(category *models.Category) error
	GetByID(id uint) (*models.Category, error)
	Update(category *models.Category) error
	Delete(id uint) error
	List() ([]models.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) GetByID(id uint) (*models.Category, error) {
	var category models.Category
	if err := r.db.Preload("Parent").First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) Update(category *models.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(id uint) error {
	// Check if there are products associated with this category
	var count int64
	if err := r.db.Model(&models.Product{}).Where("category_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	
	if count > 0 {
		return errors.New("cannot delete category with associated products")
	}
	
	return r.db.Delete(&models.Category{}, id).Error
}

func (r *categoryRepository) List() ([]models.Category, error) {
	var categories []models.Category
	if err := r.db.Preload("Parent").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
