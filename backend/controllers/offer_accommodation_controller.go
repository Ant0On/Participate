package controllers

import (
	"net/http"
	"strconv"

	"backend/models"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

func CreateAccommodationOffer(c *gin.Context) {
	var offer models.Accommodation
	CreateOffer(c, "accommodation", &offer)
}

func fetchAccommodations(c *gin.Context, filters map[string]interface{}) {
	var accommodations []models.Accommodation

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	query := models.DB.Preload("Rooms.RoomFacilities").
		Preload("Town.Country").
		Preload("GeneralFacilities")

	for key, value := range filters {
		query = query.Where(key+" = ?", value)
	}

	query = query.Offset(offset).Limit(pageSize).Find(&accommodations)

	if query.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": query.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "accommodations fetched successfully",
		"data":      accommodations,
		"page":      page,
		"page_size": pageSize,
	})
}

func GetAccommodations(c *gin.Context) {
	fetchAccommodations(c, nil)
}

func GetAccommodationByID(c *gin.Context) {
	offerID := c.Param("id")

	filters := map[string]interface{}{
		"id": offerID,
	}

	fetchAccommodations(c, filters)
}

func GetAccommodationsForHost(c *gin.Context) {
	hostID := c.Param("id")

	filters := map[string]interface{}{
		"user_id": hostID,
	}

	fetchAccommodations(c, filters)
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
