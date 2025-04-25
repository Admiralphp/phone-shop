// internal/service/category_service.go
package service

import (
	"errors"

	"phone-accessories/internal/models"
	"phone-accessories/internal/repository"
)

type CategoryService interface {
	CreateCategory(category *models.Category) error
	GetCategoryByID(id uint) (*models.Category, error)
	UpdateCategory(category *models.Category) error
	DeleteCategory(id uint) error
	ListCategories() ([]models.Category, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) CreateCategory(category *models.Category) error {
	if category.Name == "" {
		return errors.New("category name is required")
	}
	
	return s.repo.Create(category)
}

func (s *categoryService) GetCategoryByID(id uint) (*models.Category, error) {
	return s.repo.GetByID(id)
}

func (s *categoryService) UpdateCategory(category *models.Category) error {
	if category.Name == "" {
		return errors.New("category name is required")
	}
	
	return s.repo.Update(category)
}

func (s *categoryService) DeleteCategory(id uint) error {
	return s.repo.Delete(id)
}

func (s *categoryService) ListCategories() ([]models.Category, error) {
	return s.repo.List()
}
