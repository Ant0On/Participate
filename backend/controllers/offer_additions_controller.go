package controllers

import (
	"net/http"

	"backend/models"
	"backend/models/DTO"

	"github.com/gin-gonic/gin"
)

func GetAccommodationWithRoomsByID(c *gin.Context) {
	accommodationID := c.Param("id")

	var accommodationRooms []DTO.AccommodationRoom

	result := models.DB.
		Table("accommodation").
		Joins("JOIN room ON accommodation.id = room.accommodation_id").
		Joins("LEFT JOIN room_room_facilities ON room.id = room_room_facilities.room_id").
		Joins("LEFT JOIN room_facility ON room_room_facilities.room_facility_id = room_facility.id").
		Joins("LEFT JOIN accommodation_general_facilities ON accommodation.id = accommodation_general_facilities.accommodation_id").
		Joins("LEFT JOIN general_facility ON accommodation_general_facilities.general_facility_id = general_facility.id").
		Where("accommodation.id = ?", accommodationID).
		Select(`
			accommodation.title as title,
			accommodation.price_per_day as price_per_day,
			room.room_name,
			room.room_description,
			room.capacity as room_capacity,
			room.area as room_area,
			accommodation.accommodation_type as type,
			accommodation.is_animal_friendly,
			accommodation.town_id,
			accommodation.user_id,
			STRING_AGG(room_facility.name, ', ') as room_facilities,
			STRING_AGG(general_facility.name, ', ') as general_facilities
		`).
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
