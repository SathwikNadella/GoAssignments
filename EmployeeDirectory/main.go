package main

import (
	"employeeeDirectory/db"
	"employeeeDirectory/repository"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/*
CRUD

*/

func main() {
	db.Connect()

	router := mux.NewRouter()

	repo := repository.NewEmployeeRepo()

	router.HandleFunc("/employees", repo.CreateEmployee).Methods(http.MethodPost)
	router.HandleFunc("/employees/{id}", repo.GetEmployee).Methods(http.MethodGet)
	router.HandleFunc("/employees", repo.UpdateEmployee).Methods(http.MethodPatch)
	router.HandleFunc("/employees/{id}", repo.DeleteEmployee).Methods(http.MethodDelete)

	fmt.Println("Starting Server")
	http.ListenAndServe(":8080", router)
}
