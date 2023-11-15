package controllers

import (
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
)

func GetAllCountry(c *gin.Context) {
	country, err := models.GetAllCountry()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"GetAllCountry error:": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": country})
}
