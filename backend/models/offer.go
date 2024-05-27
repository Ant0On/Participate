package models

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OfferOperations interface {
	Save() error
	Update() error
	Delete() error
	UpdatePrice(price float64) error
	AddDiscount(discount float64) error
	GetID() (uint, error)
	HandleOfferImageUploads(c *gin.Context, tableName string, id uint) error
}

type Offer struct {
	Title       string  `gorm:"size:100;not null" form:"title" binding:"required,min=2,max=100"`
	Description string  `gorm:"size:300;not null" form:"description" binding:"required,min=30,max=300"`
	Capacity    int     `gorm:"not null" form:"capacity" binding:"required,gt=0"`
	Discount    float64 `gorm:"not null;default: 0.00" form:"discount"`
	TownID      uint    `gorm:"not null" form:"town_id" binding:"required"`
	UserID      uint    `gorm:"not null" form:"user_id" binding:"required"`
	Town        Town    `gorm:"foreignKey:TownID"`
}

func (o *Offer) HandleOfferImageUploads(c *gin.Context, tableName string, offerID uint) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["images"]
	if len(files) == 0 {
		return fmt.Errorf("no images were uploaded for the offer")
	}

	offerFolder := filepath.Join(fmt.Sprintf("images/offers/%s/%d", tableName, offerID))
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
