package controllers

import (
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
)

func RecommendOffer(c *gin.Context) {
	offerID := c.Param("id")

	var offer models.Offer

	if err := models.DB.First(&offer, offerID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Offer not found"})
		return
	}

	offer.IsRecommended = true

	if err := models.DB.Save(&offer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update offer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Offer changed to recommended", "offer": offer})
}
