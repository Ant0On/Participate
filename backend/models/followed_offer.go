package models

import "gorm.io/gorm"

type FollowedOffer struct {
	gorm.Model
	CustomerID uint `gorm:"not null"`
	OfferID    uint `gorm:"not null"`
}
