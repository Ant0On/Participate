package controllers

import (
	"math"
	"net/http"
	"strconv"

	"backend/logger"
	"backend/models"
	"backend/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOffer(c *gin.Context, tableName string, offer models.OfferOperations) {
	if err := c.ShouldBind(offer); err != nil {
		logger.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"c.ShouldBind: ": err.Error()})
		return
	}

	if err := offer.Save(); err != nil {
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
	tableName   string
	model       interface{}
	dto         interface{}
	selectQuery string
}

func GetOffers(c *gin.Context, parameters OfferQueryParameters) {
	var result *gorm.DB

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit := 10
	offset := (page - 1) * limit

	query := models.DB.Model(parameters.model)

	var totalRecords int64
	query.Count(&totalRecords)
	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))

	joinCondition := "JOIN town ON " + parameters.tableName + ".town_id = town.id"
	joinCondition += " JOIN country ON town.country_id = country.id"

	result = query.
		Joins(joinCondition).
		Select(parameters.selectQuery).
		Offset(offset).Limit(limit).
		Find(parameters.dto)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "offers fetched successfully",
		"data":         parameters.dto,
		"page":         page,
		"limit":        limit,
		"totalPages":   totalPages,
		"totalRecords": totalRecords,
	})
}

func GetOfferByID(c *gin.Context, parameters OfferQueryParameters) {
	offerID := c.Param("id")

	var result *gorm.DB

	joinCondition := "JOIN town ON " + parameters.tableName + ".town_id = town.id"
	joinCondition += " JOIN country ON town.country_id = country.id"

	result = models.DB.
		Model(parameters.model).
		Joins(joinCondition).
		Where(parameters.tableName+".id = ?", offerID).
		Select(parameters.selectQuery).
		Find(parameters.dto)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Offer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "offer fetched successfully", "data": parameters.dto})
}

func GetOffersForHost(c *gin.Context, parameters OfferQueryParameters) {
	var result *gorm.DB
	hostID := c.Param("id")

	joinCondition := "JOIN town ON " + parameters.tableName + ".town_id = town.id"
	joinCondition += " JOIN country ON town.country_id = country.id"

	result = models.DB.
		Model(parameters.model).
		Joins(joinCondition).
		Where(parameters.tableName+".user_id = ?", hostID).
		Select(parameters.selectQuery).
		Find(parameters.dto)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Offer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "offers fetched successfully", "data": parameters.dto})
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
		c.JSON(http.StatusBadRequest, gin.H{"offer.Delete: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "offer deleted", "data": offer})
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
