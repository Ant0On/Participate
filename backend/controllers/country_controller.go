package controllers

import (
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
)

func GetAllCountries(c *gin.Context) {
	countries, err := models.GetAllCountries()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"GetAllCountries error:": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": countries})
}
