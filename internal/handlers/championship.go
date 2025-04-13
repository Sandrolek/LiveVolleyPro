package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"volley/internal/models"
)

type ChampionshipHandler struct {
	DB *gorm.DB
}

func NewChampionshipHandler(db *gorm.DB) *ChampionshipHandler {
	return &ChampionshipHandler{DB: db}
}

// POST /championships
func (h *ChampionshipHandler) CreateChampionship(c *gin.Context) {
	var input models.CreateChampionship
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startDate, err := time.Parse("2006-01-02", input.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format"})
		return
	}

	endDate, err := time.Parse("2006-01-02", input.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format"})
		return
	}

	championship := models.Championship{
		Title:     input.Title,
		StartDate: startDate,
		EndDate:   endDate,
	}

	if err := h.DB.Create(&championship).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create championship"})
		return
	}

	c.JSON(http.StatusCreated, championship)
}

// GET /championships
func (h *ChampionshipHandler) GetAllChampionships(c *gin.Context) {
	var championships []models.Championship
	if err := h.DB.Preload("Rounds").Find(&championships).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch championships"})
		return
	}
	c.JSON(http.StatusOK, championships)
}

// GET /championships/:id
func (h *ChampionshipHandler) GetChampionshipByID(c *gin.Context) {
	id := c.Param("id")
	var championship models.Championship

	if err := h.DB.Preload("Rounds").First(&championship, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Championship not found"})
		return
	}

	c.JSON(http.StatusOK, championship)
}

// PUT /championships/:id
func (h *ChampionshipHandler) UpdateChampionship(c *gin.Context) {
	id := c.Param("id")
	var input models.UpdateChampionship

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var championship models.Championship
	if err := h.DB.First(&championship, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Championship not found"})
		return
	}

	if input.Title != nil {
		championship.Title = *input.Title
	}

	if input.StartDate != "" {
		startDate, err := time.Parse("2006-01-02", input.StartDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format"})
			return
		}
		championship.StartDate = startDate
	}

	if input.EndDate != "" {
		endDate, err := time.Parse("2006-01-02", input.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format"})
			return
		}
		championship.EndDate = endDate
	}

	if err := h.DB.Save(&championship).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update championship"})
		return
	}

	c.JSON(http.StatusOK, championship)
}

// DELETE /championships/:id
func (h *ChampionshipHandler) DeleteChampionship(c *gin.Context) {
	id := c.Param("id")

	var championship models.Championship
	if err := h.DB.First(&championship, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Championship not found"})
		return
	}

	if err := h.DB.Delete(&championship).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete championship"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Championship deleted"})
}
