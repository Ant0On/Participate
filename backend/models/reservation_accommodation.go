package models

import (
	"time"

	"gorm.io/gorm"
)

type ReservationAccommodation struct {
	gorm.Model
	Reservation
	DateFrom        time.Time `gorm:"not null" json:"date_from" binding:"required"`
	DateTo          time.Time `gorm:"not null" json:"date_to" binding:"required,gtfield=DateFrom"`
	GradeID         uint      `json:"grade_id"`
	AnimalID        uint
	AccommodationID uint
}
