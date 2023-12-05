package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Animal struct {
	gorm.Model
	Name         string `gorm:"size:30;not null" json:"name"`
	Reservations []Reservation
}

func (a *Animal) Save() error {
	if err := DB.Create(&a).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}
