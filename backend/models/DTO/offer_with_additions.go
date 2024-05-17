package DTO

import "time"

type AccommodationRoom struct {
	Title             string  `gorm:"size:100;not null" json:"title" binding:"required,min=2,max=100"`
	PricePerDay       float64 `gorm:"not null" json:"price_per_day" binding:"required,min=1"`
	RoomName          string  `gorm:"not null" json:"room_name" binding:"required,min=3"`
	RoomDescription   string  `gorm:"not null" json:"room_description" binding:"required,min=10"`
	Capacity          int     `gorm:"not null" json:"room_capacity" binding:"required,gt=0"`
	Area              int     `gorm:"not null" json:"room_area" binding:"required,gt=0"`
	RoomFacilities    string  `gorm:"not null" json:"room_facilities" binding:"required" type:"text"`
	GeneralFacilities string  `gorm:"not null" json:"general_facilities" binding:"required" type:"text"`
	AccommodationType string  `gorm:"not null" json:"accommodation_type" binding:"required,oneof=hotel hostel apartment villa guesthouse"`
	IsAnimalFriendly  bool    `gorm:"not null" json:"is_animal_friendly"`
	TownID            uint    `gorm:"not null" json:"town_id" binding:"required"`
	UserID            uint    `gorm:"not null" json:"user_id" binding:"required"`
}

type ActivityEquipment struct {
	Title         string        `gorm:"size:100;not null" json:"title" binding:"required,min=2,max=100"`
	Description   string        `gorm:"size:300;not null" json:"description" binding:"required,min=30,max=300"`
	Capacity      int           `gorm:"not null" json:"capacity" binding:"required,gt=0"`
	IsRecommended bool          `gorm:"not null" json:"is_recommended"`
	Discount      float64       `gorm:"not null;default: 0.00" json:"discount"`
	Date          time.Time     `gorm:"not null" json:"date" binding:"required"`
	Skill         string        `gorm:"type:varchar(255);not null" json:"skill_level" binding:"required,oneof=beginner intermediate advanced"`
	Type          string        `gorm:"type:varchar(255);not null" json:"activity_type" binding:"required,oneof=indoor outdoor"`
	Price         float64       `gorm:"not null" json:"price" binding:"required,gt=0"`
	Duration      time.Duration `gorm:"not null" json:"duration" binding:"required"`
	TownID        uint          `gorm:"not null" json:"town_id" binding:"required"`
	UserID        uint          `gorm:"not null" json:"user_id" binding:"required"`
	Equipment     string        `gorm:"not null" json:"equipment" binding:"required" type:"text"`
}
