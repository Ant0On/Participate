package models

import "gorm.io/gorm"

type Grade struct {
	gorm.Model
	Count        int    `gorm:"not null" json:"count"`
	Description  string `gorm:"size:50;not null;unique" json:"description"`
	Reservations []Reservation
}
