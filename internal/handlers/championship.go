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

type ChampionshipHandler struct {
	Service *services.CRUDService[orm.Championship]
}

func NewChampionshipHandler(db *gorm.DB) *ChampionshipHandler {
	return &ChampionshipHandler{
		Service: services.NewCRUDService[orm.Championship](db),
	}
}

func (h *ChampionshipHandler) Create(c *gin.Context) {
	var input dto.CreateChampionshipDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	champ := orm.Championship{
		Title:     input.Title,
		StartDate: input.StartDate,
		EndDate:   input.EndDate,
	}

	if err := h.Service.Create(&champ); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, champ)
}

func (h *ChampionshipHandler) GetAll(c *gin.Context) {
	var champs []orm.Championship

	if err := h.Service.GetAll(&champs, "championship_id ASC", "Rounds"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, champs)
}

func (h *ChampionshipHandler) Get(c *gin.Context) {
	id := c.Param("id")
	var champ orm.Championship

	if err := h.Service.GetById(id, &champ, "Rounds"); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Championship with ID %s not found", id)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, champ)
}
func (h *ChampionshipHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var input dto.UpdateChampionshipDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := make(map[string]interface{})
	if input.Title != nil {
		updates["title"] = *input.Title
	}
	if input.StartDate != nil {
		updates["start_date"] = *input.StartDate
	}
	if input.EndDate != nil {
		updates["end_date"] = *input.EndDate
	}

	if err := h.Service.Update(id, updates); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Championship with ID %s not found", id)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Optional: return updated entity
	var updated orm.Championship
	_ = h.Service.GetById(id, &updated)

	c.JSON(http.StatusOK, updated)
}

func (h *ChampionshipHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	deleted, err := h.Service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Championship with ID %s not found", id)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Championship with ID %s deleted successfully", id)})
}
