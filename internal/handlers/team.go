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

	team := orm.Team{
		Name:   input.Name,
		UserID: input.UserID,
	}

	if err := h.Service.Create(&team); err != nil {

		fmt.Println(err)
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
	var team orm.Team

	if err := h.Service.GetById(id, &team, "User", "Players"); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Team with ID %s not found", id)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, team)
}

func (h *TeamHandler) GetAll(c *gin.Context) {
	var teams []orm.Team

	if err := h.Service.GetAll(&teams, "team_id ASC", "User", "Players"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, teams)
}

func (h *TeamHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var input dto.UpdateTeamDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := map[string]interface{}{}
	if input.Name != nil {
		updates["name"] = *input.Name
	}
	if input.UserID != nil {
		updates["user_id"] = *input.UserID
	}

	if err := h.Service.Update(id, updates); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Team with ID %s not found", id)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var updated orm.Team
	_ = h.Service.GetById(id, &updated, "User", "Players")

	c.JSON(http.StatusOK, updated)
}

func (h *TeamHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	deleted, err := h.Service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Team with ID %s not found", id)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Team with ID %s deleted successfully", id)})
}
