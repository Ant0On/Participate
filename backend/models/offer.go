package models

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type OfferOperations interface {
	Save() error
	Update() error
	Delete() error
	GetID() (uint, error)
	HandleOfferImageUploads(c *gin.Context, tableName string, id uint) error
}

type Offer struct {
	Title       string  `gorm:"size:100;not null" form:"title" binding:"required,min=2,max=100"`
	Description string  `gorm:"size:300;not null" form:"description" binding:"required,min=30,max=3000"`
	Capacity    int     `gorm:"not null" form:"capacity" binding:"required,gt=0"`
	Discount    float64 `gorm:"not null;default: 0.00" form:"discount"`
	TownID      uint    `gorm:"not null" form:"town_id"`
	UserID      uint    `gorm:"not null" form:"user_id" binding:"required"`
	Town        Town    `gorm:"foreignKey:TownID"`
}

func CheckOffers() error {
	var activities []Activity
	var events []Event

	if err := DB.Model(&Activity{}).Scan(&activities).Error; err != nil {
		return fmt.Errorf("failed to retrieve activities: %w", err)
	}
	if err := DB.Model(&Event{}).Scan(&events).Error; err != nil {
		return fmt.Errorf("failed to retrieve events: %w", err)
	}

	now := time.Now()

	for _, activity := range activities {
		if activity.Date.After(now) && now.Format("2006-01-02") == activity.Date.Format("2006-01-02") {
			if err := activity.Delete(); err != nil {
				return fmt.Errorf("failed to delete activity: %w", err)
			}
		}
	}
	for _, event := range events {
		if event.DateTo.After(now) && now.Format("2006-01-02") == event.DateTo.Format("2006-01-02") {
			if err := event.Delete(); err != nil {
				return fmt.Errorf("failed to delete event: %w", err)
			}
		}
	}
	return nil
}

func (o *Offer) HandleOfferImageUploads(c *gin.Context, tableName string, offerID uint) error {
	form, err := c.MultipartForm()
	if err != nil {
		return fmt.Errorf("failed to parse multipart form: %w", err)
	}

	files := form.File["images"]
	if len(files) == 0 {
		return fmt.Errorf("no images were uploaded for the offer")
	}

	offerFolder := filepath.Join(fmt.Sprintf("images/offers/%s/%d", tableName, offerID))
	if err := os.MkdirAll(offerFolder, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create offer image folder: %w", err)
	}

	for i, file := range files {
		filename := fmt.Sprintf("%s_%d.jpeg", strconv.Itoa(int(offerID)), i)
		dst := filepath.Join(offerFolder, filename)

		if err := c.SaveUploadedFile(file, dst); err != nil {
			return fmt.Errorf("failed to save uploaded file: %w", err)
		}
	}

	return nil
}
