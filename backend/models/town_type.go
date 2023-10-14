package models

import "gorm.io/gorm"

type TownType struct {
	gorm.Model
	Name          string `gorm:"size:20;not null" json:"name"`
	MinPopulation int    `gorm:"not null" json:"min_population"`
	MaxPopulation int    `gorm:"not null" json:"max_population"`
}
