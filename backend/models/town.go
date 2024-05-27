package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Town struct {
	gorm.Model
	Name           string `gorm:"size:50;not null" json:"name" binding:"required,min=2,max=50"`
	Accommodations []Accommodation
	Activities     []Activity
	Events         []Event
	CountryID      uint    `gorm:"not null" json:"country_id" binding:"required"`
	Country        Country `gorm:"foreignKey:CountryID"`
}

func (t *Town) Save() error {
	if err := DB.Create(t).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}
