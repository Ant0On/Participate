package models

import (
	"errors"
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

func (t *Town) Save() (bool, error) {
	existingTown := Town{}
	if err := DB.Where("country_id = ? AND name = ?", t.CountryID, t.Name).First(&existingTown).Error; err == nil {
		return false, nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, fmt.Errorf("failed to check existing town: %w", err)
	}

	if err := DB.Create(t).Error; err != nil {
		return true, fmt.Errorf("DB.Create: %w", err)
	}

	return true, nil
}
