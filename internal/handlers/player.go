package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"volley/internal/models"
)

type PlayerHandler struct {
	DB *gorm.DB
}

func NewPlayerHandler(db *gorm.DB) *PlayerHandler {
	return &PlayerHandler{DB: db}
}

// POST /players
func (h *PlayerHandler) CreatePlayer(c *gin.Context) {
	var input models.CreatePlayer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var birthdatePtr *time.Time
	if input.Birthdate != "" {
		birthdate, err := time.Parse("2006-01-02", input.Birthdate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid birthdate format. Use YYYY-MM-DD"})
			return
		}
		birthdatePtr = &birthdate
	}

	player := models.Player{
		AmpluaID:  input.AmpluaID,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Birthdate: birthdatePtr,
		Gender:    input.Gender,
		Height:    input.Height,
		Number:    input.Number,
		TeamID:    input.TeamID,
	}

	if err := h.DB.Create(&player).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create player"})
		return
	}

	c.JSON(http.StatusCreated, player)
}

// GET /players
func (h *PlayerHandler) GetAllPlayers(c *gin.Context) {
	var players []models.Player
	if err := h.DB.Preload("Amplua").Preload("Team").Find(&players).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch players"})
		return
	}
	c.JSON(http.StatusOK, players)
}

// GET /players/:id
func (h *PlayerHandler) GetPlayerByID(c *gin.Context) {
	id := c.Param("id")
	var player models.Player

	if err := h.DB.Preload("Amplua").Preload("Team").First(&player, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}
	c.JSON(http.StatusOK, player)
}

// PUT /players/:id
func (h *PlayerHandler) UpdatePlayer(c *gin.Context) {
	id := c.Param("id")
	var input models.UpdatePlayer

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var player models.Player
	if err := h.DB.First(&player, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	// Обновляем только переданные поля
	if input.FirstName != "" {
		player.FirstName = input.FirstName
	}
	if input.LastName != "" {
		player.LastName = input.LastName
	}
	if input.Height != nil {
		player.Height = input.Height
	}
	if input.Number != nil {
		player.Number = input.Number
	}

	if err := h.DB.Save(&player).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update player"})
		return
	}

	c.JSON(http.StatusOK, player)
}

// DELETE /players/:id
func (h *PlayerHandler) DeletePlayer(c *gin.Context) {
	id := c.Param("id")

	var player models.Player
	if err := h.DB.First(&player, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	if err := h.DB.Delete(&player).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete player"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Player deleted"})
}
