package DTO

import (
	"time"

	"backend/models"
)

type ReservationWithOffer struct {
	ReservationID    uint             `json:"reservation_id" binding:"required"`
	Name             string           `json:"name" binding:"required,min=2"`
	Price            float64          `json:"price" binding:"required,gt=0"`
	NumberOfPeople   int              `json:"number_of_people" binding:"required,gt=0"`
	DateFrom         time.Time        `json:"date_from" binding:"required"`
	DateTo           time.Time        `json:"date_to" binding:"required,gtefield=DateFrom"`
	IsAnimalFriendly bool             `json:"is_animal_friendly"`
	OfferType        models.OfferType `json:"offer_type" binding:"required,oneof=activity event accommodation"`
	TownName         string           `json:"town_name" binding:"required,min=2"`
	CountryName      string           `json:"country_name" binding:"required,min=3"`
	OfferID          uint             `json:"offer_id" binding:"required"`
	GradeId          uint             `json:"grade_id"`
	ReservationState string           `json:"reservation_state" binding:"required,oneof=pending accepted ongoing finished rejected"`
}
