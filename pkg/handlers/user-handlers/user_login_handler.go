package userhandlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ecomm-crud/pkg/services/user-services"
	"net/http"
)


func UserTestHandler(c *gin.Context) {
	fmt.Println("This is user Handler")

	users := user_service.GetUsers()

	c.IndentedJSON(http.StatusOK, users)
}