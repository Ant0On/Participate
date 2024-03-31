package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOffer(c *gin.Context, offer interface{}) {
	if err := c.ShouldBind(offer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"c.ShouldBind: ": err.Error()})
		return
	}

	if err := offer.(OfferSaver).Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"offer.Save: ": err.Error()})
		return
	}

	id, err := offer.(OfferIdentifier).GetID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"GetID error": err.Error()})
		return
	}

	if err := offer.(OfferUploader).HandleOfferImageUploads(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"offer.HandleOfferImageUploads: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "offer created successfully!", "offer": offer})
}

type OfferSaver interface {
	Save() error
	Update() error
	Delete() error
}

type OfferUploader interface {
	HandleOfferImageUploads(c *gin.Context, id uint) error
}

type OfferIdentifier interface {
	GetID() (uint, error)
}
