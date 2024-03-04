package models

import (
	"time"

	"gorm.io/gorm"
)

type ReservationEvent struct {
	gorm.Model
	Reservation
	Date    time.Time
	EventID uint
}
