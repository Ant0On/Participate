package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Offer
	DateFrom     time.Time `gorm:"not null" form:"date_from" binding:"required"`
	DateTo       time.Time `gorm:"not null" form:"date_to" binding:"required"`
	Price        float64   `gorm:"not null" form:"price" binding:"required,gt=0"`
	Type         string    `gorm:"type:text;check:event_type IN ('conference', 'concert', 'festival', 'sports event'); column:event_type; not null" form:"event_type" binding:"required,oneof=conference concert festival 'sports event'"`
	Reservations []ReservationEvent
}

func (e *Event) BeforeDelete(tx *gorm.DB) (err error) {
	if err = tx.Where("event_id = ?", e.ID).Delete(&ReservationEvent{}).Error; err != nil {
		return fmt.Errorf("error deleting associated reservations: %w", err)
	}
	return nil
}

func (e *Event) Save() error {
	if err := DB.Create(e).Error; err != nil {
		return fmt.Errorf("failed to create event: %w", err)
	}
	return nil
}

func (e *Event) GetID() (uint, error) {
	if err := DB.First(&e).Error; err != nil {
		return 0, fmt.Errorf("failed to get event ID: %w", err)
	}
	return e.ID, nil
}

func (e *Event) Update() error {
	if err := DB.Save(e).Error; err != nil {
		return fmt.Errorf("failed to update event: %w", err)
	}
	return nil
}

func (e *Event) Delete() error {
	if err := DB.Delete(e).Error; err != nil {
		return fmt.Errorf("failed to delete event: %w", err)
	}
	return nil
}

func GetEventByID(id string) (OfferOperations, error) {
	var e *Event
	if err := DB.First(&e, id).Error; err != nil {
		return nil, fmt.Errorf("failed to get event by ID: %w", err)
	}
	return OfferOperations(e), nil
}
