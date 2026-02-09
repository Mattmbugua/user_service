package service

import (
	"errors"
	"testing"
	"user_service/internal/model"
)

type mockUserRepository struct {
	createFn func(user *model.User) (*model.User, error)
}

func (m *mockUserRepository) Create(user *model.User) (*model.User, error) {
	return m.createFn(user)
}

func (m *mockUserRepository) GetByID(id string) (model.User, error) {
	return model.User{}, nil
}

func (m *mockUserRepository) GetAll() ([]*model.User, error) {
	return []*model.User{}, nil
}

func (m *mockUserRepository) Update(user *model.User) (*model.User, error) {
	return user, nil
}

func (m *mockUserRepository) Delete(id string) error {
	return nil
}

func TestCreateUser_Success(t *testing.T) {
	mockRepo := &mockUserRepository{
		createFn: func(user *model.User) (*model.User, error) {
			user.ID = "123"
			return user, nil
		},
	}

	service := NewUserService(mockRepo)

	user, err := service.CreateUser("John Doe", "JOHN@EMAIL.COM")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if user.Email != "john@email.com" {
		t.Errorf("expected email to be lowercased, got %s", user.Email)
	}

	if user.ID == "" {
		t.Errorf("expected user ID to be set")
	}
}
func TestCreateUser_RepositoryError(t *testing.T) {
	mockRepo := &mockUserRepository{
		createFn: func(user *model.User) (*model.User, error) {
			return nil, errors.New("database error")
		},
	}

	service := NewUserService(mockRepo)

	_, err := service.CreateUser("Jane", "jane@email.com")

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
