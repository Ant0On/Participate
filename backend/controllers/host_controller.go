package controllers

import (
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
)

func GetHostByID(c *gin.Context) {
	hostID := c.Param("id")

	if hostID == "" {
		c.JSON(400, gin.H{"error": "Host ID is required"})
		return
	}

	host, err := models.GetHost(hostID)

	if err != nil {
		c.JSON(404, gin.H{"error": "Host not found"})
		return
	}

	c.JSON(http.StatusOK, host)
}
