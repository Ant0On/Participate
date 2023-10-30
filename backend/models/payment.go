package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	Type         string `gorm:"size:30;not null" json:"type"`
	Reservations []Reservation
}
