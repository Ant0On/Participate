package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	RoomNumber      int
	RoomName        string
	RoomDescription string
	Capacity        int
	Area            int
	AccommodationID uint
	RoomFacilities  []RoomFacility `gorm:"many2many:room_room_facilities;"`
}

func (r *Room) Save() error {
	if err := DB.Create(&r).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}

func GetRoomByID(id string) (*Room, error) {
	var r Room
	if err := DB.First(&r, id).Error; err != nil {
		return nil, fmt.Errorf("DB.First: %w", err)
	}
	return &r, nil
}

func (r *Room) Delete() error {
	if err := DB.Delete(&r).Error; err != nil {
		return fmt.Errorf("DB.Delete: %w", err)
	}
	return nil
}
