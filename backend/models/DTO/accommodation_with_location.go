package DTO

import "backend/models"

type AccommodationWithLocation struct {
	OfferWithLocation
	PricePerDay      float64                  `json:"price_per_day" binding:"required,gt=0"`
	IsAnimalFriendly bool                     `json:"is_animal_friendly"`
	Type             models.AccommodationType `json:"type" binding:"required,oneof=hotel hostel apartment villa guesthouse"`
	Rating           int                      `json:"rating" binding:"required,min=1,max=5"`
}
