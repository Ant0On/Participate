package models

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OfferType string

//const (
//	Activity      OfferType = "activity"
//	Event         OfferType = "event"
//	Accommodation OfferType = "accommodation"
//)

type Offer struct {
	gorm.Model
	Title         string  `gorm:"size:100;not null" form:"name" binding:"required,min=2,max=100"`
	Description   string  `gorm:"size:300;not null" form:"description" binding:"required,min=30,max=300"`
	Price         float64 `gorm:"not null" form:"price" binding:"required,gt=0"`
	Capacity      int     `gorm:"not null" form:"max_people" binding:"required,gt=0"`
	IsRecommended bool    `gorm:"not null" form:"is_recommended"`
	Discount      float64 `gorm:"not null;default: 0.00" form:"discount"`
	TownID        uint    `gorm:"not null" form:"town_id" binding:"required"`
}

func (o *Offer) Save() error {
	if err := DB.Create(&o).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}

func (o *Offer) Delete() error {
	if err := DB.Delete(&o).Error; err != nil {
		return fmt.Errorf("DB.Delete: %w", err)
	}
	return nil
}

func GetOfferByID(id string) (*Offer, error) {
	var o *Offer
	if err := DB.First(&o, id).Error; err != nil {
		return nil, fmt.Errorf("DB.First: %w", err)
	}
	return o, nil
}

func (o *Offer) Update() error {
	if err := DB.Save(&o).Error; err != nil {
		return err
	}
	return nil
}

func (o *Offer) HandleOfferImageUploads(c *gin.Context, offerID uint) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["images"]
	if len(files) == 0 {
		return fmt.Errorf("no images were uploaded for the offer")
	}

	offerFolder := filepath.Join("images/offers", fmt.Sprintf("%d", offerID))
	if err := os.MkdirAll(offerFolder, os.ModePerm); err != nil {
		return err
	}

	for i, file := range files {
		filename := fmt.Sprintf("%s_%d.jpeg", strconv.Itoa(int(offerID)), i)
		dst := filepath.Join(offerFolder, filename)

		if err := c.SaveUploadedFile(file, dst); err != nil {
			return err
		}
	}

	return nil
}

func AddRecommendedOffers() ([]Offer, error) {
	var offers []Offer

	query := DB.Model(&Offer{})

	result := query.
		Joins("INNER JOIN reservation ON offer.id = reservation.offer_id").
		Select("offer.id as ID, offer.name, offer.description, offer.price, offer.max_people, offer.is_animal_friendly," +
			"offer.is_recommended, offer.offer_type, offer.discount, offer.user_id," +
			"offer.town_id, AVG(reservation.grade_id) as avg_grade").
		Group("offer.id").
		Order("avg_grade desc").
		Limit(10).
		Find(&offers)

	if err := result.Error; err != nil {
		return nil, fmt.Errorf("DB.Error: %w", err)
	}

	return offers, nil
}
