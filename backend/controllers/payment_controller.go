package controllers

import (
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
)

func GetAllPayments(c *gin.Context) {
	payment, err := models.GetAllPayments()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"GetAllPayments error:": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": payment})
}
