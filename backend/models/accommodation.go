package models

import "gorm.io/gorm"

type AccommodationType string

const (
	Hotel      AccommodationType = "Hotel"
	Hostel     AccommodationType = "Hostel"
	Apartment  AccommodationType = "Apartment"
	Villa      AccommodationType = "Villa"
	Guesthouse AccommodationType = "Guesthouse"
)

type Accommodation struct {
	gorm.Model
	Offer
	GeneralFacilities []string
	NumberOfRooms     int
	Type              AccommodationType
	IsAnimalFriendly  bool
	RatingID          Grade
	Rooms             []Room
	Reservations      []ReservationAccommodation
}
