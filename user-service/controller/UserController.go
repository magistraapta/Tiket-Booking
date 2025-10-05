package controller

import (
	"encoding/json"
	"net/http"
	"user-service/model"
	"user-service/service"

	"github.com/gorilla/mux"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := &model.CreateUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.UserResponse{
			Success: false,
			Message: "Invalid JSON format",
			Error:   err.Error(),
		})
		return
	}
}

func (c *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.UserResponse{
			Success: false,
			Message: "User ID is required",
			Error:   "User ID must be provided in the URL path",
		})
		return
	}
	user, err := c.userService.GetUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.UserResponse{
			Success: false,
			Message: "User not found",
			Error:   err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := c.userService.GetAllUsers()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
