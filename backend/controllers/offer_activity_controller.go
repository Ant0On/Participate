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
	CreateOffer(c, "activity", &offer)
}

func GetActivities(c *gin.Context) {
	var activityWithLocation []DTO.ActivityWithLocation
	selectQuery := "activity.id as offer_id, activity.title, activity.description, " +
		"activity.price, activity.capacity, activity.skill_level," +
		"activity.is_recommended, activity.duration, activity.activity_type, activity.discount, " +
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
		"activity.is_recommended, activity.duration, activity.activity_type, activity.discount, " +
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
		"activity.is_recommended, activity.duration, activity.activity_type, activity.discount, " +
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
	DiscountOffer(c, models.GetActivityByID)
}

func ChangeActivityPrice(c *gin.Context) {
	ChangeOfferPrice(c, models.GetActivityByID)
}

func UpdateEquipment(c *gin.Context) {
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
		equipment, err := models.GetEquipmentByID(req.Equipment[i])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"models.GetEquipmentByID": err.Error()})
			return
		}
		newEquipments = append(newEquipments, equipment)
	}
	err = activity.UpdateEquipment(newEquipments)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"activity.UpdateEquipment: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Equipment added to activity successful"})

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
