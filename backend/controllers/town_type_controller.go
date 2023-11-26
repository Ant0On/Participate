package controllers

import (
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
)

func GetAllTownTypes(c *gin.Context) {
	townType, err := models.GetAllTownTypes()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"GetAllTownTypes error:": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": townType})
}
