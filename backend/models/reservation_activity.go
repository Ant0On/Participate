package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type ReservationActivity struct {
	gorm.Model
	Reservation
	Date       time.Time `gorm:"not null" json:"date" binding:"required"`
	RatingID   uint      `json:"rating_id"`
	ActivityID uint      `json:"activity_id"`
}

func (r *ReservationActivity) Save() error {
	if err := r.Validate(); err != nil {
		return fmt.Errorf("r.Validate: %v", err)
	}

	if err := DB.Create(r).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}

func (r *ReservationActivity) Update() error {
	if err := r.Validate(); err != nil {
		return fmt.Errorf("r.Validate: %v", err)
	}
	if err := DB.Save(r).Error; err != nil {
		return err
	}
	return nil
}

func (r *ReservationActivity) Delete() error {
	if err := DB.Delete(&r).Error; err != nil {
		return fmt.Errorf("DB.Delete: %w", err)
	}
	return nil
}

func GetActivityReservationById(id string) (ReservationOperations, error) {
	var r ReservationActivity
	if err := DB.First(&r, id).Error; err != nil {
		return nil, fmt.Errorf("reservation not found: %w", err)
	}
	return ReservationOperations(&r), nil
}

func GetActivityReservationsByState(state string) ([]ReservationOperations, error) {
	var reservations []ReservationActivity
	if err := DB.Model(&ReservationActivity{}).Where("reservation_state = ?", state).Scan(reservations).Error; err != nil {
		return nil, fmt.Errorf("reservation not found: %w", err)
	}
	var ops []ReservationOperations
	for _, r := range reservations {
		ops = append(ops, ReservationOperations(&r))
	}

	return ops, nil
}

func (r *ReservationActivity) ChangeState(state string) error {
	r.ReservationState = ReservationState(state)
	return r.Update()
}

func (r *ReservationActivity) ChangeCapacity() error {
	var a Activity
	if err := DB.First(&a, r.ActivityID).Error; err != nil {
		return fmt.Errorf("DB.First: %w", err)
	}
	a.Capacity -= r.NumberOfPeople
	return nil
}
