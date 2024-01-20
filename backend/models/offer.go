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
	Name             string    `gorm:"size:100;not null" form:"name"`
	Description      string    `gorm:"size:300;not null" form:"description"`
	ImageFilenames   []string  `gorm:"-" form:"image_filenames"`
	Price            float64   `gorm:"not null" form:"price"`
	MaxPeople        int       `gorm:"not null" form:"max_people"`
	IsAnimalFriendly bool      `gorm:"not null" form:"is_animal_friendly"`
	IsRecommended    bool      `gorm:"not null" form:"is_recommended"`
	OfferType        OfferType `gorm:"type:varchar(255);check:offer_type IN ('activity', 'event', 'place'); column:offer_type; not null" form:"offer_type"`
	HostID           uint      `gorm:"not null" form:"host_id"`
	TownID           uint      `gorm:"not null" form:"town_id"`
	Reservations     []Reservation
	ID               string `gorm:"type:uuid;default:uuid_generate_v4()" form:"-"`
}

func (o *Offer) Save() error {
	if err := DB.Create(&o).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}

func (o *Offer) Delete() error {
	if err := DB.Delete(&o).Error; err != nil {
		return fmt.Errorf("DB.Delete: %w", err)
	}
	return nil
}

func GetOfferByID(id string) (*Offer, error) {
	var o *Offer
	if err := DB.First(&o, id).Error; err != nil {
		return nil, fmt.Errorf("DB.First: %w", err)
	}
	return o, nil
}

func (o *Offer) Update() error {
	if err := DB.Save(&o).Error; err != nil {
		return err
	}
	return nil
}
