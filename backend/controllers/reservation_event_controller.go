package controllers

import (
	"backend/models"

	"github.com/gin-gonic/gin"
)

func AddEventReservation(c *gin.Context) {
	var reservation *models.ReservationEvent
	CreateReservation(c, reservation)
}

func GetEventReservationById(c *gin.Context) {
	GetReservationById(c, models.GetEventReservationById)
}

func GetEventReservationsByState(c *gin.Context) {
	GetReservationByState(c, models.GetEventReservationsByState)
}

func ChangeEventReservationState(c *gin.Context) {
	ChangeReservationState(c, models.GetEventReservationById)
}
