package main

import (
	"ecomm-crud/pkg/config/db"
	"ecomm-crud/pkg/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Server is starting")
	db.CreateDBConnection()
	router := gin.Default()
	for _, route := range router.Routes() {
		fmt.Println(route.Method, route.Path)
	}

	routes.RegisterUserRoutes(router)
	routes.RegisterProductRoutes(router)
	router.Run(":8080")
}
