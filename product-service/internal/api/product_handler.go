// internal/api/product_handler.go
package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	
	"product-service/internal/models"
	"product-service/internal/service"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// ListProducts godoc
// @Summary      List products
// @Description  Get paginated products with filtering options
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        categoryId   query     int     false  "Filter by category ID"
// @Param        minPrice     query     number  false  "Filter by minimum price"
// @Param        maxPrice     query     number  false  "Filter by maximum price"
// @Param        q            query     string  false  "Search query"
// @Param        inStock      query     bool    false  "Filter by stock availability"
// @Param        sortBy       query     string  false  "Sort field"
// @Param        sortDir      query     string  false  "Sort direction (asc or desc)"
// @Param        page         query     int     false  "Page number"
// @Param        pageSize     query     int     false  "Items per page"
// @Success      200          {object}  models.PaginatedResponse
// @Failure      400          {object}  ErrorResponse
// @Failure      500          {object}  ErrorResponse
// @Router       /products [get]
func (h *ProductHandler) ListProducts(c *gin.Context) {
	var filter models.ProductFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid filter parameters"})
		return
	}
	
	// Set default values if not provided
	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.PageSize <= 0 {
		filter.PageSize = 20
	}
	
	result, err := h.service.ListProducts(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, result)
}

// GetProduct godoc
// @Summary      Get product by ID
// @Description  Get detailed information about a product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  models.Product
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /products/{id} [get]
func (h *ProductHandler) GetProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid product ID"})
		return
	}
	
	product, err := h.service.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, product)
}

// CreateProduct godoc
// @Summary      Create product
// @Description  Add a new product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      models.Product  true  "Product information"
// @Success      201      {object}  models.Product
// @Failure      400      {object}  ErrorResponse
// @Failure      500      {object}  ErrorResponse
// @Router       /products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid product data"})
		return
	}
	
	if err := h.service.CreateProduct(&product); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, product)
}

// UpdateProduct godoc
// @Summary      Update product
// @Description  Update an existing product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id       path      int             true  "Product ID"
// @Param        product  body      models.Product  true  "Product information"
// @Success      200      {object}  models.Product
// @Failure      400      {object}  ErrorResponse
// @Failure      404      {object}  ErrorResponse
// @Failure      500      {object}  ErrorResponse
// @Router       /products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid product ID"})
		return
	}
	
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid product data"})
		return
	}
	
	// Ensure the ID in the path matches the product
	product.ID = uint(id)
	
	if err := h.service.UpdateProduct(&product); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, product)
}

// DeleteProduct godoc
// @Summary      Delete product
// @Description  Delete an existing product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      204  {object}  nil
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid product ID"})
		return
	}
	
	if err := h.service.DeleteProduct(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	
	c.Status(http.StatusNoContent)
}

type StockUpdateRequest struct {
	Quantity int `json:"quantity" binding:"required"`
}

// UpdateStock godoc
// @Summary      Update product stock
// @Description  Update product stock level
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id      path      int                 true  "Product ID"
// @Param        update  body      StockUpdateRequest  true  "Stock update information"
// @Success      200     {object}  models.Product
// @Failure      400     {object}  ErrorResponse
// @Failure      404     {object}  ErrorResponse
// @Failure      500     {object}  ErrorResponse
// @Router       /products/{id}/stock [patch]
func (h *ProductHandler) UpdateStock(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid product ID"})
		return
	}
	
	var req StockUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid stock update data"})
		return
	}
	
	if err := h.service.UpdateStock(uint(id), req.Quantity); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	
	product, err := h.service.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, product)
}

