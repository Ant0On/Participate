package models

import "gorm.io/gorm"

type Discount struct {
	gorm.Model
	Type         string  `gorm:"size:50;not null" json:"type"`
	Amount       float64 `gorm:"not null" json:"amount"`
	Reservations []Reservation
}
