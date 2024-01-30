package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Discount struct {
	gorm.Model
	Amount float64 `gorm:"not null" json:"amount"`
	Offers []Offer
}

func (d *Discount) Save() error {
	if err := DB.Create(&d).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}
