package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Offer
	Date         time.Time     `gorm:"not null" form:"date" binding:"required"`
	Skill        string        `gorm:"type:text;check:skill_level IN ('beginner', 'intermediate', 'advanced'); column:skill_level; not null" form:"skill_level" binding:"required,oneof=beginner intermediate advanced"`
	Type         string        `gorm:"type:text;check:activity_type IN ('indoor', 'outdoor'); column:activity_type; not null" form:"activity_type" binding:"required,oneof=indoor outdoor"`
	Price        float64       `gorm:"not null" form:"price" binding:"required,gt=0"`
	RatingAvg    float64       `gorm:"not null;default: 0.00" form:"-"`
	RatingCount  int           `gorm:"not null;default: 0" form:"-"`
	Duration     time.Duration `gorm:"not null" form:"duration" binding:"required"`
	Equipment    []Equipment   `gorm:"many2many:activity_equipment;"`
	Reservations []ReservationActivity
}

func (a *Activity) BeforeDelete(tx *gorm.DB) (err error) {
	if err = tx.Where("activity_id = ?", a.ID).Delete(&ReservationActivity{}).Error; err != nil {
		return fmt.Errorf("failed to delete associated reservations: %w", err)
	}
	return nil
}

func (a *Activity) Save() error {
	a.Duration = a.Duration / time.Hour
	if err := DB.Create(a).Error; err != nil {
		return fmt.Errorf("failed to create activity: %w", err)
	}
	return nil
}

func (a *Activity) GetID() (uint, error) {
	if err := DB.First(&a).Error; err != nil {
		return 0, fmt.Errorf("failed to retrieve activity ID: %w", err)
	}
	return a.ID, nil
}

func (a *Activity) Update() error {
	if err := DB.Save(a).Error; err != nil {
		return fmt.Errorf("failed to update activity: %w", err)
	}
	return nil
}

func (a *Activity) UpdateRating(rating int) error {
	currentSum := a.RatingAvg * float64(a.RatingCount)
	currentSum += float64(rating)
	a.RatingCount += 1
	a.RatingAvg = currentSum / float64(a.RatingCount)
	if err := DB.Save(a).Error; err != nil {
		return fmt.Errorf("failed to update activity rating: %w", err)
	}
	return nil
}

func (a *Activity) Delete() error {
	if err := DB.Delete(a).Error; err != nil {
		return fmt.Errorf("failed to delete activity: %w", err)
	}
	return nil
}

func GetActivityByID(id string) (OfferOperations, error) {
	var a *Activity
	if err := DB.First(&a, id).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve activity by ID: %w", err)
	}
	return OfferOperations(a), nil
}

func GetActivityById(id string) (*Activity, error) {
	var a Activity
	if err := DB.First(&a, id).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve activity by ID: %w", err)
	}
	return &a, nil
}

func (a *Activity) AddEquipment(equipment []Equipment) error {
	a.Equipment = equipment
	return a.Update()
}
