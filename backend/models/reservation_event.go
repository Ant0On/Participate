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
	EventID uint      `gorm:"not null" json:"event_id" binding:"required"`
}

func (r *ReservationEvent) validate() error {
	var offer Event
	if err := DB.First(&offer, r.EventID).Error; err != nil {
		return fmt.Errorf("failed to retrieve event: %w", err)
	}

	if r.NumberOfPeople > offer.Capacity {
		return fmt.Errorf("too many people added to reservation")
	}
	return nil
}

func (r *ReservationEvent) Save() error {
	if err := r.validate(); err != nil {
		return fmt.Errorf("reservation validation failed: %v", err)
	}

	if err := DB.Create(r).Error; err != nil {
		return fmt.Errorf("failed to save reservation: %w", err)
	}

	return nil
}

func (r *ReservationEvent) Update() error {
	if err := r.validate(); err != nil {
		return fmt.Errorf("reservation validation failed: %v", err)
	}
	if err := DB.Save(r).Error; err != nil {
		return fmt.Errorf("failed to update reservation: %w", err)
	}
	return nil
}

func (r *ReservationEvent) Delete() error {
	if err := DB.Delete(&r).Error; err != nil {
		return fmt.Errorf("failed to delete reservation: %w", err)
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
		return nil, fmt.Errorf("failed to retrieve reservations: %w", err)
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

func (r *ReservationEvent) ChangeCapacity() error {
	var e Event
	if err := DB.First(&e, r.EventID).Error; err != nil {
		return fmt.Errorf("failed to retrieve event: %w", err)
	}
	e.Capacity -= r.NumberOfPeople
	if err := DB.Save(&e).Error; err != nil {
		return fmt.Errorf("failed to update event capacity: %w", err)
	}
	return nil
}
