package DTO

import (
	"time"

	"backend/models"
)

type ReservationAccommodationWithOffer struct {
	ReservationID     uint                     `json:"reservation_id" binding:"required"`
	Title             string                   `json:"title" binding:"required,min=2"`
	PricePerDay       float64                  `json:"price_per_day" binding:"required,gt=0"`
	Capacity          int                      `json:"capacity" binding:"required,gt=0"`
	DateFrom          time.Time                `json:"date_from" binding:"required"`
	DateTo            time.Time                `json:"date_to" binding:"required,gtefield=DateFrom"`
	IsAnimalFriendly  bool                     `json:"is_animal_friendly"`
	AccommodationType models.AccommodationType `json:"accommodation_type" binding:"required,oneof=hotel hostel apartment villa guesthouse"`
	TownName          string                   `json:"town_name" binding:"required,min=2"`
	CountryName       string                   `json:"country_name" binding:"required,min=3"`
	AccommodationID   uint                     `json:"accommodation_id" binding:"required"`
	RatingID          uint                     `json:"rating_id"`
	ReservationState  string                   `json:"reservation_state" binding:"required,oneof=pending accepted ongoing finished rejected"`
}

type ReservationActivityWithOffer struct {
	ReservationID    uint                `json:"reservation_id" binding:"required"`
	Title            string              `json:"title" binding:"required,min=2"`
	Price            float64             `json:"price" binding:"required,gt=0"`
	Capacity         int                 `json:"capacity" binding:"required,gt=0"`
	Date             time.Time           `json:"date" binding:"required"`
	ActivityType     models.ActivityType `json:"activity_type" binding:"required,oneof=indoor outdoor"`
	TownName         string              `json:"town_name" binding:"required,min=2"`
	CountryName      string              `json:"country_name" binding:"required,min=3"`
	ActivityID       uint                `json:"activity_id" binding:"required"`
	RatingID         uint                `json:"rating_id"`
	ReservationState string              `json:"reservation_state" binding:"required,oneof=pending accepted ongoing finished rejected"`
}

type ReservationEventWithOffer struct {
	ReservationID    uint             `json:"reservation_id" binding:"required"`
	Title            string           `json:"title" binding:"required,min=2"`
	Price            float64          `json:"price" binding:"required,gt=0"`
	Capacity         int              `json:"capacity" binding:"required,gt=0"`
	Date             time.Time        `json:"date" binding:"required"`
	EventType        models.EventType `json:"event_type" binding:"required,oneof=conference concert festival 'sports event'"`
	TownName         string           `json:"town_name" binding:"required,min=2"`
	CountryName      string           `json:"country_name" binding:"required,min=3"`
	EventID          uint             `json:"event_id" binding:"required"`
	RatingID         uint             `json:"rating_id"`
	ReservationState string           `json:"reservation_state" binding:"required,oneof=pending accepted ongoing finished rejected"`
}
