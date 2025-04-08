// internal/service/product_service.go
package service

import (
	"errors"

	"product-service/internal/models"
	"product-service/internal/repository"
)

type ProductService interface {
	CreateProduct(product *models.Product) error
	GetProductByID(id uint) (*models.Product, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(id uint) error
	ListProducts(filter models.ProductFilter) (*models.PaginatedResponse, error)
	UpdateStock(id uint, quantity int) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) CreateProduct(product *models.Product) error {
	// Add validation logic here if needed
	if product.Name == "" {
		return errors.New("product name is required")
	}
	if product.Price <= 0 {
		return errors.New("product price must be greater than zero")
	}
	if product.SKU == "" {
		return errors.New("product SKU is required")
	}
	
	return s.repo.Create(product)
}

func (s *productService) GetProductByID(id uint) (*models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *productService) UpdateProduct(product *models.Product) error {
	// Add validation logic here if needed
	if product.Name == "" {
		return errors.New("product name is required")
	}
	if product.Price <= 0 {
		return errors.New("product price must be greater than zero")
	}
	
	return s.repo.Update(product)
}

func (s *productService) DeleteProduct(id uint) error {
	return s.repo.Delete(id)
}

func (s *productService) ListProducts(filter models.ProductFilter) (*models.PaginatedResponse, error) {
	return s.repo.List(filter)
}

func (s *productService) UpdateStock(id uint, quantity int) error {
	return s.repo.UpdateStock(id, quantity)
}

