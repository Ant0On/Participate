package controllers

import (
	"net/http"

	"backend/models"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

func CreateReservation(c *gin.Context, reservation models.ReservationOperations) {
	if err := c.ShouldBind(reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"c.ShouldBind: ": err.Error()})
		return
	}

	if err := reservation.Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"offer.Save: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "offer created successfully!", "offer": reservation})
}

func GetReservationById(c *gin.Context, getByID func(string) (models.ReservationOperations, error)) {
	id := c.Params.ByName("id")

	reservation, err := getByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"models.getReservationByID:": err.Error()})
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
