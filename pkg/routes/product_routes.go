package routes

import (
	product_admin_handler "ecomm-crud/pkg/handlers/product_admin_handlers"
	product_customer_handler "ecomm-crud/pkg/handlers/product_customer_handlers"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(route *gin.Engine) {

	productAdminGroup := route.Group("/admin/product")
	{
		// productAdminGroup.GET("", product_admin_handler.CreateProduct)
		productAdminGroup.POST("", product_admin_handler.CreateProduct)
	}

	productCustomerGroup := route.Group("/products")
	{
		productCustomerGroup.GET("", product_customer_handler.ListProducts)
	}
}
