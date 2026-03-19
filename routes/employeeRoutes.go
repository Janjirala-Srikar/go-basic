package routes

import (
	"employee-api/controllers"

	"github.com/gin-gonic/gin"
)

func EmployeeRoutes(r *gin.Engine) {
	r.GET("/employees", controllers.GetEmployees)
	r.POST("/employees", controllers.CreateEmployee)
	r.PUT("/employees/:id", controllers.UpdateEmployee)
	r.DELETE("/employees/:id", controllers.DeleteEmployee)
}