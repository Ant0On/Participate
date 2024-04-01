package controllers

import (
	"backend/models"
	"backend/models/DTO"

	"github.com/gin-gonic/gin"
)

func CreateEventOffer(c *gin.Context) {
	var offer models.Event
	CreateOffer(c, &offer)
}

func GetEvents(c *gin.Context) {
	var eventsWithLocation []DTO.EventWithLocation
	selectQuery := "accommodation.id as offer_id, accommodation.title, accommodation.description, " +
		"accommodation.price_per_day, accommodation.capacity, accommodation.is_animal_friendly," +
		"accommodation.is_recommended, accommodation.rating, accommodation.type, accommodation.discount, " +
		"accommodation.user_id, town.name as town_name, country.name as country_name"
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
		"event.price, event.capacity, event.is_recommended, event.type, event.discount, " +
		"event.user_id, town.name as town_name, country.name as country_name"
	GetOfferByID(c, OfferQueryParameters{
		tableName:   "event",
		model:       &models.Event{},
		dto:         &eventsWithLocation,
		selectQuery: selectQuery,
	})
}

func GetEventsForHost(c *gin.Context) {
	var eventsWithLocation DTO.EventWithLocation
	selectQuery := "event.id as offer_id, event.title, event.description, " +
		"event.price, event.capacity, event.is_recommended, event.type, event.discount, " +
		"event.user_id, town.name as town_name, country.name as country_name"
	GetOfferByID(c, OfferQueryParameters{
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

/*
func GetRecommendedEvent(c *gin.Context) {
	var recommendedOffers []DTO.EventWithLocation
	var result *gorm.DB

	query := models.DB.Model(&models.Event{})

	result = query.
		Model(&models.Event{}).
		Joins("JOIN town ON offer.town_id = town.id").
		Joins("JOIN country ON town.country_id = country.id").
		Where("is_recommended = ?", true).
		Select("Event.id as offer_id, Event.title, Event.description, " +
			"Event.price_per_day, Event.capacity, Event.is_animal_friendly," +
			"Event.is_recommended, Event.rating, Event.type, Event.discount, " +
			"Event.user_id, town.name as town_name, country.name as country_name").
		Find(&recommendedOffers)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "offers fetched successfully", "data": recommendedOffers})
}
*/

/*
func AddRecommendedOffers(c *gin.Context) {
	offers, err := models.AddRecommendedOffers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for _, offer := range offers {
		offer.IsRecommended = true
		if err := offer.Update(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": offers})
}
*/
