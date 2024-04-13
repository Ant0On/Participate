package DTO

import (
	"time"

	"backend/models"
)

type OfferWithLocation struct {
	OfferID       uint    `json:"offer_id" binding:"required"`
	Title         string  `json:"title" binding:"required,min=2"`
	Description   string  `json:"description" binding:"required,min=30"`
	Capacity      int     `json:"capacity" binding:"required,gt=0"`
	IsRecommended bool    `json:"is_recommended"`
	TownName      string  `json:"town_name" binding:"required,min=2"`
	CountryName   string  `json:"country_name" binding:"required,min=3"`
	UserID        uint    `json:"user_id" binding:"required"`
	Discount      float64 `json:"discount" binding:"required,min=0,max=100"`
}

type EventWithLocation struct {
	OfferWithLocation
	Price float64          `json:"price" binding:"required,gt=0"`
	Type  models.EventType `json:"type" binding:"required,oneof=conference concert festival 'sports event'"`
}

type ActivityWithLocation struct {
	OfferWithLocation
	Price    float64             `json:"price" binding:"required,gt=0"`
	Skill    models.SkillLevel   `json:"skill_level" binding:"required,oneof=beginner intermediate advanced"`
	Type     models.ActivityType `json:"type" binding:"required,oneof=indoor outdoor"`
	Duration time.Duration       `json:"duration" binding:"required"`
}

type AccommodationWithLocation struct {
	OfferWithLocation
	PricePerDay      float64                  `json:"price_per_day" binding:"required,gt=0"`
	IsAnimalFriendly bool                     `json:"is_animal_friendly"`
	Type             models.AccommodationType `json:"type" binding:"required,oneof=hotel hostel apartment villa guesthouse"`
	Rating           int                      `json:"rating" binding:"required,min=1,max=5"`
}
