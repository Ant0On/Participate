package controllers

import (
	"net/http"
	"strconv"

	"backend/models"

	"github.com/gin-gonic/gin"
)

func CreateChat(c *gin.Context) {
	offerId := c.Param("id")
	offer, err := models.GetOfferByID(offerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "offer not found"})
		return
	}

	offerID, err := strconv.Atoi(offerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	chat := models.Chat{
		OfferID: uint(offerID),
	}

	if err := chat.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	offer.ChatID = chat.ID

	if err := offer.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": chat})
}
