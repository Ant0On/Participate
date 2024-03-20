package controllers

import (
	"net/http"

	"backend/models"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

func AddAccommodationReservation(c *gin.Context) {
	var reservation *models.ReservationAccommodation
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"c.ShouldBindJSON: ": err.Error()})
		return
	}

	if err := reservation.ValidateDates(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reservation.ValidateDates: ": err.Error()})
		return
	}

	if err := reservation.Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reservation.Save: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "reservation added successfully!", "reservation": reservation})
}

func GetAccommodationReservationById(c *gin.Context) {
	reservationID := c.Param("id")
	if reservationID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "reservation ID is required"})
		return
	}

	reservation, err := models.GetAccommodationReservationById(reservationID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"models.GetAccommodationReservationById: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reservation)
}

func GetAccommodationReservationsByState(c *gin.Context) {
	reservationState := c.Param("state")
	if reservationState == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "reservation state is required"})
		return
	}

	reservations, err := models.GetAccommodationReservationsByState(reservationState)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"models.GetAccommodationReservationsByState ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reservations)
}

func ChangeAccommodationReservationState(c *gin.Context) {
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

	reservation, err := models.GetAccommodationReservationById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"models.GetAccommodationReservationById: ": err.Error()})
		return
	}

	reservation.ReservationState = models.ReservationState(state)

	if err := reservation.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reservation.Update: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "reservation": reservation})
}
