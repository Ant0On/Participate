package models

import (
	"time"

	"gorm.io/gorm"
)

type ReservationActivity struct {
	gorm.Model
	Reservation
	Date       time.Time `gorm:"not null" json:"date" binding:"required"`
	ActivityID uint      `json:"activity_id"`
}
