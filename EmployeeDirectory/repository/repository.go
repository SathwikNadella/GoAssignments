package repository

import (
	"context"
	"employeeeDirectory/db"
	"employeeeDirectory/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

type EmployeeRepo struct {
	employees map[int]models.Employee
}

func NewEmployeeRepo() *EmployeeRepo {
	return &EmployeeRepo{employees: make(map[int]models.Employee)}
}

var id = 1000

func (r *EmployeeRepo) CreateEmployee(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var employee models.Employee

	if err := json.NewDecoder(req.Body).Decode(&employee); err != nil {
		http.Error(w, "Invalid Request body", http.StatusBadRequest)
		return
	}

	employee.EmployeeID = id
	id += 1

	fmt.Println("Trying to insert:", employee)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.GetCollection("EmployeeDirectoryDB", "Employees")
	_, err := collection.InsertOne(ctx, employee)

	if err != nil {
		http.Error(w, "Failed to save to DB", http.StatusConflict)
	}
	w.WriteHeader(http.StatusCreated)

}

func (r *EmployeeRepo) GetEmployee(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)

	searchid, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Request body", http.StatusBadRequest)
		return
	}

	fmt.Println("Got id: ", searchid)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.GetCollection("EmployeeDirectoryDB", "Employees")
	var employee models.Employee
	err1 := collection.FindOne(ctx, bson.M{"employeeID": searchid}).Decode(&employee)

	if err1 != nil {
		fmt.Println(err1.Error())
		http.Error(w, "Failed to get employee", http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(employee)
}

func (r *EmployeeRepo) UpdateEmployee(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var emp models.Employee

	if err := json.NewDecoder(req.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid Request body", http.StatusBadRequest)
		return
	}

	fmt.Println(emp)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := db.GetCollection("EmployeeDirectoryDB", "Employees")
	var employee models.Employee
	err1 := collection.FindOneAndUpdate(ctx, bson.M{"employeeID": emp.ID()}, bson.M{"$set": emp}).Decode(&employee)

	if err1 != nil {
		fmt.Println(err1.Error())
		http.Error(w, "Failed to update employee", http.StatusUnprocessableEntity)
	}
	w.WriteHeader(http.StatusCreated)
}

func (r *EmployeeRepo) DeleteEmployee(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)

	delid, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Request body", http.StatusBadRequest)
		return
	}

	fmt.Println("Gotta delete: ", delid)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.GetCollection("EmployeeDirectoryDB", "Employees")
	var employee models.Employee
	err1 := collection.FindOneAndDelete(ctx, bson.M{"employeeID": delid}).Decode(&employee)

	if err1 != nil {
		fmt.Println(err1.Error())
		http.Error(w, "Failed to delete employee", http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(employee)
}

func (r *EmployeeRepo) ListAllEmployees() {
	fmt.Println(r.employees)
}
