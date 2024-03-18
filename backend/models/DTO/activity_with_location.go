package DTO

import (
	"time"

	"backend/models"
)

type ActivityWithLocation struct {
	OfferWithLocation
	Skill    models.SkillLevel   `json:"skill_level" binding:"required,oneof=beginner intermediate advanced"`
	Type     models.ActivityType `json:"activity_type" binding:"required,oneof=indoor outdoor"`
	Duration time.Duration       `json:"duration" binding:"required"`
}
