package models

import "gorm.io/gorm"

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
