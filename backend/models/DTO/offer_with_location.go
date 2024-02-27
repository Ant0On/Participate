package DTO

import "backend/models"

type OfferWithLocation struct {
	OfferID          uint             `json:"offer_id" binding:"required"`
	Name             string           `json:"name" binding:"required,min=2"`
	Description      string           `json:"description" binding:"required,min=30"`
	Price            float64          `json:"price" binding:"required,gt=0"`
	MaxPeople        int              `json:"max_people" binding:"required,gt=0"`
	IsAnimalFriendly bool             `json:"is_animal_friendly"`
	IsRecommended    bool             `json:"is_recommended"`
	OfferType        models.OfferType `json:"offer_type" binding:"required,oneof=activity event accommodation"`
	TownName         string           `json:"town_name" binding:"required,min=2"`
	CountryName      string           `json:"country_name" binding:"required,min=3"`
	AppUserID        uint             `json:"app_user_id" binding:"required"`
	Discount         float64          `json:"discount" binding:"required,min=0,max=100"`
}
