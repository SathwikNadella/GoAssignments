package repository

import (
	"employeeeDirectory/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type EmployeeRepo struct {
	employees map[int]models.Employee
}

func NewEmployeeRepo() *EmployeeRepo {
	return &EmployeeRepo{employees: make(map[int]models.Employee)}
}

func (r *EmployeeRepo) CreateEmployee(w http.ResponseWriter, req *http.Request) {

	var emp models.Employee

	if err := json.NewDecoder(req.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid Request body", http.StatusBadRequest)
		return
	}

	fmt.Println(emp)

	if _, exists := r.employees[emp.ID()]; exists {
		fmt.Println("Employee Already Exists")
		http.Error(w, "Invalid Request body", http.StatusUnprocessableEntity)
		return
	}

	r.employees[emp.ID()] = emp

	r.ListAllEmployees()

	w.WriteHeader(http.StatusCreated)

}

func (r *EmployeeRepo) GetEmployee(w http.ResponseWriter, req *http.Request) {

	queryValues := req.URL.Query()

	id, err := strconv.Atoi(queryValues.Get("id"))
	if err != nil {
		http.Error(w, "Invalid Request body", http.StatusBadRequest)
		return
	}

	fmt.Println(id)

	if val, exists := r.employees[id]; !exists {
		fmt.Sprintln("No Employee Found for %v", id)
		r.ListAllEmployees()
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No Employee Found"))
		return
	} else {
		r.ListAllEmployees()
		w.WriteHeader(http.StatusOK)
		jsonData, err := json.Marshal(val)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Write([]byte(jsonData))
	}

}

func (r *EmployeeRepo) UpdateEmployee(w http.ResponseWriter, req *http.Request) {

	var emp models.Employee

	if err := json.NewDecoder(req.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid Request body", http.StatusBadRequest)
		return
	}

	fmt.Println(emp)

	if _, exists := r.employees[emp.ID()]; !exists {
		fmt.Println("Employee Does not Exist")
		http.Error(w, "Invalid Request body", http.StatusUnprocessableEntity)
		return
	}

	r.employees[emp.ID()] = emp

	r.ListAllEmployees()

	w.WriteHeader(http.StatusCreated)
}

func (r *EmployeeRepo) DeleteEmployee(id int) error {

	if _, exists := r.employees[id]; !exists {
		fmt.Println("No Employee found to delete")
		return errors.New("Not a new Employee")
	}

	delete(r.employees, id)

	r.ListAllEmployees()
	return nil
}

func (r *EmployeeRepo) ListAllEmployees() {
	fmt.Println(r.employees)
}
