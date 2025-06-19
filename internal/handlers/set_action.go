package handlers

import (
	"net/http"

	"volley/internal/models/dto"
	"volley/internal/services"

	"github.com/gin-gonic/gin"
)

type SetActionHandler struct {
	Service *services.SetActionService
}

func NewSetActionHandler(service *services.SetActionService) *SetActionHandler {
	return &SetActionHandler{Service: service}
}

func (h *SetActionHandler) Create(c *gin.Context) {
	var input dto.CreateSetActionDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.RecordAction(input.SetID, input.Record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}
