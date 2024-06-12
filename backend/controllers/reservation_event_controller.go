package controllers

import (
	"backend/models"
	"backend/models/DTO"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

func AddEventReservation(c *gin.Context) {
	var reservation models.ReservationEvent
	CreateReservation(c, &reservation)
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

func GetCurrentEventReservations(c *gin.Context) {
	var currentReservations []DTO.ReservationEventWithOffer
	selectQuery := "reservation_event.id as reservation_id, reservation_event.date," +
		" event.capacity, event.title, event.price, " +
		"event.event_type, town.name as town_name, country.country_name as country_name, reservation_event.reservation_state, event.id as event_id"
	GetDTOReservation(c, ReservationQueryParameters{
		offerTableName:       "event",
		reservationTableName: "reservation_event",
		offerID:              "event_id",
		model:                &models.ReservationEvent{},
		dto:                  &currentReservations,
		selectQuery:          selectQuery,
		condition:            utils.CurrentWhereCondition,
		userRole:             "host",
	})
}

func GetReservationsEventHistory(c *gin.Context) {
	var reservationsHistory []DTO.ReservationEventWithOffer
	selectQuery := "reservation_event.id as reservation_id, reservation_event.date," +
		" event.capacity, event.title, event.price, " +
		"event.event_type, town.name as town_name, country.country_name as country_name, reservation_event.reservation_state, event.id as event_id"
	GetDTOReservation(c, ReservationQueryParameters{
		offerTableName:       "event",
		reservationTableName: "reservation_event",
		offerID:              "event_id",
		model:                &models.ReservationEvent{},
		dto:                  &reservationsHistory,
		selectQuery:          selectQuery,
		condition:            utils.HistoryWhereCondition,
		userRole:             "customer",
	})
}
