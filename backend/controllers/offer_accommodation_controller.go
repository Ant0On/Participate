package controllers

import (
	"math"
	"net/http"
	"strconv"

	"backend/models"
	"backend/models/DTO"
	"backend/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DiscountRequest struct {
	Discount float64 `json:"discount" binding:"required,gte=0,lte=100"`
}

func CreateAccommodationOffer(c *gin.Context) {
	var offer models.Accommodation

	if err := c.ShouldBind(&offer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"c.ShouldBind: ": err.Error()})
		return
	}

	if err := offer.Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"offer.Save: ": err.Error()})
		return
	}
	if err := offer.HandleOfferImageUploads(c, offer.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"offer.HandleOfferImageUploads: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "offer created successfully!", "offer": offer})
}

func GetAccommodations(c *gin.Context) {
	var accommodationWithLocation []DTO.AccommodationWithLocation
	var result *gorm.DB

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit := 10
	offset := (page - 1) * limit

	query := models.DB.Model(&models.Accommodation{})

	var totalRecords int64
	query.Count(&totalRecords)
	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))

	result = query.
		Joins("JOIN town ON offer.town_id = town.id").
		Joins("JOIN country ON town.country_id = country.id").
		Select("accommodation.id as offer_id, accommodation.title, accommodation.description, " +
			"accommodation.price_per_day, accommodation.capacity, accommodation.is_animal_friendly," +
			"accommodation.is_recommended, accommodation.rating, accommodation.type, accommodation.discount, " +
			"accommodation.user_id, town.name as town_name, country.name as country_name").
		Offset(offset).Limit(limit).
		Find(&accommodationWithLocation)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "offers fetched successfully",
		"data":         accommodationWithLocation,
		"page":         page,
		"limit":        limit,
		"totalPages":   totalPages,
		"totalRecords": totalRecords,
	})
}

func GetAccommodationByID(c *gin.Context) {
	offerID := c.Param("id")

	var accommodationWithLocation DTO.AccommodationWithLocation

	result := models.DB.
		Model(&models.Accommodation{}).
		Joins("JOIN town ON offer.town_id = town.id").
		Joins("JOIN country ON town.country_id = country.id").
		Where("accommodation.id = ?", offerID).
		Select("accommodation.id as offer_id, accommodation.title, accommodation.description, " +
			"accommodation.price_per_day, accommodation.capacity, accommodation.is_animal_friendly," +
			"accommodation.is_recommended, accommodation.rating, accommodation.type, accommodation.discount, " +
			"accommodation.user_id, town.name as town_name, country.name as country_name").
		Find(&accommodationWithLocation)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result.Error: ": err.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Offer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "offer fetched successfully", "data": accommodationWithLocation})
}

func DeleteAccommodation(c *gin.Context) {
	id := c.Params.ByName("id")

	offer, err := models.GetAccommodationByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"models.GetAccommodationByID: ": err.Error()})
		return
	}

	if err = offer.Delete(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"offer.Delete: ": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "offer deleted", "data": offer})
}

func UpdateAccommodation(c *gin.Context) {
	id := c.Params.ByName("id")

	offer, err := models.GetAccommodationByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"models.GetAccommodationByID: ": err.Error()})
		return
	}

	if err = c.ShouldBindJSON(&offer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"c.ShouldBindJSON: ": err.Error()})
		return
	}

	if err = offer.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"offer.Update: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Offer updated successfully", "offer": offer})
}

func DiscountAccommodation(c *gin.Context) {
	offerID := c.Param("id")

	var offer models.Accommodation
	var discountReq DiscountRequest

	if err := models.DB.First(&offer, offerID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Offer not found"})
		return
	}

	if err := c.ShouldBindJSON(&discountReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	offer.Discount = discountReq.Discount

	if err := offer.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"offer.Update: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Discount assigned successfully"})
}

func GetAccommodationForHost(c *gin.Context) {
	hostID := c.Param("id")

	var accommodationWithLocation []DTO.AccommodationWithLocation
	result := models.DB.
		Model(&models.Offer{}).
		Joins("JOIN town ON offer.town_id = town.id").
		Joins("JOIN country ON town.country_id = country.id").
		Where("offer.user_id = ?", hostID).
		Select("accommodation.id as offer_id, accommodation.title, accommodation.description, " +
			"accommodation.price_per_day, accommodation.capacity, accommodation.is_animal_friendly," +
			"accommodation.is_recommended, accommodation.rating, accommodation.type, accommodation.discount, " +
			"accommodation.user_id, town.name as town_name, country.name as country_name").
		Find(&accommodationWithLocation)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Offer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "offer fetched successfully", "data": accommodationWithLocation})

}

type ChangeAccommodationPriceReq struct {
	PricePerDay float64 `json:"price_per_day" binding:"required,gt=0"`
}

func ChangeAccommodationPrice(c *gin.Context) {
	offerId := c.Param("id")
	var req utils.ChangePriceReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	offer, err := models.GetAccommodationByID(offerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if offer == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "offer not found"})
		return
	}

	offer.PricePerDay = req.Price

	if err := offer.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": offer})
}

/*
func GetRecommendedAccommodation(c *gin.Context) {
	var recommendedOffers []DTO.AccommodationWithLocation
	var result *gorm.DB

	query := models.DB.Model(&models.Accommodation{})

	result = query.
		Model(&models.Accommodation{}).
		Joins("JOIN town ON offer.town_id = town.id").
		Joins("JOIN country ON town.country_id = country.id").
		Where("is_recommended = ?", true).
		Select("accommodation.id as offer_id, accommodation.title, accommodation.description, " +
			"accommodation.price_per_day, accommodation.capacity, accommodation.is_animal_friendly," +
			"accommodation.is_recommended, accommodation.rating, accommodation.type, accommodation.discount, " +
			"accommodation.user_id, town.name as town_name, country.name as country_name").
		Find(&recommendedOffers)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "offers fetched successfully", "data": recommendedOffers})
}
*/
