package DTO

import "backend/models"

type OfferWithLocation struct {
	OfferID          uint             `json:"offer_id"`
	Name             string           `json:"name"`
	Description      string           `json:"description"`
	Price            float64          `json:"price"`
	MaxPeople        int              `json:"max_people"`
	IsAnimalFriendly bool             `json:"is_animal_friendly"`
	IsRecommended    bool             `json:"is_recommended"`
	OfferType        models.OfferType `json:"offer_type"`
	TownName         string           `json:"town_name"`
	CountryName      string           `json:"country_name"`
}
