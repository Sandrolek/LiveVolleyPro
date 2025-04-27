package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"volley/internal/models"
)

type SetHandler struct {
	DB *gorm.DB
}

func NewSetHandler(db *gorm.DB) *SetHandler {
	return &SetHandler{DB: db}
}

// POST /sets
func (h *SetHandler) CreateSet(c *gin.Context) {
	var input models.CreateSet
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	set := models.Set{
		GameID:       input.GameID,
		SerialNumber: input.SerialNumber,
		TeamScore:    input.TeamScore,
		OppScore:     input.OppScore,
	}

	if err := h.DB.Create(&set).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create set"})
		return
	}

	c.JSON(http.StatusCreated, set)
}

// GET /sets
func (h *SetHandler) GetAllSets(c *gin.Context) {
	var sets []models.Set
	if err := h.DB.Preload("Game").Preload("SetActions").Find(&sets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch sets"})
		return
	}
	c.JSON(http.StatusOK, sets)
}

// GET /sets/:id
func (h *SetHandler) GetSetByID(c *gin.Context) {
	id := c.Param("id")
	var set models.Set

	if err := h.DB.Preload("Game").Preload("SetActions").First(&set, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Set not found"})
		return
	}

	c.JSON(http.StatusOK, set)
}

// PUT /sets/:id
func (h *SetHandler) UpdateSet(c *gin.Context) {
	id := c.Param("id")
	var input models.UpdateSet

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var set models.Set
	if err := h.DB.First(&set, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Set not found"})
		return
	}

	if input.SerialNumber != nil {
		set.SerialNumber = *input.SerialNumber
	}

	if input.TeamScore != nil {
		set.TeamScore = *input.TeamScore
	}

	if input.OppScore != nil {
		set.OppScore = *input.OppScore
	}

	if err := h.DB.Save(&set).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update set"})
		return
	}

	c.JSON(http.StatusOK, set)
}

// DELETE /sets/:id
func (h *SetHandler) DeleteSet(c *gin.Context) {
	id := c.Param("id")

	var set models.Set
	if err := h.DB.First(&set, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Set not found"})
		return
	}

	if err := h.DB.Delete(&set).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete set"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Set deleted"})
}
