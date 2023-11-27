package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Town struct {
	gorm.Model
	Name       string `gorm:"size:50;not null" json:"name"`
	Offers     []Offer
	CountryID  uint `gorm:"not null"`
	TownTypeID uint `gorm:"not null"`
}

func (t *Town) Save() error {
	if err := DB.Create(&t).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}
