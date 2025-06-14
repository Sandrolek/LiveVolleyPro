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

type SetHandler struct {
	Service *services.CRUDService[orm.Set]
}

func NewSetHandler(db *gorm.DB) *SetHandler {
	return &SetHandler{
		Service: services.NewCRUDService[orm.Set](db),
	}
}

func (h *SetHandler) Create(c *gin.Context) {
	var input dto.CreateSetDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	set := orm.Set{
		SerialNumber: input.SerialNumber,
		TeamScore:    input.TeamScore,
		OppScore:     input.OppScore,
		GameID:       input.GameID,
	}

	if err := h.Service.Create(&set, "Game", "SetActions"); err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 23503") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Referenced game does not exist"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, set)
}

func (h *SetHandler) Get(c *gin.Context) {
	id := c.Param("id")
	var set orm.Set

	if err := h.Service.GetById(id, &set, "Game", "SetActions"); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Set with ID %s not found", id)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, set)
}

func (h *SetHandler) GetAll(c *gin.Context) {
	var sets []orm.Set

	if err := h.Service.GetAll(&sets, "set_id ASC", "Game", "SetActions"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sets)
}

func (h *SetHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var input dto.UpdateSetDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := map[string]interface{}{}
	if input.SerialNumber != nil {
		updates["serial_number"] = *input.SerialNumber
	}
	if input.TeamScore != nil {
		updates["team_score"] = *input.TeamScore
	}
	if input.OppScore != nil {
		updates["opp_score"] = *input.OppScore
	}
	if input.GameID != nil {
		updates["game_id"] = *input.GameID
	}

	if err := h.Service.Update(id, updates); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Set with ID %s not found", id)})
		} else if strings.Contains(err.Error(), "SQLSTATE 23503") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Referenced game does not exist"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var updated orm.Set
	_ = h.Service.GetById(id, &updated, "Game", "SetActions")

	c.JSON(http.StatusOK, updated)
}

func (h *SetHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	deleted, err := h.Service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Set with ID %s not found", id)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Set with ID %s deleted successfully", id)})
}
