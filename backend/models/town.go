package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Town struct {
	gorm.Model
	Name      string  `gorm:"size:50;not null" json:"name" binding:"required,min=2,max=50"`
	Offers    []Offer `gorm:"foreignKey:OfferID"`
	CountryID uint    `gorm:"not null" json:"country_id" binding:"required"`
}

func (t *Town) Save() error {
	if err := DB.Create(t).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}

func GetTown(town *Town, id string) (*Town, error) {
	if err := DB.First(town, id).Error; err != nil {
		return nil, fmt.Errorf("DB.First: %w", err)
	}

	return town, nil
}
