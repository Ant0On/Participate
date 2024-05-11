package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type SkillLevel string
type ActivityType string

const (
	Beginner     SkillLevel = "beginner"
	Intermediate SkillLevel = "intermediate"
	Advanced     SkillLevel = "advanced"
)

const (
	Indoor  ActivityType = "indoor"
	Outdoor ActivityType = "outdoor"
)

type Activity struct {
	gorm.Model
	Offer
	Date         time.Time     `gorm:"not null" json:"date" binding:"required"`
	Skill        SkillLevel    `gorm:"type:varchar(255);check:skill_level IN ('beginner', 'intermediate', 'advanced'); column:skill_level; not null" form:"skill_level" binding:"required,oneof=beginner intermediate advanced"`
	Type         ActivityType  `gorm:"type:varchar(255);check:activity_type IN ('indoor', 'outdoor'); column:activity_type; not null" form:"activity_type" binding:"required,oneof=indoor outdoor"`
	Price        float64       `gorm:"not null" form:"price" binding:"required,gt=0"`
	Duration     time.Duration `gorm:"not null" form:"duration" binding:"required"`
	Equipment    []Equipment   `gorm:"many2many:activity_equipment;"`
	Reservations []ReservationActivity
}

func (a *Activity) Save() error {
	a.Duration.Hours()
	if err := DB.Create(a).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}

func (a *Activity) GetID() (uint, error) {
	if err := DB.First(&a).Error; err != nil {
		return 0, fmt.Errorf("DB.First: %w", err)
	}
	return a.ID, nil
}

func (a *Activity) Update() error {
	if err := DB.Save(a).Error; err != nil {
		return err
	}
	return nil
}

func (a *Activity) Delete() error {
	if err := DB.Delete(a).Error; err != nil {
		return fmt.Errorf("DB.Delete: %w", err)
	}
	return nil
}

func (a *Activity) UpdatePrice(price float64) error {
	a.Price = price
	return a.Update()
}

func (a *Activity) AddDiscount(discount float64) error {
	a.Discount = discount
	return a.Update()
}

func GetActivityByID(id string) (OfferOperations, error) {
	var a Activity
	if err := DB.First(&a, id).Error; err != nil {
		return nil, fmt.Errorf("DB.First: %w", err)
	}
	return OfferOperations(&a), nil
}
