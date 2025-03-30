package handlers

import (
	"net/http"
	"volley/internal/models"
	"volley/internal/repositories"
	"volley/internal/services"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary Get info about all users
// @Schemes
// @Description Retrieves data
// @Tags authentication
// @Accept json
// @Produce json
// @Success 200 {string} {"users": []}
// @Router /users [get]
func GetAllUsers(c *gin.Context) {
	var UserService = services.NewUserService(repositories.NewUserRepository())

	users, err := UserService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// @BasePath /api/v1

// PingExample godoc
// @Summary Create user
// @Schemes
// @Description Creates user
// @Tags authentication
// @Accept json
// @Produce json
// @Success 200 {string} {"message": "created user"}
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var UserService = services.NewUserService(repositories.NewUserRepository())

	user, err := UserService.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
