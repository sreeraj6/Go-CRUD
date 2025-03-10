package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ecomm-crud/pkg/routes"
	"ecomm-crud/pkg/config/db"
)

func main() {
	fmt.Println("Server is starting")
	db.CreateDBConnection()
	router := gin.Default()
	routes.RegisterUserRoutes(router)
	router.Run(":8080")
}