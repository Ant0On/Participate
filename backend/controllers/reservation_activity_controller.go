package controllers

import (
	"backend/models"
	"backend/models/DTO"

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

func GetPendingActivityReservations(c *gin.Context) {
	var pendingReservations []DTO.ReservationActivityWithOffer
	selectQuery := "reservation_activity.id as reservation_activity_id, reservation_activity.date, activity.capacity," +
		" reservation_activity.rating_id, activity.title, activity.price, activity.activity_type," +
		" town.name as town_name, country.name as country_name, reservation_activity.reservation_state, activity.id as activity_id"

	GetDTOReservation(c, ReservationQueryParameters{
		offerTableName:       "activity",
		reservationTableName: "reservation_activity",
		offerID:              "activity_id",
		model:                &models.ReservationActivity{},
		dto:                  &pendingReservations,
		selectQuery:          selectQuery,
		condition:            PendingWhereCondition,
	})
}
func GetReservationsActivityHistory(c *gin.Context) {
	var pendingReservations []DTO.ReservationActivityWithOffer
	selectQuery := "reservation_activity.id as reservation_activity_id, reservation_activity.date, activity.capacity," +
		" reservation_activity.rating_id, activity.title, activity.price, activity.activity_type," +
		" town.name as town_name, country.name as country_name, reservation_activity.reservation_state, activity.id as activity_id"

	GetDTOReservation(c, ReservationQueryParameters{
		offerTableName:       "activity",
		reservationTableName: "reservation_activity",
		offerID:              "activity_id",
		model:                &models.ReservationActivity{},
		dto:                  &pendingReservations,
		selectQuery:          selectQuery,
		condition:            HistoryWhereCondition,
	})
}
