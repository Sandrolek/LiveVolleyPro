package handlers

import (
	"net/http"

	"volley/internal/models/dto"
	"volley/internal/services"

	"github.com/gin-gonic/gin"
)

type StatsHandler struct {
	Service *services.StatsService
}

func NewStatsHandler(service *services.StatsService) *StatsHandler {
	return &StatsHandler{Service: service}
}

func (h *StatsHandler) GetPlayerStats(c *gin.Context) {
	var input dto.PlayerStatsRequestDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stats, err := h.Service.GetPlayerStats(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stats)
}
