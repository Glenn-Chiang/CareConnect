package handlers

import (
	"hack4good/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RecipientHandler struct {
	DB *gorm.DB
}

func (h RecipientHandler) List(c *gin.Context) {
	var recipients []models.Recipient
	if err := h.DB.Preload("User").Order("id desc").Find(&recipients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, recipients)
}

func (h RecipientHandler) ListRecipientsByCaregiver(c *gin.Context) {
	caregiverID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid caregiver id"})
		return
	}

	var recipients []models.Recipient

	err = h.DB.
		Preload("User").
		Table("recipients").
		Joins("JOIN caregiver_recipients cr ON cr.recipient_id = recipients.id").
		Where("cr.caregiver_id = ?", uint(caregiverID)).
		Order("recipients.id asc").
		Find(&recipients).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, recipients)
}

func (h RecipientHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	var recipient models.RecipientReturned
	if err := h.DB.
		Table("recipients").
		Select(`
			recipients.id,
			recipients.user_id,
			users.name,
			recipients.age,
			recipients.condition,
			recipients.likes,
			recipients.dislikes,
			recipients.phobias,
			recipients.pet_peeves
		`).
		Joins("JOIN users ON users.id = recipients.user_id").
		Where("recipients.id = ?", id).
		Scan(&recipient).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipient not found"})
		return
	}

	c.JSON(http.StatusOK, recipient)
}

func (h RecipientHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var recipient models.Recipient
	if err := h.DB.First(&recipient, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipient not found"})
		return
	}

	var req models.RecipientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.DB.Transaction(func(tx *gorm.DB) error {
		// Update recipient fields
		if req.Age != nil {
			recipient.Age = req.Age
		}
		if req.Condition != nil {
			recipient.Condition = req.Condition
		}
		if req.Likes != nil {
			recipient.Likes = req.Likes
		}
		if req.Dislikes != nil {
			recipient.Dislikes = req.Dislikes
		}
		if req.Phobias != nil {
			recipient.Phobias = req.Phobias
		}
		if req.PetPeeves != nil {
			recipient.PetPeeves = req.PetPeeves
		}

		if err := tx.Save(&recipient).Error; err != nil {
			return err
		}

		// Update user name (if provided)
		if req.Name != nil {
			if err := tx.Model(&models.User{}).
				Where("id = ?", recipient.UserID).
				Update("name", *req.Name).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Reload with user
	h.DB.Preload("User").First(&recipient, recipient.ID)

	c.JSON(http.StatusOK, recipient)
}
