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
	accommodationID := c.Param("id")

	var inputAccommodation models.Accommodation

	if err := c.ShouldBindJSON(&inputAccommodation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingAccommodation models.Accommodation
	if err := models.DB.Preload("Rooms.RoomFacilities").
		Preload("GeneralFacilities").
		First(&existingAccommodation, accommodationID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Accommodation not found"})
		return
	}

	tx := models.DB.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
		return
	}

	var town models.Town
	if err := models.DB.Where("name = ? AND country_id = ?", inputAccommodation.Town.Name, inputAccommodation.Town.CountryID).First(&town).Error; err != nil {
		town = inputAccommodation.Town
		if err := models.DB.FirstOrCreate(&town, models.Town{Name: town.Name, CountryID: town.CountryID}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create or find town"})
			return
		}
	}

	inputAccommodation.Town = town
	inputAccommodation.TownID = town.ID

	if err := tx.Model(&existingAccommodation).Updates(inputAccommodation).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update activity"})
		return
	}

	if err := tx.Model(&existingAccommodation).Association("GeneralFacilities").Clear(); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear existing general facilities"})
		return
	}

	for _, facility := range inputAccommodation.GeneralFacilities {
		var existingFacility models.GeneralFacility
		if err := tx.Where("name = ?", facility.Name).First(&existingFacility).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := tx.Model(&existingAccommodation).Association("GeneralFacilities").Append(&existingFacility); err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	for _, room := range inputAccommodation.Rooms {
		var existingRoom models.Room
		if err := tx.Where("id = ?", room.ID).First(&existingRoom).Error; err != nil {
			if err := tx.Create(&room).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create room"})
				return
			}
			if err := tx.Where("id = ?", room.ID).First(&existingRoom).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to re-fetch newly created room"})
				return
			}
		} else {
			if err := tx.Model(&existingRoom).Updates(room).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update room"})
				return
			}
		}
		if err := tx.Model(&existingRoom).Association("RoomFacilities").Clear(); err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear existing room facilities"})
			return
		}
		for _, facility := range room.RoomFacilities {
			var existingFacility models.RoomFacility
			if err := tx.Where("name = ?", facility.Name).First(&existingFacility).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := tx.Model(&existingRoom).Association("RoomFacilities").Append(&existingFacility); err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Accommodation updated successfully!", "accommodation": existingAccommodation})
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
