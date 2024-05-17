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
		c.JSON(http.StatusBadRequest, gin.H{"c.ShouldBind: ": err.Error()})
		return
	}

	for _, room := range rooms {
		if err := room.Save(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"room.Save: ": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "rooms created successfully!", "rooms": rooms})
}

func GetRooms(c *gin.Context) {
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
			"town.name as town_name, country.name as country_name").
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
