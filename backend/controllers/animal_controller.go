package controllers

import (
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
)

func AddAnimal(c *gin.Context) {
	var animal *models.Animal

	if err := c.ShouldBindJSON(&animal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error with AddAnimal": err.Error()})
		return
	}

	if err := animal.Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"animal.Save error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "animal added successfully!"})
}
