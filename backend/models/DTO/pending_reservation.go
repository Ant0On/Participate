package DTO

import "backend/models"

type PendingReservation struct {
	ReservationID    uint             `json:"reservation_id"`
	Name             string           `json:"name"`
	Price            float64          `json:"price"`
	IsAnimalFriendly bool             `json:"is_animal_friendly"`
	OfferType        models.OfferType `json:"offer_type"`
	TownName         string           `json:"town_name"`
	CountryName      string           `json:"country_name"`
}
