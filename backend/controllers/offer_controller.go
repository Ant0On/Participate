package controllers

import (
	"math"
	"net/http"
	"strconv"

	"backend/models"
	"backend/models/DTO"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetOffers(c *gin.Context) {
	var offersWithLocation []DTO.OfferWithLocation
	var result *gorm.DB
	offerType := c.Query("type")

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit := 10
	offset := (page - 1) * limit

	query := models.DB.Model(&models.Offer{})

	var totalRecords int64
	query.Count(&totalRecords)
	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))

	if offerType != "" {
		query = query.Where("offer_type = ?", offerType)
	}

	result = query.
		Joins("JOIN town ON offer.town_id = town.id").
		Joins("JOIN country ON town.country_id = country.id").
		Select("offer.id as offer_id, offer.name, offer.description, offer.price, offer.max_people, offer.is_animal_friendly," +
			"offer.is_recommended, offer.offer_type, town.name as town_name, country.name as country_name").
		Offset(offset).Limit(limit).
		Find(&offersWithLocation)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "offers fetched successfully",
		"data":         offersWithLocation,
		"page":         page,
		"limit":        limit,
		"totalPages":   totalPages,
		"totalRecords": totalRecords,
	})
}

func GetOfferByID(c *gin.Context) {
	offerID := c.Param("id")

	var offerWithLocation DTO.OfferWithLocation
	result := models.DB.
		Model(&models.Offer{}).
		Joins("JOIN town ON offer.town_id = town.id").
		Joins("JOIN country ON town.country_id = country.id").
		Where("offer.id = ?", offerID).
		Select("offer.id as offer_id, offer.name, offer.description, offer.price, offer.max_people, offer.is_animal_friendly," +
			"offer.is_recommended, offer.offer_type, offer.host_id, town.name as town_name, country.name as country_name").
		Find(&offerWithLocation)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Offer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "offer fetched successfully", "data": offerWithLocation})
}

func CreateOffer(c *gin.Context) {
	var offer models.Offer

	if err := c.ShouldBind(&offer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error with createOffer": err.Error()})
		return
	}

	if err := offer.Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"offer.CreateOffer.Save error": err.Error()})
		return
	}

	if err := offer.HandleOfferImageUploads(c, offer.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"image upload error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "offer created successfully!", "offer": offer})
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
	c.JSON(http.StatusOK, gin.H{"message": "offer deleted", "data": offer})
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

	c.JSON(http.StatusOK, gin.H{"message": "Offer updated successfully", "offer": offer})
}

func GetRecommendedOffers(c *gin.Context) {
	var recommendedOffers []DTO.OfferWithLocation
	var result *gorm.DB

	query := models.DB.Model(&models.Offer{})

	result = query.
		Joins("JOIN town ON offer.town_id = town.id").
		Joins("JOIN country ON town.country_id = country.id").
		Where("is_recommended = ?", true).
		Select("offer.id as offer_id, offer.name, offer.description, offer.images_path, offer.price, offer.max_people, offer.is_animal_friendly," +
			"offer.is_recommended, offer.offer_type, town.name as town_name, country.name as country_name").
		Find(&recommendedOffers)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "offers fetched successfully", "data": recommendedOffers})
}
