package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"volley/internal/models"
)

type RoundHandler struct {
	DB *gorm.DB
}

func NewRoundHandler(db *gorm.DB) *RoundHandler {
	return &RoundHandler{DB: db}
}

// POST /rounds
func (h *RoundHandler) CreateRound(c *gin.Context) {
	var input models.CreateRound
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var startDate, endDate *time.Time

	if input.StartDate != "" {
		sd, err := time.Parse("2006-01-02", input.StartDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format"})
			return
		}
		startDate = &sd
	}

	if input.EndDate != "" {
		ed, err := time.Parse("2006-01-02", input.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format"})
			return
		}
		endDate = &ed
	}

	round := models.Round{
		SerialNumber:   input.SerialNumber,
		ChampionshipID: input.ChampionshipID,
		StartDate:      startDate,
		EndDate:        endDate,
	}

	if err := h.DB.Create(&round).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create round"})
		return
	}

	c.JSON(http.StatusCreated, round)
}

// GET /rounds
func (h *RoundHandler) GetAllRounds(c *gin.Context) {
	var rounds []models.Round
	if err := h.DB.Preload("Championship").Preload("Games").Find(&rounds).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch rounds"})
		return
	}
	c.JSON(http.StatusOK, rounds)
}

// GET /rounds/:id
func (h *RoundHandler) GetRoundByID(c *gin.Context) {
	id := c.Param("id")
	var round models.Round

	if err := h.DB.Preload("Championship").Preload("Games").First(&round, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Round not found"})
		return
	}
	c.JSON(http.StatusOK, round)
}

// PUT /rounds/:id
func (h *RoundHandler) UpdateRound(c *gin.Context) {
	id := c.Param("id")
	var input models.UpdateRound

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var round models.Round
	if err := h.DB.First(&round, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Round not found"})
		return
	}

	if input.SerialNumber != nil {
		round.SerialNumber = *input.SerialNumber
	}

	if input.StartDate != "" {
		startDate, err := time.Parse("2006-01-02", input.StartDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format"})
			return
		}
		round.StartDate = &startDate
	}

	if input.EndDate != "" {
		endDate, err := time.Parse("2006-01-02", input.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format"})
			return
		}
		round.EndDate = &endDate
	}

	if err := h.DB.Save(&round).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update round"})
		return
	}

	c.JSON(http.StatusOK, round)
}

// DELETE /rounds/:id
func (h *RoundHandler) DeleteRound(c *gin.Context) {
	id := c.Param("id")

	var round models.Round
	if err := h.DB.First(&round, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Round not found"})
		return
	}

	if err := h.DB.Delete(&round).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete round"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Round deleted"})
}
