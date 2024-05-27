package controllers

import (
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
	UpdateOffer(c, models.GetEventByID)
}

func DiscountEvent(c *gin.Context) {
	DiscountOffer(c, models.GetEventByID)
}

func ChangeEventPrice(c *gin.Context) {
	ChangeOfferPrice(c, models.GetEventByID)
}
