package service

import (
	"net/http"
)

type EmployeeService interface {
	CreateEmployee(w http.ResponseWriter, r *http.Request) //Create
	GetEmployee(w http.ResponseWriter, r *http.Request)    //Read
	UpdateEmployee(w http.ResponseWriter, r *http.Request) //Update
	DeleteEmployee(w http.ResponseWriter, r *http.Request) //Delete
	ListAllEmployees()
}
