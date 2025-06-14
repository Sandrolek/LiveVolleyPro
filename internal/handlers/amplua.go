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

type AmpluaHandler struct {
	Service *services.CRUDService[orm.Amplua]
}

func NewAmpluaHandler(db *gorm.DB) *AmpluaHandler {
	return &AmpluaHandler{
		Service: services.NewCRUDService[orm.Amplua](db),
	}
}

func (h *AmpluaHandler) Create(c *gin.Context) {
	var input dto.CreateAmpluaDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	amplua := orm.Amplua{
		Name: input.Name,
	}

	if err := h.Service.Create(&amplua, "Players"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, amplua)
}

func (h *AmpluaHandler) Get(c *gin.Context) {
	id := c.Param("id")
	var amplua orm.Amplua

	if err := h.Service.GetById(id, &amplua, "Players"); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Amplua with ID %s not found", id)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, amplua)
}

func (h *AmpluaHandler) GetAll(c *gin.Context) {
	var ampluas []orm.Amplua

	if err := h.Service.GetAll(&ampluas, "amplua_id ASC", "Players"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ampluas)
}

func (h *AmpluaHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var input dto.UpdateAmpluaDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := map[string]interface{}{}
	if input.Name != nil {
		updates["name"] = *input.Name
	}

	if err := h.Service.Update(id, updates); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Amplua with ID %s not found", id)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var updated orm.Amplua
	_ = h.Service.GetById(id, &updated, "Players")

	c.JSON(http.StatusOK, updated)
}

func (h *AmpluaHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	deleted, err := h.Service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Amplua with ID %s not found", id)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Amplua with ID %s deleted successfully", id)})
}
