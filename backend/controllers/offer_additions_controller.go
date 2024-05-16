package controllers

import (
	"net/http"

	"backend/models"
	"backend/models/DTO"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAccommodationRoomsByID(c *gin.Context) {
	accommodationID := c.Param("id")

	var accommodationRooms []DTO.AccommodationRoom

	result := models.DB.
		Model(&models.Room{}).
		Joins("JOIN accommodation ON room.accommodation_id = accommodation.id").
		Joins("LEFT JOIN room_room_facilities ON room.id = room_room_facilities.room_id").
		Joins("LEFT JOIN room_facility ON room_room_facilities.room_facility_id = room_facility.id").
		Joins("LEFT JOIN accommodation_general_facilities ON accommodation.id = accommodation_general_facilities.accommodation_id").
		Joins("LEFT JOIN general_facility ON accommodation_general_facilities.general_facility_id = general_facility.id").
		Where("room.accommodation_id = ?", accommodationID).
		Select("accommodation.title as title, accommodation.price_per_day as price_per_day, room.room_name, room.room_description, room.capacity, STRING_AGG(room_facility.name, ', ') as room_facilities, STRING_AGG(general_facility.name, ', ') as general_facilities").
		Group("accommodation.id, room.id").
		Scan(&accommodationRooms)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Rooms not found for the accommodation"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Rooms fetched successfully", "data": accommodationRooms})
}

func GetActivitiesWithEquipment(c *gin.Context) {
	id := c.Param("id")

	var activityWithEquipment []DTO.ActivityEquipment

	var result *gorm.DB

	result = models.DB.
		Table("activity").
		Select("activity.title as title, activity.description, activity.price, STRING_AGG(equipment.name, ', ') as equipment").
		Joins("JOIN activity_equipment ON activity.id = activity_equipment.activity_id").
		Joins("JOIN equipment ON equipment.id = activity_equipment.equipment_id").
		Where("activity.id = ?", id).
		Group("activity.id").
		Scan(&activityWithEquipment)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Equipment not found for the activity"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Activities with equipment fetched successfully", "data": activityWithEquipment})
}
