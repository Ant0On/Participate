package controllers

import (
	"net/http"
	"strconv"

	"backend/models"

	"github.com/gin-gonic/gin"
)

func CreateChat(c *gin.Context) {
	offerId := c.Param("offerID")
	offer, err := models.GetOfferByID(offerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "offer not found"})
		return
	}

	if offer.OfferType != "event" {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "chat is available only for events"})
		return
	}

	offerID, err := strconv.Atoi(offerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	chat := models.Chat{
		OfferID:  uint(offerID),
		Messages: []models.Message{},
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

func GetChatByOfferID(c *gin.Context) {
	offerID := c.Param("offerID")

	chat, err := models.GetChatByOfferId(offerID)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": "Chat not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": chat})
}
