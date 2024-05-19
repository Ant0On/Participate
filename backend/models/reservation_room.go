package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type ReservationRoom struct {
	gorm.Model
	Reservation
	DateFrom time.Time `gorm:"not null" json:"date_from" binding:"required"`
	DateTo   time.Time `gorm:"not null" json:"date_to" binding:"required,gtfield=DateFrom"`
	RoomID   uint      `json:"room_id"`
}

func (r *ReservationRoom) Save() error {
	if err := r.Validate(); err != nil {
		return fmt.Errorf("r.Validate: %v", err)
	}

	if err := DB.Create(r).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}

func (r *ReservationRoom) Update() error {
	if err := r.Validate(); err != nil {
		return fmt.Errorf("r.Validate: %v", err)
	}
	if err := DB.Save(r).Error; err != nil {
		return err
	}
	return nil
}

func (r *ReservationRoom) Delete() error {
	if err := DB.Delete(&r).Error; err != nil {
		return fmt.Errorf("DB.Delete: %w", err)
	}
	return nil
}

func GetRoomReservationById(id string) (ReservationOperations, error) {
	var r ReservationRoom
	if err := DB.First(&r, id).Error; err != nil {
		return nil, fmt.Errorf("reservation not found: %w", err)
	}
	return ReservationOperations(&r), nil
}

func GetRoomReservationsByState(state string) ([]ReservationOperations, error) {
	var reservations []ReservationRoom
	if err := DB.Model(&ReservationRoom{}).Where("reservation_state = ?", state).Scan(reservations).Error; err != nil {
		return nil, fmt.Errorf("reservation not found: %w", err)
	}
	var ops []ReservationOperations
	for _, r := range reservations {
		ops = append(ops, ReservationOperations(&r))
	}

	return ops, nil
}

func (r *ReservationRoom) ChangeState(state string) error {
	r.ReservationState = ReservationState(state)
	return r.Update()
}

func (r *ReservationRoom) ChangeCapacity() error {
	return nil
}
