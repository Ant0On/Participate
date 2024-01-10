package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	Name    string `gorm:"size:30;not null" json:"name" binding:"required"`
	Data    []byte `json:"data" binding:"required"`
	OfferID uint   `gorm:"not null"`
}

func Save(name string, data []byte) error {
	DB.Create(&Image{Name: name, Data: data})

	return nil
}

func GetAllImages() ([]Image, error) {
	var images []Image
	if err := DB.Model([]Image{}).Scan(&images).Error; err != nil {
		return nil, fmt.Errorf("DB.Model([]Image).Scan: %w", err)
	}
	return images, nil
}

func GetImage() (*Image, error) {
	var image *Image
	if err := DB.Model(&Image{}).First(&image).Error; err != nil {
		return nil, err
	}
	return image, nil
}
