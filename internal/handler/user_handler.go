package handler

import (
	"net/http"

	"user_service/internal/service"
	"user_service/pkg/errors"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// Register routes
func (h *UserHandler) RegisterRoutes(router *gin.Engine) {
	router.POST("/users", h.createUser)
	router.GET("/users/:id", h.getUser)
	router.GET("/users", h.getAllUsers)
	router.PUT("/users/:id", h.updateUser)
	router.DELETE("/users/:id", h.deleteUser)

}

// getAllUsers handler for GET /users
func (h *UserHandler) getAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"users": users,
	})
}

// Handler for creating a new user
func (h *UserHandler) createUser(c *gin.Context) {
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		h.handleError(c, errors.ErrInvalidInput)
		return
	}

	user, err := h.service.CreateUser(req.Name, req.Email)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"status":  "success",
		"data":    user,
		"message": "User created successfully",
	})
}

// Handler for GET /users/:id
func (h *UserHandler) getUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.service.GetUserByID(id)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "success",
		"data":    user,
		"message": "User retrieved successfully",
	})
}

// Handler for PUT /users/:id
func (h *UserHandler) updateUser(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		h.handleError(c, errors.ErrInvalidInput)
		return
	}

	user, err := h.service.UpdateUser(id, req.Name, req.Email)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "success",
		"data":    user,
		"message": "User updated successfully",
	})
}

// Handler for DELETE /users/:id
func (h *UserHandler) deleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeleteUser(id); err != nil {
		h.handleError(c, err)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "User deleted successfully",
	}) // 204
}

// Map domain errors to HTTP status codes
func (h *UserHandler) handleError(c *gin.Context, err error) {
	switch err {
	case errors.ErrNotFound:
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case errors.ErrAlreadyExists:
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	case errors.ErrInvalidInput:
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
