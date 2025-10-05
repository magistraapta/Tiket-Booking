package main

import (
	"log"
	"net/http"
	"user-service/controller"
	"user-service/repository"
	"user-service/service"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	userRepository := repository.NewUserRepository(nil)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	router.HandleFunc("/users", userController.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userController.GetUserById).Methods("GET")
	router.HandleFunc("/users", userController.GetAllUsers).Methods("GET")

	log.Println("User service is running on port 8083")
	log.Fatal(http.ListenAndServe(":8083", router))
}
