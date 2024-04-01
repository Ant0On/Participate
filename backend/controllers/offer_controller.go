package controllers

import (
	"math"
	"net/http"
	"strconv"

	"backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

// TODO za duzo parametrów, dać do structa
func GetOffers(c *gin.Context, tableName string, model interface{}, dto interface{}, selectQuery string) {
	var result *gorm.DB

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit := 10
	offset := (page - 1) * limit

	query := models.DB.Model(model)

	var totalRecords int64
	query.Count(&totalRecords)
	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))

	joinCondition := "JOIN town ON " + tableName + ".town_id = town.id"
	joinCondition += " JOIN country ON town.country_id = country.id"

	result = query.
		Joins(joinCondition).
		Select(selectQuery).
		Offset(offset).Limit(limit).
		Find(dto)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "offers fetched successfully",
		"data":         dto,
		"page":         page,
		"limit":        limit,
		"totalPages":   totalPages,
		"totalRecords": totalRecords,
	})
}

// TODO za duzo parametrów, dać do structa
func GetOfferByID(c *gin.Context, tableName string, model interface{}, dto interface{}, selectQuery string) {
	offerID := c.Param("id")

	var result *gorm.DB

	joinCondition := "JOIN town ON " + tableName + ".town_id = town.id"
	joinCondition += " JOIN country ON town.country_id = country.id"

	result = models.DB.
		Model(model).
		Joins(joinCondition).
		Where(tableName+".id = ?", offerID).
		Select(selectQuery).
		Find(dto)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Offer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "offer fetched successfully", "data": dto})
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
