package controllers

import (
	"net/http"
	"strconv"

	"backend/models"

	"github.com/gin-gonic/gin"
)

func CreateOffer(c *gin.Context, tableName string, offer models.OfferOperations) {
	if err := c.ShouldBind(offer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	if err := offer.Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	id, err := offer.GetID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": "GetID error: " + err.Error()})
		return
	}

	if err := offer.HandleOfferImageUploads(c, tableName, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": "Failed to handle offer image uploads: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Offer created successfully!", "offer": offer})
}

type OfferQueryParameters struct {
	Model    interface{}
	Preloads []string
	Filters  map[string]interface{}
}

func FetchOffers(c *gin.Context, params OfferQueryParameters) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	query := models.DB
	for _, preload := range params.Preloads {
		query = query.Preload(preload)
	}

	for key, value := range params.Filters {
		query = query.Where(key+" = ?", value)
	}

	name := c.Query("title")
	if name != "" {
		query = query.Where("title ILIKE ?", "%"+name+"%")
	}

	var tableName string
	switch params.Model.(type) {
	case *[]models.Event:
		tableName = "event"
	case *[]models.Activity:
		tableName = "activity"
	case *[]models.Accommodation:
		tableName = "accommodation"
	default:
		tableName = ""
	}

	localisation := c.Query("localisation")
	if localisation != "" {
		query = query.Joins("JOIN town ON town.id = "+tableName+".town_id").
			Joins("JOIN country ON country.id = town.country_id").
			Where("town.name ILIKE ? OR country.country_name ILIKE ?", "%"+localisation+"%", "%"+localisation+"%")
	}

	query = query.Offset(offset).Limit(pageSize).Find(params.Model)

	if query.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": query.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Data fetched successfully",
		"data":      params.Model,
		"page":      page,
		"page_size": pageSize,
	})
}

func DeleteOffer(c *gin.Context, getByID func(string) (models.OfferOperations, error)) {
	id := c.Params.ByName("id")

	offer, err := getByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Failed to get offer by ID: " + err.Error()})
		return
	}

	if err = offer.Delete(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Failed to delete offer: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Offer deleted", "data": offer})
}
