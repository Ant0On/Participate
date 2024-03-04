package models

import (
	"time"

	"gorm.io/gorm"
)

type SkillLevel string
type ActivityType string

const (
	Beginner     SkillLevel = "Beginner"
	Intermediate SkillLevel = "Intermediate"
	Advanced     SkillLevel = "Advanced"
)

const (
	InDoor  ActivityType = "InDoor"
	Outdoor ActivityType = "Outdoor"
)

type Activity struct {
	gorm.Model
	Offer
	Date         time.Time
	Skill        SkillLevel
	Type         ActivityType
	Equipment    []string
	Duration     time.Duration
	Reservations []Reservation
}
