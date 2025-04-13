package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"volley/internal/models"
)

type SetActionHandler struct {
	DB *gorm.DB
}

func NewSetActionHandler(db *gorm.DB) *SetActionHandler {
	return &SetActionHandler{DB: db}
}

// POST /set_actions
func (h *SetActionHandler) CreateSetAction(c *gin.Context) {
	var input models.CreateSetAction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	setAction := models.SetAction{
		SetID:        input.SetID,
		PlayerID:     input.PlayerID,
		ActionRateID: input.ActionRateID,
	}

	if err := h.DB.Create(&setAction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create set_action"})
		return
	}

	c.JSON(http.StatusCreated, setAction)
}

// GET /set_actions
func (h *SetActionHandler) GetAllSetActions(c *gin.Context) {
	var setActions []models.SetAction
	if err := h.DB.Preload("Set").Preload("Player").Preload("ActionRate").Find(&setActions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch set_actions"})
		return
	}
	c.JSON(http.StatusOK, setActions)
}

// GET /set_actions/:id
func (h *SetActionHandler) GetSetActionByID(c *gin.Context) {
	id := c.Param("id")
	var setAction models.SetAction

	if err := h.DB.Preload("Set").Preload("Player").Preload("ActionRate").First(&setAction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SetAction not found"})
		return
	}

	c.JSON(http.StatusOK, setAction)
}

// PUT /set_actions/:id
func (h *SetActionHandler) UpdateSetAction(c *gin.Context) {
	id := c.Param("id")
	var input models.UpdateSetAction

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var setAction models.SetAction
	if err := h.DB.First(&setAction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SetAction not found"})
		return
	}

	if input.SetID != nil {
		setAction.SetID = *input.SetID
	}

	if input.PlayerID != nil {
		setAction.PlayerID = *input.PlayerID
	}

	if input.ActionRateID != nil {
		setAction.ActionRateID = *input.ActionRateID
	}

	if err := h.DB.Save(&setAction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update set_action"})
		return
	}

	c.JSON(http.StatusOK, setAction)
}

// DELETE /set_actions/:id
func (h *SetActionHandler) DeleteSetAction(c *gin.Context) {
	id := c.Param("id")

	var setAction models.SetAction
	if err := h.DB.First(&setAction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "SetAction not found"})
		return
	}

	if err := h.DB.Delete(&setAction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete set_action"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "SetAction deleted"})
}
