package controllers

import (
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
)

func GetAllEquipment(c *gin.Context) {
	equipment, err := models.GetAllEquipment()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Failed to get all equipment: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": equipment})
}
