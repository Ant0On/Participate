package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	Name    string `gorm:"size:30;not null" json:"name" binding:"required"`
	Data    []byte `json:"data" binding:"required"`
	OfferId uint   `gorm:"not null"`
}

func Save(name string, data []byte) error {
	DB.Create(&Image{Name: name, Data: data})

	return nil
}

func GetImage() (*Image, error) {
	var image *Image
	if err := DB.Model(&Image{}).First(&image).Error; err != nil {
		return nil, err
	}
	return image, nil
}
