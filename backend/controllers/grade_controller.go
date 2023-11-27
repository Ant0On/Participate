package controllers

import (
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
)

func GetAllGrades(c *gin.Context) {
	grade, err := models.GetGrades()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"GetAllTownTypes error:": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": grade})
}
