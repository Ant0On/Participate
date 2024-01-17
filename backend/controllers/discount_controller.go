package controllers

import (
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
)

func AddDiscount(c *gin.Context) {
	var discount *models.Discount

	if err := c.ShouldBindJSON(&discount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error with AddDiscount": err.Error()})
		return
	}

	if err := discount.Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"discount.Save error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "discount added successfully!", "data": discount})
}
