package product_admin_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ecomm-crud/pkg/models"
	"ecomm-crud/pkg/services"
)

func CreateProduct(c *gin.Context) {
	var product models.Product
	c.ShouldBindJSON(&product)
	if product.Name == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Name is required for creating product"})
		return
	}
	c.IndentedJSON(http.StatusCreated, services.CreateProduct(&product))
}
