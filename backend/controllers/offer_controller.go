package controllers

import (
	"fmt"
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OfferWithLocation struct {
	OfferID          uint             `json:"offer_id"`
	Name             string           `json:"name"`
	Description      string           `json:"description"`
	Price            float64          `json:"price"`
	MaxPeople        int              `json:"max_people"`
	IsAnimalFriendly bool             `json:"is_animal_friendly"`
	IsRecommended    bool             `json:"is_recommended"`
	OfferType        models.OfferType `json:"offer_type"`
	TownName         string           `json:"town_name"`
	CountryName      string           `json:"country_name"`
}

func GetOffers(c *gin.Context) {
	var offersWithLocation []OfferWithLocation
	var result *gorm.DB
	offerType := c.Query("type")

	query := models.DB.Model(&models.Offer{})

	if offerType != "" {
		query = query.Where("type = ?", offerType)
	}

	result = query.
		Joins("JOIN town ON offer.town_id = town.id").
		Joins("JOIN country ON town.country_id = country.id").
		Select("offer.id as offer_id, offer.name, offer.description, offer.price, offer.max_people, offer.is_animal_friendly," +
			"offer.is_recommended, offer.offer_type, town.name as town_name, country.name as country_name").
		Find(&offersWithLocation)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": query.
		Joins("JOIN town ON offer.town_id = town.id").
		Joins("JOIN country ON town.country_id = country.id").
		Select("offer.id")})
}

func GetOfferByID(c *gin.Context) {
	id := c.Params.ByName("id")

	var offer *models.Offer
	var err error

	if offer, err = models.GetOfferByID(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"models.GetOfferById:": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": offer})
}

func CreateOffer(c *gin.Context) {
	var offer models.Offer

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
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("offer deleted: %v", offer)})
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
