package controllers

import (
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
)

func CreateEventOffer(c *gin.Context) {
	var offer models.Event
	CreateOffer(c, "event", &offer)
}

func GetEvents(c *gin.Context) {
	params := OfferQueryParameters{
		Model:    &[]models.Event{},
		Preloads: []string{"Town.Country"},
		Filters:  map[string]interface{}{},
	}
	FetchOffers(c, params)
}

func GetEventByID(c *gin.Context) {
	offerID := c.Param("id")
	params := OfferQueryParameters{
		Model:    &[]models.Event{},
		Preloads: []string{"Town.Country"},
		Filters:  map[string]interface{}{"id": offerID},
	}
	FetchOffers(c, params)
}

func GetEventsForHost(c *gin.Context) {
	hostID := c.Param("id")
	params := OfferQueryParameters{
		Model:    &[]models.Event{},
		Preloads: []string{"Town.Country"},
		Filters:  map[string]interface{}{"user_id": hostID},
	}
	FetchOffers(c, params)
}

func DeleteEvent(c *gin.Context) {
	DeleteOffer(c, models.GetEventByID)
}

func UpdateEvent(c *gin.Context) {
	eventID := c.Param("id")

	var inputEvent models.Event

	if err := c.ShouldBindJSON(&inputEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingEvent models.Event
	if err := models.DB.First(&existingEvent, eventID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	var town models.Town
	if err := models.DB.Where("name = ? AND country_id = ?", inputEvent.Town.Name, inputEvent.Town.CountryID).First(&town).Error; err != nil {
		town = inputEvent.Town
		if err := models.DB.FirstOrCreate(&town, models.Town{Name: town.Name, CountryID: town.CountryID}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create or find town"})
			return
		}
	}

	inputEvent.Town = town
	inputEvent.TownID = town.ID

	if err := models.DB.Model(&existingEvent).Updates(inputEvent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!", "event": existingEvent})
}
