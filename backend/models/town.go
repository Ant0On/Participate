package models

import "gorm.io/gorm"

type Town struct {
	gorm.Model
	Name string `gorm:"size:50;not null" json:"name"`
}
