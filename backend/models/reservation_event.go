package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type ReservationEvent struct {
	gorm.Model
	Reservation
	Date    time.Time `gorm:"not null" json:"date" binding:"required"`
	EventID uint      `json:"event_id"`
}

func (r *ReservationEvent) Save() error {
	if err := r.Validate(); err != nil {
		return fmt.Errorf("r.Validate: %v", err)
	}

	if err := DB.Create(&r).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}

func (r *ReservationEvent) Update() error {
	if err := r.Validate(); err != nil {
		return fmt.Errorf("r.Validate: %v", err)
	}
	if err := DB.Save(&r).Error; err != nil {
		return err
	}
	return nil
}

func GetEventReservationById(id string) (*ReservationEvent, error) {
	var r ReservationEvent
	if err := DB.Model(&ReservationEvent{}).Where("id = ?", id).Scan(&r).Error; err != nil {
		return nil, fmt.Errorf("reservation not found: %w", err)
	}
	return &r, nil
}

func GetEventReservationsByState(state string) ([]ReservationEvent, error) {
	var reservations []ReservationEvent
	if err := DB.Model(&ReservationEvent{}).Where("reservation_state = ?", state).Scan(reservations).Error; err != nil {
		return nil, fmt.Errorf("reservation not found: %w", err)
	}
	return reservations, nil
}
