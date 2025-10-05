package service

import (
	"user-service/jwt"
	"user-service/model"
	"user-service/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) CreateUser(user *model.CreateUserRequest) (*model.UserResponse, error) {
	return s.userRepository.CreateUser(user)
}

func (s *UserService) GetUserById(id string) (*model.UserResponse, error) {
	return s.userRepository.GetUserById(id)
}

func (s *UserService) GetAllUsers() []*model.UserModel {
	return s.userRepository.GetAllUsers()
}

func (s *UserService) Login(user *model.UserModel) (*model.Token, error) {
	token, err := jwt.GenerateToken(user)
	if err != nil {
		return nil, err
	}
	return &model.Token{Token: token}, nil
}
