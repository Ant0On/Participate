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

	saved, err := town.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"town.Save error": err.Error()})
		return
	}

	if !saved {
		existingTown := models.Town{}
		if err := models.DB.Where("country_id = ? AND name = ?", town.CountryID, town.Name).First(&existingTown).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch existing town"})
			return
		}
		town.ID = existingTown.ID
	}

	c.JSON(http.StatusOK, gin.H{"message": "town added successfully!", "town": town})
}
