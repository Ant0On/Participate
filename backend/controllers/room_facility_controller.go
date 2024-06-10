package controllers

import (
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
)

func GetAllRoomFacilities(c *gin.Context) {
	roomFacilities, err := models.GetAllRoomFacilities()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Failed to get all room facilities: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": roomFacilities})
}
