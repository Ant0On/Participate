package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Town struct {
	gorm.Model
	Name       string `gorm:"size:50;not null" json:"name"`
	Offers     []Offer
	CountryID  uint `gorm:"not null" json:"country_id"`
	TownTypeID uint `json:"town_type_id"`
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
