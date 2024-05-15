package controllers

import (
	"fmt"
	"net/http"

	"backend/models"
	"backend/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateReservation(c *gin.Context, reservation models.ReservationOperations) {
	if err := c.ShouldBindJSON(reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"c.ShouldBind: ": err.Error()})
		return
	}

	if err := reservation.ChangeCapacity(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reservation.ChangeCapacity: ": err.Error()})
		return
	}

	if err := reservation.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"offer.Save: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "reservation created successfully!", "reservation": reservation})
}

func GetReservationById(c *gin.Context, getByID func(string) (models.ReservationOperations, error)) {
	id := c.Params.ByName("id")

	reservation, err := getByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"models.getReservationByID:": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reservation})
}

func GetReservationByState(c *gin.Context, getByState func(string) ([]models.ReservationOperations, error)) {
	state := c.Param("state")

	reservation, err := getByState(state)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"models.getReservationByState": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reservation})
}

func ChangeReservationState(c *gin.Context, getByID func(string) (models.ReservationOperations, error)) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	state := c.Param("state")
	if state == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "state is required"})
		return
	}

	if !utils.CheckState(state) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong state"})
		return
	}

	reservation, err := getByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error: ": err.Error()})
		return
	}

	if err := reservation.ChangeState(state); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "reservation": reservation})
}

type ReservationQueryParameters struct {
	offerTableName       string
	reservationTableName string
	offerID              string
	model                interface{}
	dto                  interface{}
	selectQuery          string
	condition            func(userID string) string
}

func GetDTOReservation(c *gin.Context, parameters ReservationQueryParameters) {
	userID := c.Param("id")

	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "app_user ID is required"})
		return
	}

	var result *gorm.DB

	query := models.DB.Model(parameters.model)

	joinCondition := fmt.Sprintf("JOIN %s ON %s.%s = %s.id", parameters.offerTableName, parameters.reservationTableName, parameters.offerID, parameters.reservationTableName)
	joinCondition += fmt.Sprintf(" JOIN town ON %s.town_id = town.id", parameters.offerTableName)
	joinCondition += " JOIN country ON town.country_id = country.id"
	joinCondition += fmt.Sprintf(" JOIN app_user ON %s.user_id = app_user.id", parameters.reservationTableName)

	result = query.
		Joins(joinCondition).
		Where(parameters.condition(userID)).
		Select(parameters.selectQuery).
		Find(parameters.dto)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNoContent, gin.H{"message": "No pending reservations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "pending reservations fetched successfully", "data": parameters.dto})
}
