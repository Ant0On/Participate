package controllers

import (
	"net/http"
	"strconv"

	"backend/models"
	"backend/models/DTO"
	"backend/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateAccommodationOffer(c *gin.Context) {
	var offer models.Accommodation
	CreateOffer(c, "accommodation", &offer)
}

func GetAccommodations(c *gin.Context) {
	var accommodations []DTO.AccommodationDTO
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit := 10
	offset := (page - 1) * limit

	query := models.DB.Table("accommodation").
		Select(`accommodation.title, accommodation.description, accommodation.capacity, 
		accommodation.discount, accommodation.town_id, accommodation.user_id, town.name as town_name, 
		country.name as country_name`).
		Joins("JOIN town ON accommodation.town_id = town.id").
		Joins("JOIN country ON town.country_id = country.id").
		Preload("GeneralFacilities", func(db *gorm.DB) *gorm.DB {
			return db.Model(&models.GeneralFacility{}).Select("general_facility.*")
		}).
		Preload("Rooms", func(db *gorm.DB) *gorm.DB {
			return db.Model(&models.Room{}).Select("room.*").
				Joins("JOIN accommodation ON room.accommodation_id = accommodation.id").
				Joins("JOIN room_room_facilities ON room.id = room_room_facilities.room_id").
				Joins("JOIN room_facility ON room_room_facilities.room_facility_id = room_facility.id").
				Where("accommodation.id = accommodation.id").
				Preload("RoomFacilities")
		}).
		Offset(offset).
		Limit(limit)

	if err := query.Scan(&accommodations).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "accommodations fetched successfully",
		"data":    accommodations,
		"page":    page,
		"limit":   limit,
	})
}

func GetAccommodationByID(c *gin.Context) {
	var accommodationWithLocation DTO.AccommodationDTO
	selectQuery := "accommodation.id as offer_id, accommodation.title, accommodation.description, " +
		"accommodation.price_per_day, accommodation.capacity, accommodation.is_animal_friendly," +
		"accommodation.accommodation_type as type, accommodation.discount, " +
		"accommodation.user_id, town.name as town_name, country.name as country_name"
	GetOfferByID(c, OfferQueryParameters{
		tableName:   "accommodation",
		model:       &models.Accommodation{},
		dto:         &accommodationWithLocation,
		selectQuery: selectQuery,
	})
}

func GetAccommodationsForHost(c *gin.Context) {
	var accommodationWithLocation []DTO.AccommodationDTO
	selectQuery := "accommodation.id as offer_id, accommodation.title, accommodation.description, " +
		"accommodation.price_per_day, accommodation.capacity, accommodation.is_animal_friendly," +
		"accommodation.accommodation_type as type, accommodation.discount, " +
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
