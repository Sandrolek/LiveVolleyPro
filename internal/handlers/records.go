package handlers

import (
	"net/http"

	"volley/internal/models/dto"
	"volley/internal/services"

	"github.com/gin-gonic/gin"
)

type RecordHandler struct {
	Service *services.SetActionService
}

func NewRecordHandler(service *services.SetActionService) *RecordHandler {
	return &RecordHandler{Service: service}
}

func (h *RecordHandler) RecordAction(c *gin.Context) {
	var input dto.RecordActionDTO
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
