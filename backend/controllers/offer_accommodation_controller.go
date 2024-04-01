package controllers

import (
	"net/http"

	"backend/models"
	"backend/models/DTO"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

type DiscountRequest struct {
	Discount float64 `json:"discount" binding:"required,gte=0,lte=100"`
}

func CreateAccommodationOffer(c *gin.Context) {
	var offer models.Accommodation
	CreateOffer(c, &offer)
}

func GetAccommodations(c *gin.Context) {
	var accommodationWithLocation []DTO.AccommodationWithLocation
	selectQuery := "accommodation.id as offer_id, accommodation.title, accommodation.description, " +
		"accommodation.price_per_day, accommodation.capacity, accommodation.is_animal_friendly," +
		"accommodation.is_recommended, accommodation.rating, accommodation.type, accommodation.discount, " +
		"accommodation.user_id, town.name as town_name, country.name as country_name"
	GetOffers(c, "accommodation", &models.Accommodation{}, &accommodationWithLocation, selectQuery)
}

func GetAccommodationByID(c *gin.Context) {
	var accommodationWithLocation DTO.AccommodationWithLocation
	selectQuery := "accommodation.id as offer_id, accommodation.title, accommodation.description, " +
		"accommodation.price_per_day, accommodation.capacity, accommodation.is_animal_friendly," +
		"accommodation.is_recommended, accommodation.rating, accommodation.type, accommodation.discount, " +
		"accommodation.user_id, town.name as town_name, country.name as country_name"
	GetOfferByID(c, "accommodation", &models.Accommodation{}, &accommodationWithLocation, selectQuery)
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
