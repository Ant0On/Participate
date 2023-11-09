package models

import (
	"fmt"

	"gorm.io/gorm"
)

type OfferType string

const (
	Activity OfferType = "activity"
	Event    OfferType = "event"
	Place    OfferType = "place"
)

type Offer struct {
	gorm.Model
	Name             string    `gorm:"size:100;not null" json:"name"`
	Description      string    `gorm:"size:300;not null" json:"description"`
	Price            float64   `gorm:"not null" json:"price"`
	MaxPeople        int       `gorm:"not null" json:"max_people"`
	IsAnimalFriendly bool      `gorm:"not null" json:"is_animal_friendly"`
	IsRecommended    bool      `gorm:"not null" json:"is_recommended"`
	OfferType        OfferType `gorm:"type:varchar(255);check:offer_type IN ('activity', 'event', 'place'); column:offer_type; not null" json:"offer_type"`
	OfferCategoryID  uint      `gorm:"not null"`
	HostID           uint      `gorm:"not null"`
	TownID           uint      `gorm:"not null"`
	Reservations     []Reservation
}

func (o *Offer) Save() error {
	if err := DB.Create(&o).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}
