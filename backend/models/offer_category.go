package models

import "gorm.io/gorm"

type OfferCategory struct {
	gorm.Model
	Name   string `gorm:"size:50;not null;unique" json:"name"`
	Offers []Offer
}
