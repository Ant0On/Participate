package models

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	DateFrom time.Time `gorm:"not null" json:"date_from"`
	DateTo   time.Time `gorm:"not null" json:"date_to"`
}
