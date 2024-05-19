package controllers

import (
	"net/http"

	"backend/models"
	"backend/models/DTO"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

func CreateAccommodationOffer(c *gin.Context) {
	var offer models.Accommodation
	CreateOffer(c, "accommodation", &offer)
}

func GetAccommodations(c *gin.Context) {
	var accommodationWithLocation []DTO.AccommodationWithLocation
	selectQuery := "accommodation.id as offer_id, accommodation.title, accommodation.description, " +
		"accommodation.price_per_day, accommodation.capacity, accommodation.is_animal_friendly," +
		"accommodation.is_recommended, accommodation.accommodation_type as type, accommodation.discount, " +
		"accommodation.user_id, town.name as town_name, country.name as country_name"
	GetOffers(c, OfferQueryParameters{
		tableName:   "accommodation",
		model:       &models.Accommodation{},
		dto:         &accommodationWithLocation,
		selectQuery: selectQuery,
	})
}

func GetAccommodationByID(c *gin.Context) {
	var accommodationWithLocation DTO.AccommodationWithLocation
	selectQuery := "accommodation.id as offer_id, accommodation.title, accommodation.description, " +
		"accommodation.price_per_day, accommodation.capacity, accommodation.is_animal_friendly," +
		"accommodation.is_recommended, accommodation.accommodation_type as type, accommodation.discount, " +
		"accommodation.user_id, town.name as town_name, country.name as country_name"
	GetOfferByID(c, OfferQueryParameters{
		tableName:   "accommodation",
		model:       &models.Accommodation{},
		dto:         &accommodationWithLocation,
		selectQuery: selectQuery,
	})
}

func GetAccommodationsForHost(c *gin.Context) {
	var accommodationWithLocation DTO.AccommodationWithLocation
	selectQuery := "accommodation.id as offer_id, accommodation.title, accommodation.description, " +
		"accommodation.price_per_day, accommodation.capacity, accommodation.is_animal_friendly," +
		"accommodation.is_recommended, accommodation.accommodation_type as type, accommodation.discount, " +
		"accommodation.user_id, town.name as town_name, country.name as country_name"
	GetOffersForHost(c, OfferQueryParameters{
		tableName:   "accommodation",
		model:       &models.Accommodation{},
		dto:         &accommodationWithLocation,
		selectQuery: selectQuery,
	})
}

func DeleteAccommodation(c *gin.Context) {
	DeleteOffer(c, models.GetAccommodationByID)
}

func UpdateAccommodation(c *gin.Context) {
	UpdateOffer(c, models.GetAccommodationByID)
}

func DiscountAccommodation(c *gin.Context) {
	DiscountOffer(c, models.GetAccommodationByID)
}

func ChangeAccommodationPrice(c *gin.Context) {
	ChangeOfferPrice(c, models.GetAccommodationByID)
}

func AddGeneralFacilities(c *gin.Context) {
	id := c.Param("id")
	var req utils.FacilitiesRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accommodation, err := models.GetAccommodationById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var facilities []models.GeneralFacility

	for i := 0; i < len(req.Facilities); i++ {
		facility, err := models.GetFacilityByName(req.Facilities[i])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"models.GetFacilityByName": err.Error()})
			return
		}
		facilities = append(facilities, facility)
	}
	err = accommodation.AddFacilities(facilities)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"accommodation.UpdateFacilities: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Facilities added to accommodation successful", "data": id})

}

func AddRoomFacilities(c *gin.Context) {
	id := c.Param("id")
	var req utils.FacilitiesRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room, err := models.GetRoomByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var facilities []models.RoomFacility

	for i := 0; i < len(req.Facilities); i++ {
		facility, err := models.GetRoomFacilityByName(req.Facilities[i])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"models.GetRoomFacilityByName": err.Error()})
			return
		}
		facilities = append(facilities, facility)
	}
	err = room.AddFacilities(facilities)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"room.AddFacilities: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Facilities added to room successful", "data": id})

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
