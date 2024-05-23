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

func (a *Accommodation) Search(title, location string) ([]interface{}, error) {
	var accommodations []Accommodation
	query := DB.Model(&Accommodation{})

	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}
	if location != "" {
		query = query.Where("location LIKE ?", "%"+location+"%")
	}

	if err := query.Find(&accommodations).Error; err != nil {
		return nil, fmt.Errorf("DB.Find: %w", err)
	}

	results := make([]interface{}, len(accommodations))
	for i, v := range accommodations {
		results[i] = v
	}

	return results, nil
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

func GetAccommodationById(id string) (*Accommodation, error) {
	var a Accommodation
	if err := DB.First(&a, id).Error; err != nil {
		return nil, fmt.Errorf("DB.First: %w", err)
	}
	return &a, nil
}

func (a *Accommodation) AddFacilities(facilities []GeneralFacility) error {
	a.GeneralFacilities = facilities
	return a.Update()
}
