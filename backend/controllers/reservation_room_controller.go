package controllers

import (
	"backend/models"
	"backend/models/DTO"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

func AddRoomReservation(c *gin.Context) {
	var reservation models.ReservationRoom
	CreateReservation(c, &reservation)
}

func GetRoomReservationById(c *gin.Context) {
	GetReservationById(c, models.GetRoomReservationById)
}

func GetRoomReservationsByState(c *gin.Context) {
	GetReservationByState(c, models.GetRoomReservationsByState)
}

func ChangeRoomReservationState(c *gin.Context) {
	ChangeReservationState(c, models.GetRoomReservationById)
}

func GetCurrentRoomReservations(c *gin.Context) {
	var currentReservations []DTO.ReservationRoomWithOffer
	selectQuery := "reservation_room.id as reservation_id, reservation_room.date_from," +
		" reservation_room.date_to, room.capacity, accommodation.price_per_day as price_per_day," +
		" room.room_name as name, reservation_room.reservation_state, room.id as room_id"
	GetDTOReservation(c, ReservationQueryParameters{
		offerTableName:       "room",
		reservationTableName: "reservation_room",
		offerID:              "room_id",
		model:                &models.ReservationRoom{},
		dto:                  &currentReservations,
		selectQuery:          selectQuery,
		condition:            utils.CurrentWhereCondition,
		userRole:             "host",
	})
}

func GetReservationsRoomHistory(c *gin.Context) {
	var reservationsHistory []DTO.ReservationRoomWithOffer
	selectQuery := "reservation_room.id as reservation_id, reservation_room.date_from," +
		" reservation_room.date_to, room.capacity, reservation_room.rating_id, accommodation.price_per_day as price_per_day," +
		" room.room_name as name, reservation_room.reservation_state, room.id as room_id"
	GetDTOReservation(c, ReservationQueryParameters{
		offerTableName:       "room",
		reservationTableName: "reservation_room",
		offerID:              "room_id",
		model:                &models.ReservationRoom{},
		dto:                  &reservationsHistory,
		selectQuery:          selectQuery,
		condition:            utils.HistoryWhereCondition,
		userRole:             "customer",
	})
}
