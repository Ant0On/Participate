package controllers

import (
	"net/http"

	"backend/models"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

func CreateAccommodationOffer(c *gin.Context) {
	var offer models.Accommodation
	CreateOffer(c, "accommodation", &offer)
}

func GetAccommodations(c *gin.Context) {
	params := OfferQueryParameters{
		Model:    &[]models.Accommodation{},
		Preloads: []string{"Rooms.RoomFacilities", "Town.Country", "GeneralFacilities"},
		Filters:  map[string]interface{}{},
	}
	FetchOffers(c, params)
}

func GetAccommodationByID(c *gin.Context) {
	offerID := c.Param("id")
	params := OfferQueryParameters{
		Model:    &[]models.Accommodation{},
		Preloads: []string{"Rooms.RoomFacilities", "Town.Country", "GeneralFacilities"},
		Filters:  map[string]interface{}{"id": offerID},
	}
	FetchOffers(c, params)
}

func GetAccommodationsForHost(c *gin.Context) {
	hostID := c.Param("id")
	params := OfferQueryParameters{
		Model:    &[]models.Accommodation{},
		Preloads: []string{"Rooms.RoomFacilities", "Town.Country", "GeneralFacilities"},
		Filters:  map[string]interface{}{"user_id": hostID},
	}
	FetchOffers(c, params)
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
