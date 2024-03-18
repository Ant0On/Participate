package DTO

type OfferWithLocation struct {
	OfferID       uint    `json:"offer_id" binding:"required"`
	Title         string  `json:"title" binding:"required,min=2"`
	Description   string  `json:"description" binding:"required,min=30"`
	Capacity      int     `json:"capacity" binding:"required,gt=0"`
	IsRecommended bool    `json:"is_recommended"`
	TownName      string  `json:"town_name" binding:"required,min=2"`
	CountryName   string  `json:"country_name" binding:"required,min=3"`
	UserID        uint    `json:"user_id" binding:"required"`
	Discount      float64 `json:"discount" binding:"required,min=0,max=100"`
}
