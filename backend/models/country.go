package models

import "gorm.io/gorm"

type Country struct {
	gorm.Model
	Name string `gorm:"size:40;not null;unique" json:"name"`
	Code string `gorm:"size:2;not null;unique" json:"code"`
}
