package models

import "gorm.io/gorm"

type Place struct {
	gorm.Model
	Name             string  `gorm:"size:100;not null" json:"name"`
	Price            float64 `gorm:"not null" json:"price"`
	Description      string  `gorm:"size:300;not null" json:"description"`
	MaxPeople        int     `gorm:"not null" json:"max_people"`
	IsAnimalFriendly bool    `gorm:"not null" json:"is_animal_friendly"`
	IsRecommended    bool    `gorm:"not null" json:"is_recommended"`
}
