package controllers

import (
	"backend/models"
	"backend/models/DTO"

	"github.com/gin-gonic/gin"
)

func AddRoomReservation(c *gin.Context) {
	var reservation *models.ReservationRoom
	CreateReservation(c, reservation)
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

func GetPendingRoomReservations(c *gin.Context) {
	var pendingReservations []DTO.ReservationRoomWithOffer
	selectQuery := "reservation_room.id as reservation_room_id, reservation_room.date_from," +
		" reservation_room.date_to, room.capacity," +
		" room.room_name, reservation_room.reservation_state, room.id as room_id"
	GetDTOReservation(c, ReservationQueryParameters{
		offerTableName:       "room",
		reservationTableName: "reservation_room",
		offerID:              "room_id",
		model:                &models.ReservationRoom{},
		dto:                  &pendingReservations,
		selectQuery:          selectQuery,
		condition:            PendingWhereCondition,
	})
}

func GetReservationsRoomHistory(c *gin.Context) {
	var reservationsHistory []DTO.ReservationRoomWithOffer
	selectQuery := "reservation_room.id as reservation_room_id, reservation_room.date_from," +
		" reservation_room.date_to, room.capacity," +
		" room.room_name, reservation_room.reservation_state, room.id as room_id"
	GetDTOReservation(c, ReservationQueryParameters{
		offerTableName:       "room",
		reservationTableName: "reservation_room",
		offerID:              "room_id",
		model:                &models.ReservationRoom{},
		dto:                  &reservationsHistory,
		selectQuery:          selectQuery,
		condition:            HistoryWhereCondition,
	})
}
