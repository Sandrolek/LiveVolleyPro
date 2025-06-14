package handlers

import (
	"net/http"

	"fmt"

	"volley/internal/models/dto"
	"volley/internal/models/orm"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ChampionshipHandler struct {
	DB *gorm.DB
}

func NewChampionshipHandler(db *gorm.DB) *ChampionshipHandler {
	return &ChampionshipHandler{DB: db}
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

	if err := h.DB.Create(&champ).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, champ)
}

func (h *ChampionshipHandler) GetAll(c *gin.Context) {
	var champs []orm.Championship
	h.DB.Find(&champs)
	c.JSON(http.StatusOK, champs)
}

func (h *ChampionshipHandler) Get(c *gin.Context) {
	id := c.Param("id")
	var champ orm.Championship
	if err := h.DB.First(&champ, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.JSON(http.StatusOK, champ)
}

func (h *ChampionshipHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var champ orm.Championship
	if err := h.DB.First(&champ, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	var input dto.UpdateChampionshipDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Model(&champ).Updates(input)
	c.JSON(http.StatusOK, champ)
}

func (h *ChampionshipHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	result := h.DB.Delete(&orm.Championship{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Championship with ID %s not found", id),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Championship with ID %s deleted successfully", id),
	})
}
