package DTO

import "backend/models"

type EventWithLocation struct {
	OfferWithLocation
	Price float64          `json:"price" binding:"required,gt=0"`
	Type  models.EventType `json:"type" binding:"required,oneof=conference concert festival 'sports event'"`
}
