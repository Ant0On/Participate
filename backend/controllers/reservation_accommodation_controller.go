package controllers

import (
	"fmt"

	"backend/models"
	"backend/models/DTO"

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

func PendingWhereCondition(userID string) string {
	return fmt.Sprintf("app_user.id = '%s' AND reservation_state = 'pending'", userID)
}

func HistoryWhereCondition(userID string) string {
	return fmt.Sprintf("app_user.id = '%s' AND reservation_state in ('finished', 'accepted', 'rejected')", userID)
}

func GetPendingAccommodationReservations(c *gin.Context) {
	var pendingReservations []DTO.ReservationAccommodationWithOffer
	selectQuery := "reservation_accommodation.id as reservation_accommodation_id, reservation_accommodation.date_from," +
		" reservation_accommodation.date_to, accommodation.capacity, reservation_accommodation.rating_id," +
		" accommodation.title, accommodation.price_per_day, accommodation.is_animal_friendly, town.name as town_name," +
		" country.name as country_name, accommodation.accommodation_type, reservation_accommodation.reservation_state, accommodation.id as accommodation_id"
	GetDTOReservation(c, ReservationQueryParameters{
		offerTableName:       "accommodation",
		reservationTableName: "reservation_accommodation",
		offerID:              "accommodation_id",
		model:                &models.ReservationAccommodation{},
		dto:                  &pendingReservations,
		selectQuery:          selectQuery,
		condition:            PendingWhereCondition,
	})
}

func GetReservationsAccommodationHistory(c *gin.Context) {
	var reservationsHistory []DTO.ReservationAccommodationWithOffer
	selectQuery := "reservation_accommodation.id as reservation_accommodation_id, reservation_accommodation.date_from," +
		" reservation_accommodation.date_to, accommodation.capacity, reservation_accommodation.rating_id," +
		" accommodation.title, accommodation.price_per_day, accommodation.is_animal_friendly, town.name as town_name," +
		" country.name as country_name, accommodation.accommodation_type, reservation_accommodation.reservation_state, accommodation.id as accommodation_id"
	GetDTOReservation(c, ReservationQueryParameters{
		offerTableName:       "accommodation",
		reservationTableName: "reservation_accommodation",
		offerID:              "accommodation_id",
		model:                &models.ReservationAccommodation{},
		dto:                  &reservationsHistory,
		selectQuery:          selectQuery,
		condition:            HistoryWhereCondition,
	})
}
