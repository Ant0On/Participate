package controllers

import (
	"math"
	"net/http"
	"strconv"

	"backend/models"
	"backend/models/DTO"
	"backend/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateActivityOffer(c *gin.Context) {
	var offer models.Activity
	CreateOffer(c, &offer)
}

func GetActivities(c *gin.Context) {
	var activityWithLocation []DTO.ActivityWithLocation
	var result *gorm.DB

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit := 10
	offset := (page - 1) * limit

	query := models.DB.Model(&models.Activity{})

	var totalRecords int64
	query.Count(&totalRecords)
	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))

	result = query.
		Joins("JOIN town ON offer.town_id = town.id").
		Joins("JOIN country ON town.country_id = country.id").
		Select("activity.id as offer_id, activity.title, activity.description, " +
			"activity.price, activity.capacity, activity.skill_level," +
			"activity.is_recommended, activity.duration, activity.type, activity.discount, " +
			"activity.user_id, town.name as town_name, country.name as country_name").
		Offset(offset).Limit(limit).
		Find(&activityWithLocation)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "offers fetched successfully",
		"data":         activityWithLocation,
		"page":         page,
		"limit":        limit,
		"totalPages":   totalPages,
		"totalRecords": totalRecords,
	})
}

func GetActivityByID(c *gin.Context) {
	offerID := c.Param("id")

	var activityWithLocation DTO.ActivityWithLocation

	result := models.DB.
		Model(&models.Activity{}).
		Joins("JOIN town ON offer.town_id = town.id").
		Joins("JOIN country ON town.country_id = country.id").
		Where("activity.id = ?", offerID).
		Select("activity.id as offer_id, activity.title, activity.description, " +
			"activity.price, activity.capacity, activity.skill_level," +
			"activity.is_recommended, activity.duration, activity.type, activity.discount, " +
			"activity.user_id, town.name as town_name, country.name as country_name").
		Find(&activityWithLocation)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result.Error: ": err.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Offer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "offer fetched successfully", "data": activityWithLocation})
}

func DeleteActivity(c *gin.Context) {
	id := c.Params.ByName("id")

	offer, err := models.GetActivityByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"models.GetActivityByID: ": err.Error()})
		return
	}

	if err = offer.Delete(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"offer.Delete: ": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "offer deleted", "data": offer})
}

func UpdateActivity(c *gin.Context) {
	id := c.Params.ByName("id")

	offer, err := models.GetActivityByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"models.GetActivityByID: ": err.Error()})
		return
	}

	if err = c.ShouldBindJSON(&offer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"c.ShouldBindJSON: ": err.Error()})
		return
	}

	if err = offer.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"offer.Update: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Offer updated successfully", "offer": offer})
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

func GetActivityForHost(c *gin.Context) {
	hostID := c.Param("id")

	var activityWithLocation []DTO.ActivityWithLocation
	result := models.DB.
		Model(&models.Offer{}).
		Joins("JOIN town ON offer.town_id = town.id").
		Joins("JOIN country ON town.country_id = country.id").
		Where("offer.user_id = ?", hostID).
		Select("activity.id as offer_id, activity.title, activity.description, " +
			"activity.price, activity.capacity, activity.skill_level," +
			"activity.is_recommended, activity.duration, activity.type, activity.discount, " +
			"activity.user_id, town.name as town_name, country.name as country_name").
		Find(&activityWithLocation)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Offer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "offer fetched successfully", "data": activityWithLocation})

}

func ChangeActivityPrice(c *gin.Context) {
	offerId := c.Param("id")
	var req utils.ChangePriceReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	offer, err := models.GetActivityByID(offerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if offer == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "offer not found"})
		return
	}

	offer.Price = req.Price

	if err := offer.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
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
