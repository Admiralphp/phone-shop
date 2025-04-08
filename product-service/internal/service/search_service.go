// internal/service/search_service.go
package service

import (
	"product-service/internal/models"
	"product-service/internal/repository"
)

type SearchService interface {
	Search(query string, page, pageSize int) (*models.PaginatedResponse, error)
}

type searchService struct {
	repo repository.ProductRepository
}

func NewSearchService(repo repository.ProductRepository) SearchService {
	return &searchService{repo: repo}
}

func (s *searchService) Search(query string, page, pageSize int) (*models.PaginatedResponse, error) {
	return s.repo.Search(query, page, pageSize)
}
