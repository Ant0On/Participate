package controllers

import (
	"backend/models"
	"backend/models/DTO"

	"github.com/gin-gonic/gin"
)

func CreateEventOffer(c *gin.Context) {
	var offer models.Event
	CreateOffer(c, "event", &offer)
}

func GetEvents(c *gin.Context) {
	var eventsWithLocation []DTO.EventWithLocation
	selectQuery := "event.id as offer_id, event.title, event.description, " +
		"event.price, event.capacity, event.event_type as type, event.discount, " +
		"event.user_id, town.name as town_name, country.name as country_name"
	GetOffers(c, OfferQueryParameters{
		tableName:   "event",
		model:       &models.Event{},
		dto:         &eventsWithLocation,
		selectQuery: selectQuery,
	})
}

func GetEventByID(c *gin.Context) {
	var eventsWithLocation DTO.EventWithLocation
	selectQuery := "event.id as offer_id, event.title, event.description, " +
		"event.price, event.capacity, event.event_type as type, event.discount, " +
		"event.user_id, town.name as town_name, country.name as country_name"
	GetOfferByID(c, OfferQueryParameters{
		tableName:   "event",
		model:       &models.Event{},
		dto:         &eventsWithLocation,
		selectQuery: selectQuery,
	})
}

func GetEventsForHost(c *gin.Context) {
	var eventsWithLocation []DTO.EventWithLocation
	selectQuery := "event.id as offer_id, event.title, event.description, " +
		"event.price, event.capacity, event.event_type as type, event.discount, " +
		"event.user_id, town.name as town_name, country.name as country_name"
	GetOffersForHost(c, OfferQueryParameters{
		tableName:   "event",
		model:       &models.Event{},
		dto:         &eventsWithLocation,
		selectQuery: selectQuery,
	})
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
