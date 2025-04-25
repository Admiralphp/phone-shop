// internal/api/routes.go
package api

import (
	"github.com/gin-gonic/gin"
	
	"phone-accessories/internal/service"
)

func SetupRoutes(router *gin.Engine, 
	productService service.ProductService, 
	categoryService service.CategoryService,
	searchService service.SearchService) {
	
	// API versioning
	v1 := router.Group("/api/v1")
	
	// Health check endpoint
	v1.GET("/health", HealthCheck)
	
	// Product routes
	products := v1.Group("/products")
	{
		products.GET("", NewProductHandler(productService).ListProducts)
		products.POST("", NewProductHandler(productService).CreateProduct)
		products.GET("/:id", NewProductHandler(productService).GetProduct)
		products.PUT("/:id", NewProductHandler(productService).UpdateProduct)
		products.DELETE("/:id", NewProductHandler(productService).DeleteProduct)
		products.PATCH("/:id/stock", NewProductHandler(productService).UpdateStock)
	}
	
	// Category routes
	categories := v1.Group("/categories")
	{
		categories.GET("", NewCategoryHandler(categoryService).ListCategories)
		categories.POST("", NewCategoryHandler(categoryService).CreateCategory)
		categories.GET("/:id", NewCategoryHandler(categoryService).GetCategory)
		categories.PUT("/:id", NewCategoryHandler(categoryService).UpdateCategory)
		categories.DELETE("/:id", NewCategoryHandler(categoryService).DeleteCategory)
	}
	
	// Search route
	v1.GET("/search", NewSearchHandler(searchService).Search)
}

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
		"service": "product-service",
	})
}
