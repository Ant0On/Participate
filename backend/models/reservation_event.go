package models

import (
	"time"

	"gorm.io/gorm"
)

type ReservationEvent struct {
	gorm.Model
	Reservation
	Date    time.Time `gorm:"not null" json:"date" binding:"required"`
	EventID uint      `json:"event_id"`
}
