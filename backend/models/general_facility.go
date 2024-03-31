package models

import (
	"fmt"

	"gorm.io/gorm"
)

type GeneralFacility struct {
	gorm.Model
	Name            string
	AccommodationID uint
}

var GeneralFacilitiesList = []GeneralFacility{
	{Name: "Swimming Pool"},
	{Name: "Gym"},
	{Name: "Spa"},
	{Name: "Restaurant"},
	{Name: "Bar"},
	{Name: "Lounge"},
	{Name: "Conference Room"},
	{Name: "Business Center"},
	{Name: "WiFi"},
	{Name: "Parking"},
	{Name: "24-Hour Front Desk"},
	{Name: "Fitness Center"},
	{Name: "Laundry Service"},
	{Name: "Room Service"},
	{Name: "Concierge Service"},
	{Name: "Outdoor Pool"},
	{Name: "Children's Playground"},
	{Name: "Tennis Court"},
	{Name: "Library"},
	{Name: "Garden"},
	{Name: "Sauna"},
	{Name: "Jacuzzi"},
	{Name: "Billiards Room"},
	{Name: "Cinema"},
	{Name: "Karaoke Room"},
	{Name: "Bowling Alley"},
	{Name: "BBQ Area"},
	{Name: "Shuttle Service"},
}

func (gf *GeneralFacility) save() error {
	if err := DB.Create(&gf).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}
	return nil
}

func AddGeneralFacilities() error {
	for _, facility := range GeneralFacilitiesList {
		if err := facility.save(); err != nil {
			return fmt.Errorf("country.Save: %w", err)
		}
	}
	return nil
}

func GetAllGeneralFacilities() ([]GeneralFacility, error) {
	var generalFacilities []GeneralFacility

	if err := DB.Order("name").Find(&generalFacilities).Error; err != nil {
		return nil, fmt.Errorf("DB.Order().Find(): %w", err)
	}
	return generalFacilities, nil
}
