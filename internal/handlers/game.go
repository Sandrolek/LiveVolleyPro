package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"volley/internal/models/dto"
	"volley/internal/models/orm"
	"volley/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GameHandler struct {
	Service *services.CRUDService[orm.Game]
}

func NewGameHandler(db *gorm.DB) *GameHandler {
	return &GameHandler{
		Service: services.NewCRUDService[orm.Game](db),
	}
}

func (h *GameHandler) Create(c *gin.Context) {
	var input dto.CreateGameDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	game := orm.Game{
		Date:      input.Date,
		Win:       input.Win,
		RoundID:   input.RoundID,
		TeamID:    input.TeamID,
		OppTeamID: input.OppTeamID,
	}

	if err := h.Service.Create(&game, "Round", "Team", "OppTeam", "Sets", "GamePlayers"); err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 23503") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "One or more referenced IDs (Round, Team, OppTeam) do not exist"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, game)
}

func (h *GameHandler) Get(c *gin.Context) {
	id := c.Param("id")
	var game orm.Game

	if err := h.Service.GetById(id, &game, "Round", "Team", "OppTeam", "Sets", "GamePlayers"); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Game with ID %s not found", id)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, game)
}

func (h *GameHandler) GetAll(c *gin.Context) {
	var games []orm.Game

	if err := h.Service.GetAll(&games, "game_id ASC", "Round", "Team", "OppTeam", "Sets", "GamePlayers"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, games)
}

func (h *GameHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var input dto.UpdateGameDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := make(map[string]interface{})
	if input.Date != nil {
		updates["date"] = *input.Date
	}
	if input.Win != nil {
		updates["win"] = *input.Win
	}
	if input.RoundID != nil {
		updates["round_id"] = *input.RoundID
	}
	if input.TeamID != nil {
		updates["team_id"] = *input.TeamID
	}
	if input.OppTeamID != nil {
		updates["opp_team_id"] = *input.OppTeamID
	}

	if err := h.Service.Update(id, updates); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Game with ID %s not found", id)})
		} else if strings.Contains(err.Error(), "SQLSTATE 23503") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "One or more referenced IDs do not exist"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var updated orm.Game
	_ = h.Service.GetById(id, &updated, "Round", "Team", "OppTeam", "Sets", "GamePlayers")

	c.JSON(http.StatusOK, updated)
}

func (h *GameHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	deleted, err := h.Service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Game with ID %s not found", id)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Game with ID %s deleted successfully", id)})
}
