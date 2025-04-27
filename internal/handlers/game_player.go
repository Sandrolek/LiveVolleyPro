package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"volley/internal/models"
)

type GamePlayerHandler struct {
	DB *gorm.DB
}

func NewGamePlayerHandler(db *gorm.DB) *GamePlayerHandler {
	return &GamePlayerHandler{DB: db}
}

// POST /game_players
func (h *GamePlayerHandler) CreateGamePlayer(c *gin.Context) {
	var input models.CreateGamePlayer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gamePlayer := models.GamePlayer{
		PlayerID: input.PlayerID,
		GameID:   input.GameID,
	}

	if err := h.DB.Create(&gamePlayer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create game_player"})
		return
	}

	c.JSON(http.StatusCreated, gamePlayer)
}

// GET /game_players
func (h *GamePlayerHandler) GetAllGamePlayers(c *gin.Context) {
	var gamePlayers []models.GamePlayer
	if err := h.DB.Preload("Player").Preload("Game").Find(&gamePlayers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch game_players"})
		return
	}
	c.JSON(http.StatusOK, gamePlayers)
}

// GET /game_players/:id
func (h *GamePlayerHandler) GetGamePlayerByID(c *gin.Context) {
	id := c.Param("id")
	var gamePlayer models.GamePlayer

	if err := h.DB.Preload("Player").Preload("Game").First(&gamePlayer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "GamePlayer not found"})
		return
	}

	c.JSON(http.StatusOK, gamePlayer)
}

// PUT /game_players/:id
func (h *GamePlayerHandler) UpdateGamePlayer(c *gin.Context) {
	id := c.Param("id")
	var input models.UpdateGamePlayer

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var gamePlayer models.GamePlayer
	if err := h.DB.First(&gamePlayer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "GamePlayer not found"})
		return
	}

	if input.PlayerID != nil {
		gamePlayer.PlayerID = *input.PlayerID
	}

	if input.GameID != nil {
		gamePlayer.GameID = *input.GameID
	}

	if err := h.DB.Save(&gamePlayer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update game_player"})
		return
	}

	c.JSON(http.StatusOK, gamePlayer)
}

// DELETE /game_players/:id
func (h *GamePlayerHandler) DeleteGamePlayer(c *gin.Context) {
	id := c.Param("id")

	var gamePlayer models.GamePlayer
	if err := h.DB.First(&gamePlayer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "GamePlayer not found"})
		return
	}

	if err := h.DB.Delete(&gamePlayer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete game_player"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "GamePlayer deleted"})
}
