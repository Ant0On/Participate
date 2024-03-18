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

func GetActivityReservationById(id string) (*ReservationActivity, error) {
	var r *ReservationActivity
	if err := DB.Model(&ReservationActivity{}).Where("id = ?", id).Scan(r).Error; err != nil {
		return nil, fmt.Errorf("reservation not found: %w", err)
	}
	return r, nil
}

func GetActivityReservationsByState(state string) ([]ReservationActivity, error) {
	var reservations []ReservationActivity
	if err := DB.Model(&ReservationActivity{}).Where("reservation_state = ?", state).Scan(reservations).Error; err != nil {
		return nil, fmt.Errorf("reservation not found: %w", err)
	}
	return reservations, nil
}
