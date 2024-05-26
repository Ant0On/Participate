package DTO

import (
	"time"

	"backend/models"
)

type OfferWithLocation struct {
	OfferID       uint    `json:"offer_id"`
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	Capacity      int     `json:"capacity"`
	IsRecommended bool    `json:"is_recommended"`
	TownName      string  `json:"town_name"`
	CountryName   string  `json:"country_name"`
	UserID        uint    `json:"user_id"`
	Discount      float64 `json:"discount"`
}

type EventWithLocation struct {
	OfferWithLocation
	Price float64          `json:"price"`
	Type  models.EventType `json:"type"`
}

type ActivityWithLocation struct {
	OfferWithLocation
	Price     float64             `json:"price"`
	Skill     models.SkillLevel   `json:"skill_level"`
	Type      models.ActivityType `json:"type"`
	Duration  time.Duration       `json:"duration"`
	Equipment string              `json:"equipment"`
}

type AccommodationWithLocation struct {
	OfferWithLocation
	PricePerDay      float64                  `json:"price_per_day"`
	IsAnimalFriendly bool                     `json:"is_animal_friendly"`
	Type             models.AccommodationType `json:"type"`
	Rating           int                      `json:"rating"`
}
