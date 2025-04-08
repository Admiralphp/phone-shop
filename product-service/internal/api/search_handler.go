// internal/api/search_handler.go
package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	
	"product-service/internal/service"
)

type SearchHandler struct {
	service service.SearchService
}

func NewSearchHandler(service service.SearchService) *SearchHandler {
	return &SearchHandler{service: service}
}

// Search godoc
// @Summary      Search products
// @Description  Search products by query
// @Tags         search
// @Accept       json
// @Produce      json
// @Param        q         query     string  true   "Search query"
// @Param        page      query     int     false  "Page number"
// @Param        pageSize  query     int     false  "Items per page"
// @Success      200       {object}  models.PaginatedResponse
// @Failure      400       {object}  ErrorResponse
// @Failure      500       {object}  ErrorResponse
// @Router       /search [get]
func (h *SearchHandler) Search(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Search query is required"})
		return
	}
	
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	if err != nil || pageSize < 1 {
		pageSize = 20
	}
	
	result, err := h.service.Search(query, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, result)
}


