package DTO

import (
	"time"

	"backend/models"
)

type ActivityWithLocation struct {
	OfferWithLocation
	Price    float64             `json:"price" binding:"required,gt=0"`
	Skill    models.SkillLevel   `json:"skill_level" binding:"required,oneof=beginner intermediate advanced"`
	Type     models.ActivityType `json:"type" binding:"required,oneof=indoor outdoor"`
	Duration time.Duration       `json:"duration" binding:"required"`
}
