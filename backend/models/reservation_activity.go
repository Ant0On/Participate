package models

import (
	"time"

	"gorm.io/gorm"
)

type ReservationActivity struct {
	gorm.Model
	Reservation
	Date       time.Time
	ActivityID uint
}
