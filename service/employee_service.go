// service/employee_service.go
package service

import (
	"assesment/model"
	repo "assesment/repository"
)

type EmployeeService struct {
	EmployeeRepo repo.EmployeeRepo
}

func NewEmployeeService(employeeRepo repo.EmployeeRepo) *EmployeeService {
	return &EmployeeService{EmployeeRepo: employeeRepo}
}

func (es *EmployeeService) ListEmployees(page, pageSize int) ([]model.Employee, error) {
	return es.EmployeeRepo.ListEmployees(page, pageSize)
}

func (es *EmployeeService) CreateEmployee(employee model.Employee) int {
	return es.EmployeeRepo.CreateEmployee(employee)
}

func (es *EmployeeService) GetEmployeeByID(id int) (model.Employee, error) {
	return es.EmployeeRepo.GetEmployeeByID(id)
}

func (es *EmployeeService) UpdateEmployee(employee model.Employee) error {
	return es.EmployeeRepo.UpdateEmployee(employee)
}

func (es *EmployeeService) DeleteEmployee(id int) error {
	return es.EmployeeRepo.DeleteEmployee(id)
}
