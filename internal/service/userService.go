package service

import (
	"strings"

	"user_service/internal/model"
	"user_service/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(name, email string) (*model.User, error) {

	user := &model.User{
		Name:  name,
		Email: strings.ToLower(email),
	}

	// Call repository to create user
	createdUser, err := s.repo.Create(user)

	if err != nil {
		return nil, err
	}
	sendWelcomeEmailAsync(user.Name, user.Email)
	return createdUser, nil
}

// GetUserByID
func (s *UserService) GetUserByID(id string) (*model.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAllUsers

func (s *UserService) GetAllUsers() ([]*model.User, error) {
	return s.repo.GetAll()
}

//update user using id

func (s *UserService) UpdateUser(id, name, email string) (*model.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if name != "" {
		user.Name = name
	}
	if email != "" {
		user.Email = strings.ToLower(email)
	}

	return s.repo.Update(&user)
}

// Delete a user by ID
func (s *UserService) DeleteUser(id string) error {
	return s.repo.Delete(id)
}
