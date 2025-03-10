package routes

import (
	"github.com/gin-gonic/gin"
	"ecomm-crud/pkg/handlers/user-handlers"
)

func RegisterUserRoutes(route *gin.Engine) {
	
	userGroup := route.Group("/users") 
	{
		userGroup.GET("/", userhandlers.UserTestHandler)
	}
}