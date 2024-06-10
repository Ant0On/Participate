package controllers

import (
	"backend/models"
	"backend/models/DTO"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

func AddActivityReservation(c *gin.Context) {
	var reservation models.ReservationActivity
	CreateReservation(c, &reservation)
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

func GetCurrentActivityReservations(c *gin.Context) {
	var pendingReservations []DTO.ReservationActivityWithOffer
	selectQuery := "reservation_activity.id as reservation_id, reservation_activity.date, activity.capacity," +
		" reservation_activity.rating_id, activity.title, activity.price, activity.skill_level as skill, activity.activity_type," +
		" town.name as town_name, country.country_name as country_name, reservation_activity.reservation_state, activity.id as activity_id"

	GetDTOReservation(c, ReservationQueryParameters{
		offerTableName:       "activity",
		reservationTableName: "reservation_activity",
		offerID:              "activity_id",
		model:                &models.ReservationActivity{},
		dto:                  &pendingReservations,
		selectQuery:          selectQuery,
		condition:            utils.CurrentWhereCondition,
		userRole:             "host",
	})
}
func GetReservationsActivityHistory(c *gin.Context) {
	var pendingReservations []DTO.ReservationActivityWithOffer
	selectQuery := "reservation_activity.id as reservation_id, reservation_activity.date, activity.capacity," +
		" reservation_activity.rating_id, activity.title, activity.price, activity.skill_level as skill, activity.activity_type," +
		" town.name as town_name, country.country_name as country_name, reservation_activity.reservation_state, activity.id as activity_id"

	GetDTOReservation(c, ReservationQueryParameters{
		offerTableName:       "activity",
		reservationTableName: "reservation_activity",
		offerID:              "activity_id",
		model:                &models.ReservationActivity{},
		dto:                  &pendingReservations,
		selectQuery:          selectQuery,
		condition:            utils.HistoryWhereCondition,
		userRole:             "customer",
	})
}
