package main

import (
	"log"
	"net/http"

	"github.com/eduufreire/rest-api-users/api"
)

func main() {

	userHandler := api.Init()

	http.HandleFunc("POST /user/", userHandler.CreateUser)
	http.HandleFunc("GET /user/", userHandler.GetAllUsers)
	http.HandleFunc("GET /user/{id}", userHandler.GetUserById)
	http.HandleFunc("DELETE /user/{id}", userHandler.DeleteUser)
	http.HandleFunc("PUT /user/{id}", userHandler.EditUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
