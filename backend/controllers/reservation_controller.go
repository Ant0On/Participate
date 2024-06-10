package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"backend/models"
	"backend/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReservationQueryParameters struct {
	offerTableName       string
	reservationTableName string
	offerID              string
	model                interface{}
	dto                  interface{}
	selectQuery          string
	condition            func(userID string) string
	userRole             string
}

func CreateReservation(c *gin.Context, reservation models.ReservationOperations) {
	if err := c.ShouldBindJSON(reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	if err := reservation.ChangeCapacity(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": "Failed to change capacity: " + err.Error()})
		return
	}

	if err := reservation.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": "Failed to save reservation: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reservation created successfully!", "reservation": reservation})
}

func GetReservationById(c *gin.Context, getByID func(string) (models.ReservationOperations, error)) {
	id := c.Param("id")

	reservation, err := getByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": "Failed to get reservation by ID: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reservation})
}

func GetReservationByState(c *gin.Context, getByState func(string) ([]models.ReservationOperations, error)) {
	state := c.Param("state")

	reservation, err := getByState(state)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Failed to get reservation by state: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reservation})
}

func ChangeReservationState(c *gin.Context, getByID func(string) (models.ReservationOperations, error)) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "ID is required"})
		return
	}

	state := c.Param("state")
	if state == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "State is required"})
		return
	}

	if !utils.CheckState(state) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "Wrong state"})
		return
	}

	reservation, err := getByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found", "message": "Failed to get reservation by ID: " + err.Error()})
		return
	}

	if err := reservation.ChangeState(state); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "message": "Failed to change reservation state: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success", "reservation": reservation})
}

func GetDTOReservation(c *gin.Context, parameters ReservationQueryParameters) {
	userID := c.Param("id")

	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": "App user ID is required"})
		return
	}

	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt <= 0 {
		pageInt = 1
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeInt <= 0 {
		pageSizeInt = 10
	}

	offset := (pageInt - 1) * pageSizeInt

	var result *gorm.DB

	query := models.DB.Model(parameters.model)

	joinCondition := fmt.Sprintf("JOIN %s ON %s.%s = %s.id", parameters.offerTableName, parameters.reservationTableName, parameters.offerID, parameters.offerTableName)
	if parameters.offerTableName == "room" {
		joinCondition += " JOIN accommodation ON room.accommodation_id = accommodation.id"
		joinCondition += " JOIN town ON accommodation.town_id = town.id"
	} else {
		joinCondition += fmt.Sprintf(" JOIN town ON %s.town_id = town.id", parameters.offerTableName)
	}
	joinCondition += " JOIN country ON town.country_id = country.id"
	if parameters.offerTableName == "room" && parameters.userRole == "host" {
		joinCondition += " JOIN app_user ON accommodation.user_id = app_user.id"
	} else if parameters.userRole == "customer" {
		joinCondition += fmt.Sprintf(" JOIN app_user ON %s.user_id = app_user.id", parameters.reservationTableName)
	} else {
		joinCondition += fmt.Sprintf(" JOIN app_user ON %s.user_id = app_user.id", parameters.offerTableName)
	}

	result = query.
		Joins(joinCondition).
		Where(parameters.condition(userID)).
		Select(parameters.selectQuery).
		Offset(offset).
		Limit(pageSizeInt).
		Find(parameters.dto)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request", "message": err.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNoContent, gin.H{"message": "No reservations found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Reservations fetched successfully",
		"data":      parameters.dto,
		"page":      pageInt,
		"page_size": pageSizeInt,
	})
}
