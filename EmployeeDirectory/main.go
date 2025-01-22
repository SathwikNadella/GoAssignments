package main

import (
	"employeeeDirectory/repository"
	"employeeeDirectory/service"
	"fmt"
	"net/http"
)

/*
CRUD

*/

func main() {

	repo := repository.NewEmployeeRepo()

	Execute(repo)

}

func Execute(repo service.EmployeeService) {

	http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request Method: %s", r.Method)
		switch r.Method {

		case http.MethodPost: //Create
			{

				repo.CreateEmployee(w, r)

			}
		case http.MethodPatch: //Update
			{
				repo.UpdateEmployee(w, r)
			}
		case http.MethodGet: //Read
			{
				repo.GetEmployee(w, r)
			}
		case http.MethodDelete: //Delete
			{
				repo.DeleteEmployee(w, r)
			}
		default:
			{
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}

		}
	})

	fmt.Println("Starting Server")

	http.ListenAndServe(":8080", nil)

}
