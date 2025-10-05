package repository

import (
	"errors"
	"log"
	"user-service/model"

	"github.com/google/uuid"
)

var userDummy = []*model.UserModel{
	{
		ID:       "1",
		Username: "John Doe",
		Email:    "john.doe@example.com",
		Password: "password",
	},
}

type UserRepository struct{}

func NewUserRepository(userDummy []*model.UserModel) *UserRepository {

	return &UserRepository{}
}

func (r *UserRepository) CreateUser(userRequest *model.CreateUserRequest) (*model.UserResponse, error) {
	user := &model.UserModel{
		Username: userRequest.Username,
		Email:    userRequest.Email,
		Password: userRequest.Password,
		ID:       uuid.New().String(),
	}
	userDummy = append(userDummy, user)
	log.Println("User created: ", user.ID)
	return &model.UserResponse{
		Success: true,
		Message: "User created successfully",
		Data:    user,
	}, nil
}

func (r *UserRepository) GetUserById(id string) (*model.UserResponse, error) {
	for _, user := range userDummy {
		if user.ID == id {
			log.Println("User found: ", user)
			return &model.UserResponse{
				Success: true,
				Message: "User found successfully",
				Data:    user,
			}, nil
		}
	}
	log.Println("User not found")
	return nil, errors.New("user not found")
}

func (r *UserRepository) GetAllUsers() []*model.UserModel {
	return userDummy
}
