package controllers

import (
	"net/http"
	"strconv"

	"backend/logger"
	"backend/models"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

func CreateOffer(c *gin.Context, tableName string, offer models.OfferOperations) {
	if err := c.ShouldBind(offer); err != nil {
		logger.Logger.Errorf(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"c.ShouldBind: ": err.Error()})
		return
	}

	if err := offer.Save(); err != nil {
		logger.Logger.Errorf(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"offer.Save: ": err.Error()})
		return
	}

	id, err := offer.GetID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"GetID error": err.Error()})
		return
	}

	if err := offer.HandleOfferImageUploads(c, tableName, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"offer.HandleOfferImageUploads: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "offer created successfully!", "offer": offer})
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

	query = query.Offset(offset).Limit(pageSize).Find(params.Model)

	if query.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": query.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "data fetched successfully",
		"data":      params.Model,
		"page":      page,
		"page_size": pageSize,
	})
}

func DeleteOffer(c *gin.Context, getByID func(string) (models.OfferOperations, error)) {
	id := c.Params.ByName("id")

	offer, err := getByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"models.OfferByID:": err.Error()})
		return
	}

	if err = offer.Delete(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"offer.Delete: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "offer deleted", "data": offer})
}

func UpdateOffer(c *gin.Context, getByID func(string) (models.OfferOperations, error)) {
	id := c.Params.ByName("id")

	offer, err := getByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"models.OfferByID:": err.Error()})
		return
	}

	if err = offer.Update(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"offer.Update: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "offer updated", "data": offer})
}

func DiscountOffer(c *gin.Context, getByID func(string) (models.OfferOperations, error)) {
	id := c.Params.ByName("id")
	var req utils.DiscountRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	offer, err := getByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if offer == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "offer not found"})
		return
	}

	if err := offer.AddDiscount(req.Discount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": offer})
}

func ChangeOfferPrice(c *gin.Context, getByID func(string) (models.OfferOperations, error)) {
	id := c.Params.ByName("id")
	var req utils.ChangePriceReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	offer, err := getByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if offer == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "offer not found"})
		return
	}

	if err := offer.UpdatePrice(req.Price); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": offer})
}
