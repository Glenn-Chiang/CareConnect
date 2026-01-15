package handlers

import (
	"hack4good/internal/models"
	"net/http"
	"strconv"

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
 
func (h CaregiverHandler) ListByRecipient(c *gin.Context) {
	recipientID64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid recipient id"})
		return
	}
	recipientID := uint(recipientID64)

	var caregivers []models.Caregiver
	err = h.DB.
		Preload("User").
		Table("caregivers").
		Joins("JOIN caregiver_recipients cr ON cr.caregiver_id = caregivers.id").
		Where("cr.recipient_id = ?", recipientID).
		Order("caregivers.id asc").
		Find(&caregivers).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, caregivers)
}

func (h CaregiverHandler) GetByUserID(c *gin.Context) {
	userID64, err := strconv.ParseUint(c.Param("userId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	userID := uint(userID64)

	var caregiver models.Caregiver
	if err := h.DB.Preload("User").First(&caregiver, "user_id = ?", userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "caregiver not found"})
		return
	}

	c.JSON(http.StatusOK, caregiver)
}
