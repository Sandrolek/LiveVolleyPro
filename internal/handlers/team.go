package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"volley/internal/models"
)

type TeamHandler struct {
	DB *gorm.DB
}

func NewTeamHandler(db *gorm.DB) *TeamHandler {
	return &TeamHandler{DB: db}
}

// POST /teams
func (h *TeamHandler) CreateTeam(c *gin.Context) {
	var input models.CreateTeam
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	team := models.Team{
		UserID: input.UserID,
		Name:   input.Name,
	}

	if err := h.DB.Create(&team).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create team"})
		return
	}

	c.JSON(http.StatusCreated, team)
}

// GET /teams
func (h *TeamHandler) GetAllTeams(c *gin.Context) {
	var teams []models.Team
	if err := h.DB.Preload("User").Preload("Players").Find(&teams).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch teams"})
		return
	}
	c.JSON(http.StatusOK, teams)
}

// GET /teams/:id
func (h *TeamHandler) GetTeamByID(c *gin.Context) {
	id := c.Param("id")
	var team models.Team

	if err := h.DB.Preload("User").Preload("Players").First(&team, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}
	c.JSON(http.StatusOK, team)
}

// PUT /teams/:id
func (h *TeamHandler) UpdateTeam(c *gin.Context) {
	id := c.Param("id")
	var input models.UpdateTeam

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var team models.Team
	if err := h.DB.First(&team, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}

	team.Name = input.Name
	if err := h.DB.Save(&team).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update team"})
		return
	}

	c.JSON(http.StatusOK, team)
}

// DELETE /teams/:id
func (h *TeamHandler) DeleteTeam(c *gin.Context) {
	id := c.Param("id")

	var team models.Team
	if err := h.DB.First(&team, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}

	if err := h.DB.Delete(&team).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete team"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Team deleted"})
}
