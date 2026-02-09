package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"user_service/internal/repository"
	"user_service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() (*gin.Engine, *service.UserService) {
	// In-memory repo
	userRepo := repository.NewInMemoryUserRepository()
	userService := service.NewUserService(userRepo)
	handler := NewUserHandler(userService)

	// Gin router
	router := gin.Default()
	handler.RegisterRoutes(router)

	return router, userService
}

func TestCreateUser(t *testing.T) {
	router, _ := setupTestRouter()

	payload := map[string]string{
		"name":  "Alice",
		"email": "alice@example.com",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var res map[string]interface{}
	err := json.Unmarshal(resp.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.Equal(t, "success", res["status"])
}

func TestGetAllUsers(t *testing.T) {
	router, userService := setupTestRouter()

	// Add a user first
	_, _ = userService.CreateUser("Bob", "bob@example.com")

	req, _ := http.NewRequest("GET", "/users", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var res map[string]interface{}
	err := json.Unmarshal(resp.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.NotEmpty(t, res["users"])
}

func TestGetUserByID(t *testing.T) {
	router, userService := setupTestRouter()

	user, _ := userService.CreateUser("Charlie", "charlie@example.com")

	req, _ := http.NewRequest("GET", "/users/"+user.ID, nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var res map[string]interface{}
	err := json.Unmarshal(resp.Body.Bytes(), &res)
	assert.NoError(t, err)
	data := res["data"].(map[string]interface{})
	assert.Equal(t, "Charlie", data["name"].(string))
}

func TestDeleteUser(t *testing.T) {
	router, userService := setupTestRouter()

	user, _ := userService.CreateUser("Dave", "dave@example.com")

	req, _ := http.NewRequest("DELETE", "/users/"+user.ID, nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
