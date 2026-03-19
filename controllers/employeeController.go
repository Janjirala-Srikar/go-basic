package controllers

import (
	"employee-api/config"
	"employee-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEmployees(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, name, email, role FROM employees")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var employees []models.Employee

	for rows.Next() {
		var emp models.Employee
		rows.Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Role)
		employees = append(employees, emp)
	}

	c.JSON(http.StatusOK, employees)
}

func CreateEmployee(c *gin.Context) {
	var emp models.Employee

	if err := c.BindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.DB.Exec("INSERT INTO employees (name, email, role) VALUES (?, ?, ?)",
		emp.Name, emp.Email, emp.Role)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee created"})
}


func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")

	var emp models.Employee
	if err := c.BindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.DB.Exec(
		"UPDATE employees SET name=?, email=?, role=? WHERE id=?",
		emp.Name, emp.Email, emp.Role, id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee updated"})
}

func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")

	_, err := config.DB.Exec("DELETE FROM employees WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
}