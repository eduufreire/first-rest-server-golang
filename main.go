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
	log.Fatal(http.ListenAndServe(":8080", nil))
}
