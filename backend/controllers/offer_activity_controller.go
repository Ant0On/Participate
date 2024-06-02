package controllers

import (
	"net/http"

	"backend/models"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

func CreateActivityOffer(c *gin.Context) {
	var offer models.Activity
	CreateOffer(c, "activity", &offer)
}

func GetActivities(c *gin.Context) {
	params := OfferQueryParameters{
		Model:    &[]models.Activity{},
		Preloads: []string{"Town.Country", "Equipment"},
		Filters:  map[string]interface{}{},
	}
	FetchOffers(c, params)
}

func GetActivityByID(c *gin.Context) {
	offerID := c.Param("id")
	params := OfferQueryParameters{
		Model:    &[]models.Activity{},
		Preloads: []string{"Town.Country", "Equipment"},
		Filters:  map[string]interface{}{"id": offerID},
	}
	FetchOffers(c, params)
}

func GetActivitiesForHost(c *gin.Context) {
	hostID := c.Param("id")
	params := OfferQueryParameters{
		Model:    &[]models.Activity{},
		Preloads: []string{"Town.Country", "Equipment"},
		Filters:  map[string]interface{}{"user_id": hostID},
	}
	FetchOffers(c, params)
}

func DeleteActivity(c *gin.Context) {
	DeleteOffer(c, models.GetActivityByID)
}

func UpdateActivity(c *gin.Context) {
	activityID := c.Param("id")

	var inputActivity models.Activity

	if err := c.ShouldBindJSON(&inputActivity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingActivity models.Activity
	if err := models.DB.Preload("Equipment").First(&existingActivity, activityID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		return
	}

	tx := models.DB.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
		return
	}

	var town models.Town
	if err := models.DB.Where("name = ? AND country_id = ?", inputActivity.Town.Name, inputActivity.Town.CountryID).First(&town).Error; err != nil {
		town = inputActivity.Town
		if err := models.DB.FirstOrCreate(&town, models.Town{Name: town.Name, CountryID: town.CountryID}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create or find town"})
			return
		}
	}

	inputActivity.Town = town
	inputActivity.TownID = town.ID

	if err := tx.Model(&existingActivity).Updates(inputActivity).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update activity"})
		return
	}

	if err := tx.Model(&existingActivity).Association("Equipment").Clear(); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear existing equipment"})
		return
	}

	for _, equipment := range inputActivity.Equipment {
		var existingEquipment models.Equipment
		if err := tx.Where("name = ?", equipment.Name).First(&existingEquipment).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := tx.Model(&existingActivity).Association("Equipment").Append(&existingEquipment); err != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Activity updated successfully!", "activity": existingActivity})
}

func DiscountActivity(c *gin.Context) {
	DiscountOffer(c, models.GetActivityByID)
}

func ChangeActivityPrice(c *gin.Context) {
	ChangeOfferPrice(c, models.GetActivityByID)
}

func AddEquipment(c *gin.Context) {
	id := c.Param("id")
	var req utils.EquipmentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	activity, err := models.GetActivityById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newEquipments []models.Equipment

	for i := 0; i < len(req.Equipment); i++ {
		equipment, err := models.GetEquipmentByName(req.Equipment[i])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"models.GetEquipmentByName": err.Error()})
			return
		}
		newEquipments = append(newEquipments, equipment)
	}
	err = activity.AddEquipment(newEquipments)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"activity.UpdateEquipment: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Equipment added to activity successful"})

}
