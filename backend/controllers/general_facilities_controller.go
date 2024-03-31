package controllers

import (
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
)

func GetAllGeneralFacilities(c *gin.Context) {
	generalFacilities, err := models.GetAllGeneralFacilities()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"GetAllGeneralFacilities error:": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": generalFacilities})
}
