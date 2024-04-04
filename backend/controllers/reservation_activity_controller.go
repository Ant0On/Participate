package controllers

import (
	"backend/models"

	"github.com/gin-gonic/gin"
)

func AddActivityReservation(c *gin.Context) {
	var reservation *models.ReservationActivity
	CreateReservation(c, reservation)
}

func GetActivityReservationById(c *gin.Context) {
	GetReservationById(c, models.GetActivityReservationById)
}

func GetActivityReservationsByState(c *gin.Context) {
	GetReservationByState(c, models.GetActivityReservationsByState)
}

func ChangeActivityReservationState(c *gin.Context) {
	ChangeReservationState(c, models.GetActivityReservationById)
}
