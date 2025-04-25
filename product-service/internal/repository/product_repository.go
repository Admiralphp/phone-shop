// internal/repository/product_repository.go
package repository

import (
	"errors"
	"math"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"phone-accessories/internal/models"
)

type ProductRepository interface {
	Create(product *models.Product) error
	GetByID(id uint) (*models.Product, error)
	Update(product *models.Product) error
	Delete(id uint) error
	List(filter models.ProductFilter) (*models.PaginatedResponse, error)
	Search(query string, page, pageSize int) (*models.PaginatedResponse, error)
	UpdateStock(id uint, quantity int) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) GetByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := r.db.Preload("Category").First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}

func (r *productRepository) List(filter models.ProductFilter) (*models.PaginatedResponse, error) {
	var products []models.Product
	var totalItems int64

	query := r.db.Model(&models.Product{})

	// Apply filters
	if filter.CategoryID != nil {
		query = query.Where("category_id = ?", *filter.CategoryID)
	}
	if filter.MinPrice != nil {
		query = query.Where("price >= ?", *filter.MinPrice)
	}
	if filter.MaxPrice != nil {
		query = query.Where("price <= ?", *filter.MaxPrice)
	}
	if filter.InStock != nil && *filter.InStock {
		query = query.Where("stock_level > 0")
	}
	if filter.SearchQuery != "" {
		search := "%" + filter.SearchQuery + "%"
		query = query.Where("name ILIKE ? OR description ILIKE ?", search, search)
	}

	// Count total items
	if err := query.Count(&totalItems).Error; err != nil {
		return nil, err
	}

	// Apply sorting
	if filter.SortBy != "" {
		direction := "ASC"
		if strings.ToUpper(filter.SortDirection) == "DESC" {
			direction = "DESC"
		}
		query = query.Order(clause.OrderByColumn{Column: clause.Column{Name: filter.SortBy}, Desc: direction == "DESC"})
	} else {
		query = query.Order("id ASC")
	}

	// Apply pagination
	if filter.Page < 1 {
		filter.Page = 1
	}
	if filter.PageSize < 1 {
		filter.PageSize = 20
	}
	offset := (filter.Page - 1) * filter.PageSize

	if err := query.Preload("Category").Offset(offset).Limit(filter.PageSize).Find(&products).Error; err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(filter.PageSize)))
	return &models.PaginatedResponse{
		Items:      products,
		Page:       filter.Page,
		PageSize:   filter.PageSize,
		TotalItems: totalItems,
		TotalPages: totalPages,
	}, nil
}

func (r *productRepository) Search(query string, page, pageSize int) (*models.PaginatedResponse, error) {
	var products []models.Product
	var totalItems int64

	searchQuery := "%" + query + "%"

	db := r.db.Model(&models.Product{}).
		Where("name ILIKE ? OR description ILIKE ? OR sku ILIKE ?", searchQuery, searchQuery, searchQuery)

	if err := db.Count(&totalItems).Error; err != nil {
		return nil, err
	}

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}

	offset := (page - 1) * pageSize

	if err := db.Preload("Category").Offset(offset).Limit(pageSize).Find(&products).Error; err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))
	return &models.PaginatedResponse{
		Items:      products,
		Page:       page,
		PageSize:   pageSize,
		TotalItems: totalItems,
		TotalPages: totalPages,
	}, nil
}

func (r *productRepository) UpdateStock(id uint, quantity int) error {
	return r.db.Model(&models.Product{}).Where("id = ?", id).
		Update("stock_level", gorm.Expr("stock_level + ?", quantity)).Error
}
