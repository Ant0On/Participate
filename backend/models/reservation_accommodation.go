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
	GradeID         uint      `json:"grade_id"`
	AnimalID        uint      `json:"animal_id"`
	AccommodationID uint      `json:"accommodation_id"`
}

func (r *ReservationAccommodation) ValidateDates() error {
	if r.DateFrom.Before(time.Now()) || r.DateTo.Before(time.Now()) {
		return fmt.Errorf("reservation dates cannot be in the past")
	}

	return nil
}
