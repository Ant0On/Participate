package controllers

import (
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetOffers(c *gin.Context) {
	var offers []models.Offer
	var result *gorm.DB
	offerType := c.Query("type")
	if offerType != "" {
		result = models.DB.Where("type = ?", offerType).Find(&offers)
	} else {
		result = models.DB.Find(&offers)
	}

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"models.GetOffers error": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": result})
}

func CreateOffer(c *gin.Context) {
	var offer *models.Offer

	if err := c.ShouldBindJSON(&offer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error with createOfferInput": err.Error()})
		return
	}

	if err := offer.Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"customer.SaveCustomer error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "registration success!"})
}
