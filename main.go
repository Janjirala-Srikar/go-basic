package main

import (
	"employee-api/config"
	"employee-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	routes.EmployeeRoutes(r)

	r.Run(":8080")
}