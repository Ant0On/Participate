package controllers

import (
	"backend/models"

	"github.com/gin-gonic/gin"
)

func AddAccommodationReservation(c *gin.Context) {
	var reservation *models.ReservationAccommodation
	CreateReservation(c, reservation)
}

func GetAccommodationReservationById(c *gin.Context) {
	GetReservationById(c, models.GetAccommodationReservationById)
}

func GetAccommodationReservationsByState(c *gin.Context) {
	GetReservationByState(c, models.GetAccommodationReservationsByState)
}

func ChangeAccommodationReservationState(c *gin.Context) {
	ChangeReservationState(c, models.GetAccommodationReservationById)
}
