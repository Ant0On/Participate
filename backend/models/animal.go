package models

import "gorm.io/gorm"

type Animal struct {
	gorm.Model
	Name         string `gorm:"size:30;not null" json:"name"`
	Size         string `gorm:"size:20;not null" json:"size"`
	Reservations []Reservation
}
