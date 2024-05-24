package controllers

import (
	"backend/models"
	"backend/models/DTO"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

func AddAccommodationReservation(c *gin.Context) {
	var reservation models.ReservationAccommodation
	CreateReservation(c, &reservation)
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

func GetCurrentAccommodationReservations(c *gin.Context) {
	var currentReservations []DTO.ReservationAccommodationWithOffer
	selectQuery := "reservation_accommodation.id as reservation_id, reservation_accommodation.date_from," +
		" reservation_accommodation.date_to, accommodation.capacity, reservation_accommodation.rating_id," +
		" accommodation.title, accommodation.price_per_day, accommodation.is_animal_friendly, town.name as town_name," +
		" country.name as country_name, accommodation.accommodation_type as type, reservation_accommodation.reservation_state, accommodation.id as accommodation_id"
	GetDTOReservation(c, ReservationQueryParameters{
		offerTableName:       "accommodation",
		reservationTableName: "reservation_accommodation",
		offerID:              "accommodation_id",
		model:                &models.ReservationAccommodation{},
		dto:                  &currentReservations,
		selectQuery:          selectQuery,
		condition:            utils.CurrentWhereCondition,
		userRole:             "host",
	})
}

func GetReservationsAccommodationHistory(c *gin.Context) {
	var reservationsHistory []DTO.ReservationAccommodationWithOffer
	selectQuery := "reservation_accommodation.id as reservation_id, reservation_accommodation.date_from," +
		" reservation_accommodation.date_to, accommodation.capacity, reservation_accommodation.rating_id," +
		" accommodation.title, accommodation.price_per_day, accommodation.is_animal_friendly, town.name as town_name," +
		" country.name as country_name, accommodation.accommodation_type as type, reservation_accommodation.reservation_state, accommodation.id as accommodation_id"
	GetDTOReservation(c, ReservationQueryParameters{
		offerTableName:       "accommodation",
		reservationTableName: "reservation_accommodation",
		offerID:              "accommodation_id",
		model:                &models.ReservationAccommodation{},
		dto:                  &reservationsHistory,
		selectQuery:          selectQuery,
		condition:            utils.HistoryWhereCondition,
		userRole:             "customer",
	})
}
