package controllers

import (
	"fmt"
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

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result.Scan(&offers)

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": offers})
}

func CreateOffer(c *gin.Context) {
	var offer *models.Offer

	if err := c.ShouldBindJSON(&offer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error with createOffer": err.Error()})
		return
	}

	if err := offer.Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"offer.CreateOffer.Save error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "offer created successfully!"})
}

func DeleteOffer(c *gin.Context) {
	id := c.Params.ByName("id")

	var offer *models.Offer
	var err error

	if offer, err = models.GetOfferByID(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"offer.DeleteOffer.First error": err.Error()})
		return
	}

	if err = offer.Delete(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"offer.DeleteOffer.Delete error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("offer deleted, id: %s", id)})
}

func UpdateOffer(c *gin.Context) {
	id := c.Params.ByName("id")

	var offer *models.Offer
	var err error

	if offer, err = models.GetOfferByID(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"offer.UpdateOffer.First error": err.Error()})
		return
	}

	if err = c.ShouldBindJSON(&offer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error with UpdateOffer": err.Error()})
		return
	}

	if err = offer.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error with UpdateOffer.Update": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Offer updated successfully"})
}
