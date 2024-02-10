package DTO

import (
	"time"

	"backend/models"
)

type ReservationWithOffer struct {
	ReservationID    uint             `json:"reservation_id"`
	Name             string           `json:"name"`
	Price            float64          `json:"price"`
	DateFrom         time.Time        `json:"date_from"`
	DateTo           time.Time        `json:"date_to"`
	IsAnimalFriendly bool             `json:"is_animal_friendly"`
	OfferType        models.OfferType `json:"offer_type"`
	TownName         string           `json:"town_name"`
	CountryName      string           `json:"country_name"`
	OfferID          uint             `json:"offer_id"`
	ReservationState string           `json:"reservation_state"`
}
