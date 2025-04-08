// internal/models/product.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"size:255;not null"`
	Description string         `json:"description" gorm:"type:text"`
	Price       float64        `json:"price" gorm:"not null"`
	SKU         string         `json:"sku" gorm:"size:50;uniqueIndex;not null"`
	StockLevel  int            `json:"stockLevel" gorm:"not null;default:0"`
	ImageURL    string         `json:"imageUrl" gorm:"size:255"`
	CategoryID  uint           `json:"categoryId"`
	Category    Category       `json:"category" gorm:"foreignKey:CategoryID"`
	Attributes  JSON           `json:"attributes" gorm:"type:jsonb"`
	IsActive    bool           `json:"isActive" gorm:"default:true"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// ProductFilter represents the filter options for products
type ProductFilter struct {
	CategoryID    *uint    `form:"categoryId"`
	MinPrice      *float64 `form:"minPrice"`
	MaxPrice      *float64 `form:"maxPrice"`
	SearchQuery   string   `form:"q"`
	InStock       *bool    `form:"inStock"`
	SortBy        string   `form:"sortBy"`
	SortDirection string   `form:"sortDir"`
	Page          int      `form:"page,default=1"`
	PageSize      int      `form:"pageSize,default=20"`
}

// JSON is a custom type for handling JSON in GORM
type JSON map[string]interface{}

// Category represents a product category
type Category struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"size:100;not null;uniqueIndex"`
	Description string         `json:"description" gorm:"type:text"`
	ParentID    *uint          `json:"parentId"`
	Parent      *Category      `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	ImageURL    string         `json:"imageUrl" gorm:"size:255"`
	IsActive    bool           `json:"isActive" gorm:"default:true"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// PaginatedResponse represents a paginated response
type PaginatedResponse struct {
	Items      interface{} `json:"items"`
	Page       int         `json:"page"`
	PageSize   int         `json:"pageSize"`
	TotalItems int64       `json:"totalItems"`
	TotalPages int         `json:"totalPages"`
}
