package models

import (
	"fmt"

	"gorm.io/gorm"
)

type RoomFacility struct {
	gorm.Model
	Name string
	Room []Room `gorm:"many2many:room_room_facilities;"`
}

var RoomFacilitiesList = []RoomFacility{
	{Name: "Television"},
	{Name: "Air Conditioning"},
	{Name: "Mini Fridge"},
	{Name: "Safe"},
	{Name: "Coffee Maker"},
	{Name: "Microwave"},
	{Name: "Kettle"},
	{Name: "Iron and Ironing Board"},
	{Name: "Hair Dryer"},
	{Name: "Desk"},
	{Name: "Ocean View"},
	{Name: "Mountain View"},
	{Name: "Balcony"},
	{Name: "Bathtub"},
	{Name: "Shower"},
	{Name: "WiFi"},
	{Name: "Room Service"},
	{Name: "Breakfast Included"},
	{Name: "In-Room Safe"},
	{Name: "Telephone"},
	{Name: "DVD Player"},
	{Name: "Alarm Clock"},
	{Name: "Robes"},
	{Name: "Slippers"},
	{Name: "Toiletries"},
	{Name: "Work Desk"},
	{Name: "Sofa Bed"},
	{Name: "Fireplace"},
	{Name: "Refrigerator"},
	{Name: "Dining Area"},
}

func (rf *RoomFacility) save() error {
	if err := DB.Create(&rf).Error; err != nil {
		return fmt.Errorf("failed to save room facility: %w", err)
	}
	return nil
}

func AddRoomFacilities() error {
	for _, facility := range RoomFacilitiesList {
		if err := facility.save(); err != nil {
			return fmt.Errorf("failed to add room facility: %w", err)
		}
	}
	return nil
}

func GetAllRoomFacilities() ([]RoomFacility, error) {
	var roomFacilities []RoomFacility

	if err := DB.Order("name").Find(&roomFacilities).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve all room facilities: %w", err)
	}
	return roomFacilities, nil
}
