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
func (r *ReservationEvent) Delete() error {
	if err := DB.Delete(&r).Error; err != nil {
		return fmt.Errorf("DB.Delete: %w", err)
	}
	return nil
}

func GetEventReservationById(id string) (ReservationOperations, error) {
	var r ReservationEvent
	if err := DB.First(&r, id).Error; err != nil {
		return nil, fmt.Errorf("reservation not found: %w", err)
	}
	return ReservationOperations(&r), nil
}

func GetEventReservationsByState(state string) ([]ReservationOperations, error) {
	var reservations []ReservationEvent
	if err := DB.Model(&ReservationEvent{}).Where("reservation_state = ?", state).Scan(reservations).Error; err != nil {
		return nil, fmt.Errorf("reservation not found: %w", err)
	}
	var ops []ReservationOperations
	for _, r := range reservations {
		ops = append(ops, ReservationOperations(&r))
	}

	return ops, nil
}

func (r *ReservationEvent) ChangeState(state string) error {
	r.ReservationState = ReservationState(state)
	return r.Update()
}
