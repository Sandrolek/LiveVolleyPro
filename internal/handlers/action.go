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

type ActionHandler struct {
	Service *services.CRUDService[orm.Action]
}

func NewActionHandler(db *gorm.DB) *ActionHandler {
	return &ActionHandler{
		Service: services.NewCRUDService[orm.Action](db),
	}
}

func (h *ActionHandler) Create(c *gin.Context) {
	var input dto.CreateActionDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	action := orm.Action{
		Name: input.Name,
	}

	if err := h.Service.Create(&action, "Rates"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, action)
}

func (h *ActionHandler) GetAll(c *gin.Context) {
	var actions []orm.Action
	if err := h.Service.GetAll(&actions, "action_id ASC", "Rates"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, actions)
}

func (h *ActionHandler) Get(c *gin.Context) {
	id := c.Param("id")
	var action orm.Action
	if err := h.Service.GetById(id, &action, "Rates"); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Action with ID %s not found", id)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, action)
}

func (h *ActionHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var input dto.UpdateActionDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := map[string]interface{}{}
	if input.Name != nil {
		updates["name"] = *input.Name
	}

	if err := h.Service.Update(id, updates); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Action with ID %s not found", id)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var updated orm.Action
	_ = h.Service.GetById(id, &updated, "Rates")
	c.JSON(http.StatusOK, updated)
}

func (h *ActionHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	deleted, err := h.Service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Action with ID %s not found", id)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Action with ID %s deleted successfully", id)})
}
