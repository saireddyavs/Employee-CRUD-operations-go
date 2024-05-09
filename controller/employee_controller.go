package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"assesment/model"
	"assesment/service"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	EmployeeService *service.EmployeeService
}

func NewEmployeeController(employeeService *service.EmployeeService) *EmployeeController {
	return &EmployeeController{EmployeeService: employeeService}
}

func (ec *EmployeeController) ListEmployees(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}

	employees, err := ec.EmployeeService.ListEmployees(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch employees"})
		return
	}
	c.JSON(http.StatusOK, employees)
}

func (ec *EmployeeController) CreateEmployee(c *gin.Context) {
	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := ec.EmployeeService.CreateEmployee(employee)
	employee.ID = id

	c.JSON(http.StatusCreated, employee)
}

func (ec *EmployeeController) GetEmployeeByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	employee, err := ec.EmployeeService.GetEmployeeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func (ec *EmployeeController) UpdateEmployee(c *gin.Context) {
	var updatedEmployee model.Employee
	if err := c.ShouldBindJSON(&updatedEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}
	updatedEmployee.ID = id

	err = ec.EmployeeService.UpdateEmployee(updatedEmployee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update employee:" + fmt.Sprint(err)})
		return
	}

	c.JSON(http.StatusOK, updatedEmployee)
}

func (ec *EmployeeController) DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	err = ec.EmployeeService.DeleteEmployee(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete employee:" + fmt.Sprint(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}
