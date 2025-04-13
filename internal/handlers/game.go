package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"volley/internal/models"
)

type GameHandler struct {
	DB *gorm.DB
}

func NewGameHandler(db *gorm.DB) *GameHandler {
	return &GameHandler{DB: db}
}

// POST /games
func (h *GameHandler) CreateGame(c *gin.Context) {
	var input models.CreateGame
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	date, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}

	game := models.Game{
		RoundID:   input.RoundID,
		TeamID:    input.TeamID,
		OppTeamID: input.OppTeamID,
		Date:      date,
		Win:       input.Win,
	}

	if err := h.DB.Create(&game).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create game"})
		return
	}

	c.JSON(http.StatusCreated, game)
}

// GET /games
func (h *GameHandler) GetAllGames(c *gin.Context) {
	var games []models.Game
	if err := h.DB.Preload("Round").Preload("Team").Preload("OppTeam").Preload("Players").Preload("Sets").Find(&games).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch games"})
		return
	}
	c.JSON(http.StatusOK, games)
}

// GET /games/:id
func (h *GameHandler) GetGameByID(c *gin.Context) {
	id := c.Param("id")
	var game models.Game

	if err := h.DB.Preload("Round").Preload("Team").Preload("OppTeam").Preload("Players").Preload("Sets").First(&game, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}
	c.JSON(http.StatusOK, game)
}

// PUT /games/:id
func (h *GameHandler) UpdateGame(c *gin.Context) {
	id := c.Param("id")
	var input models.UpdateGame

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var game models.Game
	if err := h.DB.First(&game, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	if input.Date != "" {
		parsedDate, err := time.Parse("2006-01-02", input.Date)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
			return
		}
		game.Date = parsedDate
	}

	if input.Win != nil {
		game.Win = *input.Win
	}

	if err := h.DB.Save(&game).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update game"})
		return
	}

	c.JSON(http.StatusOK, game)
}

// DELETE /games/:id
func (h *GameHandler) DeleteGame(c *gin.Context) {
	id := c.Param("id")

	var game models.Game
	if err := h.DB.First(&game, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	if err := h.DB.Delete(&game).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete game"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game deleted"})
}
