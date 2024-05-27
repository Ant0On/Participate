package controllers

import (
	"math"
	"net/http"
	"strconv"

	"backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRooms(c *gin.Context) {
	var rooms []models.Room

	if err := c.ShouldBindJSON(&rooms); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"c.ShouldBindJSON: ": err.Error()})
		return
	}

	tx := models.DB.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
		return
	}

	for _, room := range rooms {
		if err := tx.Create(&room).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"room.Save: ": err.Error()})
			return
		}

		for _, facility := range room.RoomFacilities {
			if err := tx.Where("name = ?", facility.Name).First(&facility).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusBadRequest, gin.H{"facility.Save: ": err.Error()})
				return
			}
			if err := tx.Model(&room).Association("RoomFacilities").Append(&facility); err != nil {
				tx.Rollback()
				c.JSON(http.StatusBadRequest, gin.H{"association.Save: ": err.Error()})
				return
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "rooms created successfully!", "rooms": rooms})
}

func GetRoomsForAccommodation(c *gin.Context) {
	var rooms []models.Room
	var result *gorm.DB
	accommodationId := c.Query("id")

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit := 10
	offset := (page - 1) * limit

	query := models.DB.Model(&models.Room{})

	var totalRecords int64
	query.Count(&totalRecords)
	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))

	query = query.Where("accommodation_id = ?", accommodationId)

	result = query.
		Select("room.id as room_id, room.room_number, room.room_name, " +
			"room.room_description, room.capacity, room.area, room.accommodation_id, " +
			"town.name as town_name, country.country_name as country_name").
		Offset(offset).Limit(limit).
		Find(&rooms)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "rooms fetched successfully",
		"data":         rooms,
		"page":         page,
		"limit":        limit,
		"totalPages":   totalPages,
		"totalRecords": totalRecords,
	})
}
