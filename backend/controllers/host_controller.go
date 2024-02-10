package controllers

import (
	"net/http"

	"backend/models"
	"backend/models/DTO"

	"github.com/gin-gonic/gin"
)

func GetHostByID(c *gin.Context) {
	hostID := c.Param("id")

	if hostID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Host ID is required"})
		return
	}

	host, err := models.GetHost(hostID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Host not found"})
		return
	}

	c.JSON(http.StatusOK, host)
}

func GetPendingReservations(c *gin.Context) {
	hostID := c.Param("id")

	if hostID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Host ID is required"})
		return
	}

	var pendingReservations []DTO.ReservationWithOffer
	result := models.DB.
		Model(&models.Reservation{}).
		Joins("JOIN offer ON reservation.offer_id = offer.id").
		Joins("JOIN town ON offer.town_id = town.id").
		Joins("JOIN country ON town.country_id = country.id").
		Joins("JOIN host ON offer.host_id = host.id").
		Where("host.id = ? AND reservation_state = 'pending'", hostID).
		Select("reservation.id as reservation_id, reservation.date_from, reservation.date_to, reservation.number_of_people, offer.name," + "" +
			"offer.price, offer.is_animal_friendly, offer.offer_type, town.name as town_name, country.name as country_name, offer.id as offer_id").
		Find(&pendingReservations)

	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No pending reservations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "pending reservations fetched successfully", "data": pendingReservations})
}
