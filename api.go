package main

import (
	"fmt"
	"log"
	"net/http"
	Auth "restapi/package/Auth"
	Jobs "restapi/package/Jobs"
	Users "restapi/package/Users"

	Companies "restapi/package/Companies"

	"github.com/gorilla/mux"
)

func main() {
	myRouter := mux.NewRouter()
	//AUTH
	myRouter.HandleFunc("/login", Auth.GetLogin).Methods("POST")
	myRouter.HandleFunc("/register", Auth.Register).Methods("POST")
	myRouter.HandleFunc("/logout", Auth.Logout).Methods("GET")

	// USERS (ADMIN)
	myRouter.HandleFunc("/users", Users.GetUsers).Methods("GET")
	myRouter.HandleFunc("/users/{id}", Users.GetUser).Methods("GET")
	myRouter.HandleFunc("/users/{id}", Users.DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/users/{id}", Users.UpdateUser).Methods("PUT")

	//JOBS
	myRouter.HandleFunc("/jobs", Jobs.GetJobs).Methods("GET")
	myRouter.HandleFunc("/jobs", Jobs.PostJob).Methods("POST")
	myRouter.HandleFunc("/jobs/{id}", Jobs.GetJob).Methods("GET")
	myRouter.HandleFunc("/jobs/{id}", Jobs.ApplyJob).Methods("POST")
	myRouter.HandleFunc("/jobs/{id}", Jobs.UpdateJob).Methods("PUT")
	myRouter.HandleFunc("/jobs/{id}", Jobs.DeleteJob).Methods("DELETE")

	//Company
	myRouter.HandleFunc("/company", Companies.GetCompanies).Methods("GET")
	myRouter.HandleFunc("/company/{id}", Companies.GetCompany).Methods("GET")
	myRouter.HandleFunc("/company", Companies.CreateCompany).Methods("POST")
	myRouter.HandleFunc("/company/{id}", Companies.DeleteCompany).Methods("DELETE")

	fmt.Println("PORT RUNNING : 8000")
	log.Fatalln(http.ListenAndServe(":8000", myRouter))
}
