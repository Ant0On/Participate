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
	UpdateOffer(c, models.GetActivityByID)
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
