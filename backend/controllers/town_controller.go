package controllers

import (
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
)

func AddTown(c *gin.Context) {
	var town *models.Town

	if err := c.ShouldBindJSON(&town); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error with AddTown": err.Error()})
		return
	}

	if err := town.Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"town.Save error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "town added successfully!", "town": town})
}

func GetTown(c *gin.Context) {
	townID := c.Param("id")

	if townID == "" {
		c.JSON(400, gin.H{"error": "Town ID is required"})
		return
	}

	var town models.Town

	_, err := models.GetTown(&town, townID)

	if err != nil {
		c.JSON(404, gin.H{"error": "Town not found"})
		return
	}

	c.JSON(http.StatusOK, town)
}
