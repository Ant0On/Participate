package DTO

import (
	"time"

	"backend/models"
)

type OfferWithLocation struct {
	OfferID     uint    `json:"offer_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Capacity    int     `json:"capacity"`
	TownName    string  `json:"town_name"`
	CountryName string  `json:"country_name"`
	UserID      uint    `json:"user_id"`
	Discount    float64 `json:"discount"`
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

type AccommodationDTO struct {
	OfferWithLocation
	Capacity          int           `json:"capacity"`
	Discount          float64       `json:"discount"`
	TownID            uint          `json:"town_id"`
	UserID            uint          `json:"user_id"`
	NumberOfRooms     int           `json:"number_of_rooms"`
	Type              string        `json:"type"`
	IsAnimalFriendly  bool          `json:"is_animal_friendly"`
	PricePerDay       float64       `json:"price_per_day"`
	GeneralFacilities []string      `json:"general_facilities"`
	Rooms             []models.Room `json:"rooms"`
}
