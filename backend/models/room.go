package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	RoomNumber      int            `gorm:"not null" json:"number" binding:"required,gt=0"`
	RoomName        string         `gorm:"not null" json:"name" binding:"required,min=3"`
	RoomDescription string         `gorm:"not null" json:"description" binding:"required,min=10"`
	Capacity        int            `gorm:"not null" json:"capacity" binding:"required,gt=0"`
	Area            int            `gorm:"not null" json:"area" binding:"required,gt=0"`
	AccommodationID uint           `gorm:"not null" json:"accommodation_id" binding:"required"`
	RoomFacilities  []RoomFacility `gorm:"many2many:room_room_facilities;"`
}

func (r *Room) Save() error {
	if err := DB.Create(&r).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}
func (r *Room) Update() error {
	if err := DB.Save(&r).Error; err != nil {
		return err
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

func (r *Room) AddFacilities(facilities []RoomFacility) error {
	r.RoomFacilities = facilities
	return r.Update()
}
