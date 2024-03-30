package models

import (
	"fmt"

	"gorm.io/gorm"
)

type AccommodationType string

const (
	Hotel      AccommodationType = "hotel"
	Hostel     AccommodationType = "hostel"
	Apartment  AccommodationType = "apartment"
	Villa      AccommodationType = "villa"
	Guesthouse AccommodationType = "guesthouse"
)

type Accommodation struct {
	gorm.Model
	Offer
	GeneralFacilities []string          `gorm:"not null" form:"general_facilities" binding:"required"`
	NumberOfRooms     int               `gorm:"not null" form:"number_of_rooms" binding:"required,min=1"`
	Type              AccommodationType `gorm:"type:varchar(255);check:accommodation_type IN ('hotel', 'hostel', 'apartment', 'villa', 'guesthouse'); column:accommodation_type; not null" form:"accommodation_type" binding:"required,oneof=hotel hostel apartment villa guesthouse"`
	IsAnimalFriendly  bool              `gorm:"not null" form:"is_animal_friendly"`
	PricePerDay       float64           `gorm:"not null" form:"price_per_day" binding:"required,min=1"`
	TownID            uint              `gorm:"not null" form:"town_id" binding:"required"`
	UserID            uint              `gorm:"not null" form:"user_id" binding:"required"`
	Rooms             []Room
	Reservations      []ReservationAccommodation
}

func (a *Accommodation) Save() error {
	if err := DB.Create(&a).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}

func (a *Accommodation) Update() error {
	if err := DB.Save(&a).Error; err != nil {
		return err
	}
	return nil
}

func (a *Accommodation) Delete() error {
	if err := DB.Delete(&a).Error; err != nil {
		return fmt.Errorf("DB.Delete: %w", err)
	}
	return nil
}

func GetAccommodationByID(id string) (*Accommodation, error) {
	var a *Accommodation
	if err := DB.First(&a, id).Error; err != nil {
		return nil, fmt.Errorf("DB.First: %w", err)
	}
	return a, nil
}
