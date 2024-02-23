package controllers

import (
	"net/http"
	"strconv"

	"backend/models"

	"github.com/gin-gonic/gin"
)

func SendMessage(c *gin.Context) {
	var message models.Message

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := c.Param("id")
	chatId := c.Param("chatId")

	user, err := models.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "AppUser not found"})
		return
	}

	chatID, err := strconv.Atoi(chatId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	message.Email = user.Email
	message.AppUserID = user.ID
	message.ChatID = uint(chatID)

	if err := message.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": message})

}

func GetAllMessages(c *gin.Context) {
	offerID := c.Param("offerID")

	chat, err := models.GetChatByOfferId(offerID)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": "Chat not found"})
		return
	}

	messages, err := models.GetAllMessages(chat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if messages == nil {
		c.JSON(http.StatusNoContent, gin.H{"error": "Messages not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": messages})
}
