package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	RoomNumber      int
	RoomName        string
	RoomDescription string
	Capacity        int
	Area            int
	RoomFacilities  []string
	AccommodationID uint
}
