package controllers

import (
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
)

func CreateRooms(c *gin.Context) {
	var rooms []models.Room

	if err := c.ShouldBindJSON(&rooms); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	tx := models.DB.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": "Failed to start transaction"})
		return
	}

	for _, room := range rooms {
		if err := tx.Create(&room).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Failed to save room: " + err.Error()})
			return
		}

		for _, facility := range room.RoomFacilities {
			if err := tx.Where("name = ?", facility.Name).First(&facility).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Failed to find facility: " + err.Error()})
				return
			}
			if err := tx.Model(&room).Association("RoomFacilities").Append(&facility); err != nil {
				tx.Rollback()
				c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Failed to associate facility with room: " + err.Error()})
				return
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": "Failed to commit transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Rooms created successfully!", "rooms": rooms})
}

func DeleteRoom(c *gin.Context) {
	roomID := c.Param("id")

	var room models.Room
	if err := models.DB.First(&room, roomID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found", "message": "Room not found"})
		return
	}

	if err := room.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": "Failed to delete room"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room deleted successfully"})
}
