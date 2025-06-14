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

type PlayerHandler struct {
	Service *services.CRUDService[orm.Player]
}

func NewPlayerHandler(db *gorm.DB) *PlayerHandler {
	return &PlayerHandler{
		Service: services.NewCRUDService[orm.Player](db),
	}
}

func (h *PlayerHandler) Create(c *gin.Context) {
	var input dto.CreatePlayerDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	player := orm.Player{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Birthdate: input.Birthdate,
		Gender:    input.Gender,
		Height:    input.Height,
		Number:    input.Number,
		TeamID:    input.TeamID,
		AmpluaID:  input.AmpluaID,
	}

	if err := h.Service.Create(&player, "Team", "Amplua"); err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 23503") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Referenced team or amplua does not exist"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, player)
}

func (h *PlayerHandler) Get(c *gin.Context) {
	id := c.Param("id")
	var player orm.Player

	if err := h.Service.GetById(id, &player, "Team", "Amplua"); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Player with ID %s not found", id)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, player)
}

func (h *PlayerHandler) GetAll(c *gin.Context) {
	var players []orm.Player

	if err := h.Service.GetAll(&players, "player_id ASC", "Team", "Amplua"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, players)
}

func (h *PlayerHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var input dto.UpdatePlayerDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := map[string]interface{}{}
	if input.FirstName != nil {
		updates["first_name"] = *input.FirstName
	}
	if input.LastName != nil {
		updates["last_name"] = *input.LastName
	}
	if input.Birthdate != nil {
		updates["birthdate"] = *input.Birthdate
	}
	if input.Gender != nil {
		updates["gender"] = *input.Gender
	}
	if input.Height != nil {
		updates["height"] = *input.Height
	}
	if input.Number != nil {
		updates["number"] = *input.Number
	}
	if input.TeamID != nil {
		updates["team_id"] = *input.TeamID
	}
	if input.AmpluaID != nil {
		updates["amplua_id"] = *input.AmpluaID
	}

	if err := h.Service.Update(id, updates); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Player with ID %s not found", id)})
		} else if strings.Contains(err.Error(), "SQLSTATE 23503") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Referenced team or amplua does not exist"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var updated orm.Player
	_ = h.Service.GetById(id, &updated, "Team", "Amplua")

	c.JSON(http.StatusOK, updated)
}

func (h *PlayerHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	deleted, err := h.Service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Player with ID %s not found", id)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Player with ID %s deleted successfully", id)})
}
