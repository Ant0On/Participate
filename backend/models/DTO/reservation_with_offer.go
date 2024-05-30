package DTO

import (
	"time"

	"backend/models"
)

type ReservationAccommodationWithOffer struct {
	ReservationID     uint                     `json:"reservation_id"`
	Title             string                   `json:"title"`
	PricePerDay       float64                  `json:"price_per_day"`
	Capacity          int                      `json:"capacity"`
	DateFrom          time.Time                `json:"date_from"`
	DateTo            time.Time                `json:"date_to"`
	IsAnimalFriendly  bool                     `json:"is_animal_friendly"`
	AccommodationType models.AccommodationType `json:"type"`
	TownName          string                   `json:"town_name"`
	CountryName       string                   `json:"country_name"`
	AccommodationID   uint                     `json:"accommodation_id"`
	RatingID          uint                     `json:"rating_id"`
	ReservationState  string                   `json:"reservation_state"`
}

type ReservationActivityWithOffer struct {
	ReservationID    uint                `json:"reservation_id"`
	Title            string              `json:"title"`
	Price            float64             `json:"price"`
	Capacity         int                 `json:"capacity"`
	Date             time.Time           `json:"date"`
	ActivityType     models.ActivityType `json:"type"`
	TownName         string              `json:"town_name"`
	CountryName      string              `json:"country_name"`
	ActivityID       uint                `json:"activity_id"`
	RatingID         uint                `json:"rating_id"`
	ReservationState string              `json:"reservation_state"`
}

type ReservationEventWithOffer struct {
	ReservationID    uint             `json:"reservation_id"`
	Title            string           `json:"title"`
	Price            float64          `json:"price"`
	Capacity         int              `json:"capacity"`
	Date             time.Time        `json:"date"`
	EventType        models.EventType `json:"type"`
	TownName         string           `json:"town_name"`
	CountryName      string           `json:"country_name"`
	EventID          uint             `json:"event_id"`
	ReservationState string           `json:"reservation_state"`
}

type ReservationRoomWithOffer struct {
	ReservationID    uint      `json:"reservation_id"`
	Name             string    `json:"name"`
	PricePerDay      float64   `json:"price_per_day"`
	Capacity         int       `json:"capacity"`
	DateFrom         time.Time `json:"date_from"`
	DateTo           time.Time `json:"date_to"`
	IsAnimalFriendly bool      `json:"is_animal_friendly"`
	RoomID           uint      `json:"room_id"`
	RatingID         uint      `json:"rating_id"`
	ReservationState string    `json:"reservation_state"`
}
