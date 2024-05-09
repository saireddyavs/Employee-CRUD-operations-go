// repo/employee_repo.go
package repo

import (
	"errors"
	"sync"

	"assesment/model"
)

type EmployeeRepo interface {
	ListEmployees(page, pageSize int) ([]model.Employee, error)
	CreateEmployee(employee model.Employee) int
	GetEmployeeByID(id int) (model.Employee, error)
	UpdateEmployee(employee model.Employee) error
	DeleteEmployee(id int) error
}

type employeeRepository struct {
	sync.RWMutex
	employees map[int]model.Employee
	idCounter int
}

func NewEmployeeRepo() EmployeeRepo {
	return &employeeRepository{
		employees: make(map[int]model.Employee),
		idCounter: 1,
	}
}

func (er *employeeRepository) ListEmployees(page, pageSize int) ([]model.Employee, error) {
	er.RLock()
	defer er.RUnlock()

	startIndex := (page - 1) * pageSize
	if startIndex >= len(er.employees) {
		return nil, nil
	}

	endIndex := startIndex + pageSize
	if endIndex > len(er.employees) {
		endIndex = len(er.employees)
	}

	return er.sliceMapValues(er.employees, startIndex, endIndex), nil
}

func (er *employeeRepository) sliceMapValues(m map[int]model.Employee, startIndex, endIndex int) []model.Employee {
	employees := make([]model.Employee, 0, endIndex-startIndex)
	for i := startIndex; i < endIndex; i++ {
		employees = append(employees, m[i+1])
	}
	return employees
}

func (er *employeeRepository) CreateEmployee(employee model.Employee) int {
	er.Lock()
	defer er.Unlock()
	employee.ID = er.idCounter
	er.employees[er.idCounter] = employee
	er.idCounter++
	return employee.ID
}

func (er *employeeRepository) GetEmployeeByID(id int) (model.Employee, error) {
	er.RLock()
	defer er.RUnlock()
	emp, ok := er.employees[id]
	if !ok {
		return model.Employee{}, errors.New("employee not found")
	}
	return emp, nil
}

func (er *employeeRepository) UpdateEmployee(employee model.Employee) error {
	er.Lock()
	defer er.Unlock()
	if _, ok := er.employees[employee.ID]; !ok {
		return errors.New("employee not found")
	}
	er.employees[employee.ID] = employee
	return nil
}

func (er *employeeRepository) DeleteEmployee(id int) error {
	er.Lock()
	defer er.Unlock()
	if _, ok := er.employees[id]; !ok {
		return errors.New("employee not found")
	}
	delete(er.employees, id)
	return nil
}
