package handlers

import (
	"fmt"
	"hack4good/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CaregiverHandler struct {
	DB *gorm.DB
}

func (h CaregiverHandler) List(c *gin.Context) {
	var caregivers []models.Caregiver
	if err := h.DB.Order("id desc").Find(&caregivers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, caregivers)
}

func (h CaregiverHandler) Update(c *gin.Context) {
	userIDStr := c.Param("id")

	var userID uint
	if _, err := fmt.Sscan(userIDStr, &userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var req models.UpdateUserNameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := h.DB.Model(&models.User{}).
		Where("id = ?", userID).
		Update("name", req.Name)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   userID,
		"name": req.Name,
	})
}
