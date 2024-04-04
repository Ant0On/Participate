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

func (r *ReservationAccommodation) Delete() error {
	if err := DB.Delete(&r).Error; err != nil {
		return fmt.Errorf("DB.Delete: %w", err)
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
		return nil, fmt.Errorf("reservation not found: %w", err)
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
