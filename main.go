// main.go
package main

import (
	"assesment/controller"
	repo "assesment/repository"
	"assesment/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	employeeRepo := repo.NewEmployeeRepo()
	employeeService := service.NewEmployeeService(employeeRepo)
	employeeController := controller.NewEmployeeController(employeeService)

	router.GET("/employees", employeeController.ListEmployees)
	router.POST("/employees", employeeController.CreateEmployee)
	router.GET("/employees/:id", employeeController.GetEmployeeByID)
	router.PUT("/employees/:id", employeeController.UpdateEmployee)
	router.DELETE("/employees/:id", employeeController.DeleteEmployee)

	router.Run(":8080")
}
