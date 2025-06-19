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

type TeamHandler struct {
	Service *services.CRUDService[orm.Team]
}

func NewTeamHandler(db *gorm.DB) *TeamHandler {
	return &TeamHandler{
		Service: services.NewCRUDService[orm.Team](db),
	}
}

func (h *TeamHandler) Create(c *gin.Context) {
	var input dto.CreateTeamDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("userID").(uint)

	team := orm.Team{
		Name:   input.Name,
		UserID: int(userID),
	}

	if err := h.Service.Create(&team); err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 23503") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User with the specified ID does not exist"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, team)
}

func (h *TeamHandler) Get(c *gin.Context) {
	id := c.Param("id")
	userID := c.MustGet("userID").(uint)

	var team orm.Team
	if err := h.Service.GetById(id, &team, "User", "Players"); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Team with ID %s not found", id)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if team.UserID != int(userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have access to this team"})
		return
	}

	c.JSON(http.StatusOK, team)
}

func (h *TeamHandler) GetAll(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var teams []orm.Team
	if err := h.Service.GetWhere(&teams, "user_id = ?", userID, "team_id ASC", "User", "Players"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, teams)
}

func (h *TeamHandler) Update(c *gin.Context) {
	id := c.Param("id")
	userID := c.MustGet("userID").(uint)

	var team orm.Team
	if err := h.Service.GetById(id, &team); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Team with ID %s not found", id)})
		return
	}

	if team.UserID != int(userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to update this team"})
		return
	}

	var input dto.UpdateTeamDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := map[string]interface{}{}
	if input.Name != nil {
		updates["name"] = *input.Name
	}

	if err := h.Service.Update(id, updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updated orm.Team
	_ = h.Service.GetById(id, &updated, "User", "Players")

	c.JSON(http.StatusOK, updated)
}

func (h *TeamHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	userID := c.MustGet("userID").(uint)

	var team orm.Team
	if err := h.Service.GetById(id, &team); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Team with ID %s not found", id)})
		return
	}

	if team.UserID != int(userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to delete this team"})
		return
	}

	_, err := h.Service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Team with ID %s deleted successfully", id)})
}
