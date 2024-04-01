package controllers

import (
	"net/http"

	"backend/models"
	"backend/models/DTO"
	"backend/utils"

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
	offerID := c.Param("id")

	var offer models.Event
	var discountReq DiscountRequest

	if err := models.DB.First(&offer, offerID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Offer not found"})
		return
	}

	if err := c.ShouldBindJSON(&discountReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	offer.Discount = discountReq.Discount

	if err := offer.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"offer.Update: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Discount assigned successfully"})
}

func ChangeEventPrice(c *gin.Context) {
	offerId := c.Param("id")
	var req utils.ChangePriceReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	offer, err := models.GetEventByID(offerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if offer == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "offer not found"})
		return
	}

	if a, ok := offer.(*models.Event); ok {
		if err := a.UpdatePrice(req.Price); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	if err := offer.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": offer})
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
