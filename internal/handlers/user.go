package handlers

import (
	"fmt"
	"net/http"

	"volley/internal/models/dto"
	"volley/internal/models/orm"
	"volley/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	Service *services.CRUDService[orm.User]
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		Service: services.NewCRUDService[orm.User](db),
	}
}

func (h *UserHandler) Create(c *gin.Context) {
	var input dto.CreateUserDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := orm.User{
		Name:     input.Name,
		Password: input.Password,
		Email:    input.Email,
	}

	if err := h.Service.Create(&user, "Teams"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) Get(c *gin.Context) {
	id := c.Param("id")
	var user orm.User

	if err := h.Service.GetById(id, &user, "Teams"); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User with ID %s not found", id)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetAll(c *gin.Context) {
	var users []orm.User

	if err := h.Service.GetAll(&users, "user_id ASC", "Teams"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var input dto.UpdateUserDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := map[string]interface{}{}
	if input.Name != nil {
		updates["name"] = *input.Name
	}
	if input.Password != nil {
		updates["password"] = *input.Password
	}
	if input.Email != nil {
		updates["email"] = *input.Email
	}

	if err := h.Service.Update(id, updates); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User with ID %s not found", id)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var updated orm.User
	_ = h.Service.GetById(id, &updated, "Teams")

	c.JSON(http.StatusOK, updated)
}

func (h *UserHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	deleted, err := h.Service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User with ID %s not found", id)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User with ID %s deleted successfully", id)})
}
