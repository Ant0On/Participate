package controllers

import (
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
)

func GetCountryByName(c *gin.Context) {
	name := c.Params.ByName("name")
	country, err := models.GetCountryByName(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"GetCountryBYName error:": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": country})
}
