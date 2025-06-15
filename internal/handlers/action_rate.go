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

type ActionRateHandler struct {
	Service *services.CRUDService[orm.ActionRate]
}

func NewActionRateHandler(db *gorm.DB) *ActionRateHandler {
	return &ActionRateHandler{
		Service: services.NewCRUDService[orm.ActionRate](db),
	}
}

func (h *ActionRateHandler) Create(c *gin.Context) {
	var input dto.CreateActionRateDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rate := orm.ActionRate{
		HelpText:  input.HelpText,
		Signature: input.Signature,
		ActionID:  input.ActionID,
	}

	if err := h.Service.Create(&rate, "Action"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, rate)
}

func (h *ActionRateHandler) GetAll(c *gin.Context) {
	var rates []orm.ActionRate
	if err := h.Service.GetAll(&rates, "action_rate_id ASC", "Action"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rates)
}

func (h *ActionRateHandler) Get(c *gin.Context) {
	id := c.Param("id")
	var rate orm.ActionRate
	if err := h.Service.GetById(id, &rate, "Action"); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("ActionRate with ID %s not found", id)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, rate)
}

func (h *ActionRateHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var input dto.UpdateActionRateDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := map[string]interface{}{}
	if input.HelpText != nil {
		updates["help_text"] = *input.HelpText
	}
	if input.Signature != nil {
		updates["signature"] = *input.Signature
	}
	if input.ActionID != nil {
		updates["action_id"] = *input.ActionID
	}

	if err := h.Service.Update(id, updates); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("ActionRate with ID %s not found", id)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var updated orm.ActionRate
	_ = h.Service.GetById(id, &updated, "Action")
	c.JSON(http.StatusOK, updated)
}

func (h *ActionRateHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	deleted, err := h.Service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("ActionRate with ID %s not found", id)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("ActionRate with ID %s deleted successfully", id)})
}
