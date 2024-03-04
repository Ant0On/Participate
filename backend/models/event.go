package models

import (
	"time"

	"gorm.io/gorm"
)

type EventType string

const (
	Conference  EventType = "Conference"
	Concert     EventType = "Concert"
	Festival    EventType = "Festival"
	SportsEvent EventType = "Sports event"
)

type Event struct {
	gorm.Model
	Offer
	Date         time.Time
	Type         EventType
	Reservations []ReservationEvent
}
