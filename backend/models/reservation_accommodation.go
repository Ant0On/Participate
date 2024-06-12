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
	AccommodationID uint      `gorm:"not null" json:"accommodation_id" binding:"required"`
}

func (r *ReservationAccommodation) validate() error {
	var accommodation Accommodation
	if err := DB.First(&accommodation, r.AccommodationID).Error; err != nil {
		return fmt.Errorf("failed to retrieve accommodation: %w", err)
	}

	if r.NumberOfPeople > accommodation.Capacity {
		return fmt.Errorf("too many people added to reservation")
	}
	return nil
}

func (r *ReservationAccommodation) Save() error {
	if err := r.validate(); err != nil {
		return fmt.Errorf("reservation validation failed: %v", err)
	}

	if err := DB.Create(r).Error; err != nil {
		return fmt.Errorf("failed to save reservation: %w", err)
	}

	return nil
}

func (r *ReservationAccommodation) Update() error {
	if err := r.validate(); err != nil {
		return fmt.Errorf("reservation validation failed: %v", err)
	}
	if err := DB.Save(r).Error; err != nil {
		return fmt.Errorf("failed to update reservation: %w", err)
	}
	return nil
}

func (r *ReservationAccommodation) Delete() error {
	if err := DB.Delete(&r).Error; err != nil {
		return fmt.Errorf("failed to delete reservation: %w", err)
	}
	return nil
}

func GetAccommodationReservationById(id string) (ReservationOperations, error) {
	var r ReservationAccommodation
	if err := DB.First(&r, id).Error; err != nil {
		return nil, fmt.Errorf("reservation not found: %w", err)
	}
	return ReservationOperations(&r), nil
}

func GetAccommodationReservationsByState(state string) ([]ReservationOperations, error) {
	var reservations []ReservationAccommodation
	if err := DB.Model(&ReservationAccommodation{}).Where("reservation_state = ?", state).Scan(reservations).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve reservations: %w", err)
	}
	var ops []ReservationOperations
	for _, r := range reservations {
		ops = append(ops, ReservationOperations(&r))
	}

	return ops, nil
}

func (r *ReservationAccommodation) ChangeState(state string) error {
	r.ReservationState = ReservationState(state)
	return r.Update()
}

func (r *ReservationAccommodation) ChangeCapacity() error {
	return nil
}
