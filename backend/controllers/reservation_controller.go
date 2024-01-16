package controllers

import (
	"net/http"

	"backend/models"

	"github.com/gin-gonic/gin"
)

func AddReservation(c *gin.Context) {
	var reservation *models.Reservation

	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error with AddReservation": err.Error()})
		return
	}

	if err := reservation.Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reservation.Save error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "reservation added successfully!"})
}

func GetReservationById(c *gin.Context) {
	reservationID := c.Param("id")

	if reservationID == "" {
		c.JSON(400, gin.H{"error": "reservation ID is required"})
		return
	}

	reservation, err := models.GetReservationById(reservationID)

	if err != nil {
		c.JSON(404, gin.H{"error": "Reservation not found"})
		return
	}

	c.JSON(http.StatusOK, reservation)
}

func ChangeReservationState(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "id is required"})
		return
	}
	state := c.Param("state")
	if state == "" {
		c.JSON(400, gin.H{"error": "state is required"})
		return
	}
	if !checkState(state) {
		c.JSON(400, gin.H{"error": "wrong state"})
		return
	}

	reservation, err := models.GetReservationById(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "reservation not found"})
		return
	}

	reservation.ReservationState = models.ReservationState(state)

	if err := reservation.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reservation.Save": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func checkState(state string) bool {
	switch state {
	case "pending":
		return true
	case "accepted":
		return true
	case "ongoing":
		return true
	case "finished":
		return true
	case "rejected":
		return true
	default:
		return false
	}
}
