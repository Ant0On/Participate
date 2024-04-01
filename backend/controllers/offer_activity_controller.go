package controllers

import (
	"net/http"

	"backend/models"
	"backend/models/DTO"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

func CreateActivityOffer(c *gin.Context) {
	var offer models.Activity
	CreateOffer(c, &offer)
}

func GetActivities(c *gin.Context) {
	var activityWithLocation []DTO.ActivityWithLocation
	selectQuery := "activity.id as offer_id, activity.title, activity.description, " +
		"activity.price, activity.capacity, activity.skill_level," +
		"activity.is_recommended, activity.duration, activity.type, activity.discount, " +
		"activity.user_id, town.name as town_name, country.name as country_name"
	GetOffers(c, OfferQueryParameters{
		tableName:   "activity",
		model:       &models.Activity{},
		dto:         &activityWithLocation,
		selectQuery: selectQuery,
	})
}

func GetActivityByID(c *gin.Context) {
	var activityWithLocation DTO.ActivityWithLocation
	selectQuery := "activity.id as offer_id, activity.title, activity.description, " +
		"activity.price, activity.capacity, activity.skill_level," +
		"activity.is_recommended, activity.duration, activity.type, activity.discount, " +
		"activity.user_id, town.name as town_name, country.name as country_name"
	GetOfferByID(c, OfferQueryParameters{
		tableName:   "activity",
		model:       &models.Activity{},
		dto:         &activityWithLocation,
		selectQuery: selectQuery,
	})
}

func GetActivitiesForHost(c *gin.Context) {
	var activityWithLocation DTO.ActivityWithLocation
	selectQuery := "activity.id as offer_id, activity.title, activity.description, " +
		"activity.price, activity.capacity, activity.skill_level," +
		"activity.is_recommended, activity.duration, activity.type, activity.discount, " +
		"activity.user_id, town.name as town_name, country.name as country_name"
	GetOffersForHost(c, OfferQueryParameters{
		tableName:   "activity",
		model:       &models.Activity{},
		dto:         &activityWithLocation,
		selectQuery: selectQuery,
	})
}

func DeleteActivity(c *gin.Context) {
	DeleteOffer(c, models.GetActivityByID)
}

func UpdateActivity(c *gin.Context) {
	UpdateOffer(c, models.GetActivityByID)
}

func DiscountActivity(c *gin.Context) {
	offerID := c.Param("id")

	var offer models.Activity
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

func ChangeActivityPrice(c *gin.Context) {
	offerID := c.Param("id")
	var req utils.ChangePriceReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	offer, err := models.GetActivityByID(offerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if offer == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "offer not found"})
		return
	}

	if a, ok := offer.(*models.Activity); ok {
		if err := a.UpdatePrice(req.Price); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": offer})
}

/*
func GetRecommendedActivity(c *gin.Context) {
	var recommendedOffers []DTO.ActivityWithLocation
	var result *gorm.DB

	query := models.DB.Model(&models.Activity{})

	result = query.
		Model(&models.Activity{}).
		Joins("JOIN town ON offer.town_id = town.id").
		Joins("JOIN country ON town.country_id = country.id").
		Where("is_recommended = ?", true).
		Select("Activity.id as offer_id, Activity.title, Activity.description, " +
			"Activity.price_per_day, Activity.capacity, Activity.is_animal_friendly," +
			"Activity.is_recommended, Activity.rating, Activity.type, Activity.discount, " +
			"Activity.user_id, town.name as town_name, country.name as country_name").
		Find(&recommendedOffers)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "offers fetched successfully", "data": recommendedOffers})
}
*/
