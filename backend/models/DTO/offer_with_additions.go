package DTO

import "time"

type AccommodationRoom struct {
	Title             string  `json:"title"`
	PricePerDay       float64 `json:"price_per_day"`
	RoomName          string  `json:"room_name"`
	RoomDescription   string  `json:"room_description"`
	Capacity          int     `json:"room_capacity"`
	Area              int     `json:"room_area"`
	RoomFacilities    string  `json:"room_facilities" type:"text"`
	GeneralFacilities string  `json:"general_facilities" type:"text"`
	AccommodationType string  `json:"accommodation_type"`
	IsAnimalFriendly  bool    `json:"is_animal_friendly"`
	TownID            uint    `json:"town_id"`
	UserID            uint    `json:"user_id"`
}

type ActivityEquipment struct {
	Title         string        `json:"title"`
	Description   string        `json:"description"`
	Capacity      int           `json:"capacity"`
	IsRecommended bool          `json:"is_recommended"`
	Discount      float64       `json:"discount"`
	Date          time.Time     `json:"date"`
	Skill         string        `json:"skill_level"`
	Type          string        `json:"activity_type"`
	Price         float64       `json:"price"`
	Duration      time.Duration `json:"duration"`
	TownID        uint          `json:"town_id"`
	UserID        uint          `json:"user_id"`
	Equipment     string        `json:"equipment" type:"text"`
}
