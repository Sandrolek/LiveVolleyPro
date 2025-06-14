package handlers

import (
	"fmt"
	"net/http"

	"volley/internal/models/dto"
	"volley/internal/models/orm"
	"volley/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoundHandler struct {
	Service *services.CRUDService[orm.Round]
}

func NewRoundHandler(db *gorm.DB) *RoundHandler {
	return &RoundHandler{
		Service: services.NewCRUDService[orm.Round](db),
	}
}

func (h *RoundHandler) Create(c *gin.Context) {
	var input dto.CreateRoundDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	round := orm.Round{
		SerialNumber:   input.SerialNumber,
		StartDate:      input.StartDate,
		EndDate:        input.EndDate,
		ChampionshipID: input.ChampionshipID,
	}

	if err := h.Service.Create(&round); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, round)
}

func (h *RoundHandler) Get(c *gin.Context) {
	id := c.Param("id")
	var round orm.Round

	if err := h.Service.GetById(id, &round, "Championship"); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Round with ID %s not found", id)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, round)
}

func (h *RoundHandler) GetAll(c *gin.Context) {
	var rounds []orm.Round

	if err := h.Service.GetAll(&rounds, "round_id ASC", "Championship"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rounds)
}

func (h *RoundHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var input dto.UpdateRoundDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := make(map[string]interface{})
	if input.SerialNumber != nil {
		updates["serial_number"] = *input.SerialNumber
	}
	if input.StartDate != nil {
		updates["start_date"] = *input.StartDate
	}
	if input.EndDate != nil {
		updates["end_date"] = *input.EndDate
	}
	if input.ChampionshipID != nil {
		updates["championship_id"] = *input.ChampionshipID
	}

	if err := h.Service.Update(id, updates); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Round with ID %s not found", id)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var updated orm.Round
	_ = h.Service.GetById(id, &updated)
	c.JSON(http.StatusOK, updated)
}

func (h *RoundHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	deleted, err := h.Service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Round with ID %s not found", id)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Round with ID %s deleted successfully", id)})
}
