package models

import (
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OfferType string

const (
	Activity      OfferType = "activity"
	Event         OfferType = "event"
	Accommodation OfferType = "accommodation"
)

type Offer struct {
	gorm.Model
	Name             string        `gorm:"size:100;not null" form:"name" binding:"required,min=2,max=100"`
	Description      string        `gorm:"size:300;not null" form:"description" binding:"required,min=30,max=300"`
	Price            float64       `gorm:"not null" form:"price" binding:"required,gt=0"`
	MaxPeople        int           `gorm:"not null" form:"max_people" binding:"required,gt=0"`
	IsAnimalFriendly bool          `gorm:"not null" form:"is_animal_friendly"`
	IsRecommended    bool          `gorm:"not null" form:"is_recommended"`
	OfferType        OfferType     `gorm:"type:varchar(255);check:offer_type IN ('activity', 'event', 'accommodation'); column:offer_type; not null" form:"offer_type" binding:"required,oneof=activity event accommodation"`
	Discount         float64       `gorm:"not null;default: 0.00" form:"discount" binding:"gte=0,lte=100"`
	HostID           uint          `gorm:"not null" form:"host_id" binding:"required"`
	TownID           uint          `gorm:"not null" form:"town_id" binding:"required"`
	Reservations     []Reservation `gorm:"foreignKey:OfferID"`
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

func GetOffersForHost(hostID string) ([]Offer, error) {
	var offers []Offer
	if err := DB.Find(&offers).Where("host_id = ?", hostID).Error; err != nil {
		return nil, fmt.Errorf("DB.Find: %w", err)
	}
	return offers, nil
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

	files := form.File["image"]
	if len(files) == 0 {
		return fmt.Errorf("no image was uploaded for the offer")
	}

	file := files[0]

	filename := fmt.Sprintf("%d.jpeg", offerID)
	dst := filepath.Join("images/offers", filename)

	if err := c.SaveUploadedFile(file, dst); err != nil {
		return err
	}

	return nil
}
