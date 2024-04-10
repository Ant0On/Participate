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
	NumberOfRooms     int               `gorm:"not null" form:"number_of_rooms" binding:"required,min=1"`
	Type              AccommodationType `gorm:"type:varchar(255);check:accommodation_type IN ('hotel', 'hostel', 'apartment', 'villa', 'guesthouse'); column:accommodation_type; not null" form:"type" binding:"required,oneof=hotel hostel apartment villa guesthouse"`
	IsAnimalFriendly  bool              `gorm:"not null" form:"is_animal_friendly"`
	PricePerDay       float64           `gorm:"not null" form:"price_per_day" binding:"required,min=1"`
	TownID            uint              `gorm:"not null" form:"town_id" binding:"required"`
	UserID            uint              `gorm:"not null" form:"user_id" binding:"required"`
	GeneralFacilities []GeneralFacility `gorm:"many2many:accommodation_general_facilities;"`
	Rooms             []Room
	Reservations      []ReservationAccommodation
}

func (a *Accommodation) Save() error {
	if err := DB.Create(&a).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}

func (a *Accommodation) GetID() (uint, error) {
	if err := DB.First(&a).Error; err != nil {
		return 0, fmt.Errorf("DB.First: %w", err)
	}
	return a.ID, nil
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

func (a *Accommodation) UpdatePrice(price float64) error {
	a.PricePerDay = price
	return a.Update()
}

func (a *Accommodation) AddDiscount(discount float64) error {
	a.Discount = discount
	return a.Update()
}

func GetAccommodationByID(id string) (OfferOperations, error) {
	var a Accommodation
	if err := DB.First(&a, id).Error; err != nil {
		return nil, fmt.Errorf("DB.First: %w", err)
	}
	return OfferOperations(&a), nil
}
