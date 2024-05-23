package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type EventType string

const (
	Conference  EventType = "conference"
	Concert     EventType = "concert"
	Festival    EventType = "festival"
	SportsEvent EventType = "sports event"
)

type Event struct {
	gorm.Model
	Offer
	DateFrom     time.Time `gorm:"not null" form:"date_from" binding:"required"`
	DateTo       time.Time `gorm:"not null" form:"date_to" binding:"required"`
	Price        float64   `gorm:"not null" form:"price" binding:"required,gt=0"`
	Type         EventType `gorm:"type:varchar(255);check:event_type IN ('conference', 'concert', 'festival', 'sports event'); column:event_type; not null" form:"event_type" binding:"required,oneof=conference concert festival 'sports event'"`
	Reservations []ReservationEvent
}

func (e *Event) Save() error {
	if err := DB.Create(e).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}

func (e *Event) GetID() (uint, error) {
	if err := DB.First(&e).Error; err != nil {
		return 0, fmt.Errorf("DB.First: %w", err)
	}
	return e.ID, nil
}

func SearchEvents(title, location string) ([]Event, error) {
	var events []Event
	query := DB.Model(&Event{})

	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}
	if location != "" {
		query = query.Where("location LIKE ?", "%"+location+"%")
	}

	return events, nil
}

func (e *Event) Update() error {
	if err := DB.Save(e).Error; err != nil {
		return err
	}
	return nil
}

func (e *Event) Delete() error {
	if err := DB.Delete(e).Error; err != nil {
		return fmt.Errorf("DB.Delete: %w", err)
	}
	return nil
}

func (e *Event) UpdatePrice(price float64) error {
	e.Price = price
	return e.Update()
}

func (e *Event) AddDiscount(discount float64) error {
	e.Discount = discount
	return e.Update()
}

func GetEventByID(id string) (OfferOperations, error) {
	var e Event
	if err := DB.First(e, id).Error; err != nil {
		return nil, fmt.Errorf("DB.First: %w", err)
	}
	return OfferOperations(&e), nil
}
