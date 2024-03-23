package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type ReservationAccommodation struct {
	gorm.Model
	Reservation
	DateFrom        time.Time `gorm:"not null" json:"date_from" binding:"required"`
	DateTo          time.Time `gorm:"not null" json:"date_to" binding:"required,gtfield=DateFrom"`
	RatingID        uint      `json:"rating_id"`
	AnimalID        uint      `json:"animal_id"`
	AccommodationID uint      `json:"accommodation_id"`
}

func (r *ReservationAccommodation) ValidateDates() error {
	if r.DateFrom.Before(time.Now()) || r.DateTo.Before(time.Now()) {
		return fmt.Errorf("reservation dates cannot be in the past")
	}

	return nil
}
func (r *ReservationAccommodation) Save() error {
	if err := r.Validate(); err != nil {
		return fmt.Errorf("r.Validate: %v", err)
	}

	if err := DB.Create(r).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}

func (r *ReservationAccommodation) Update() error {
	if err := r.Validate(); err != nil {
		return fmt.Errorf("r.Validate: %v", err)
	}
	if err := DB.Save(r).Error; err != nil {
		return err
	}
	return nil
}

func GetAccommodationReservationById(id string) (*ReservationAccommodation, error) {
	var r ReservationAccommodation
	if err := DB.Model(&ReservationAccommodation{}).Where("id = ?", id).Scan(&r).Error; err != nil {
		return nil, fmt.Errorf("reservation not found: %w", err)
	}
	return &r, nil
}

func GetAccommodationReservationsByState(state string) ([]ReservationAccommodation, error) {
	var reservations []ReservationAccommodation
	if err := DB.Model(&ReservationAccommodation{}).Where("reservation_state = ?", state).Scan(reservations).Error; err != nil {
		return nil, fmt.Errorf("reservation not found: %w", err)
	}
	return reservations, nil
}
