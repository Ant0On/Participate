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
	RatingID uint      `json:"rating_id"`
	RoomID   uint      `gorm:"not null" json:"room_id" binding:"required"`
}

func (r *ReservationRoom) validate() error {
	var offer Room
	if err := DB.First(&offer, r.RoomID).Error; err != nil {
		return fmt.Errorf("failed to retrieve room: %w", err)
	}

	if r.NumberOfPeople > offer.Capacity {
		return fmt.Errorf("too many people added to reservation")
	}
	return nil
}

func (r *ReservationRoom) Save() error {
	if err := r.validate(); err != nil {
		return fmt.Errorf("reservation validation failed: %v", err)
	}

	if err := DB.Create(r).Error; err != nil {
		return fmt.Errorf("failed to save reservation: %w", err)
	}

	return nil
}

func (r *ReservationRoom) Update() error {
	if err := r.validate(); err != nil {
		return fmt.Errorf("reservation validation failed: %v", err)
	}
	if err := DB.Save(r).Error; err != nil {
		return fmt.Errorf("failed to update reservation: %w", err)
	}
	return nil
}

func (r *ReservationRoom) Delete() error {
	if err := DB.Delete(&r).Error; err != nil {
		return fmt.Errorf("failed to delete reservation: %w", err)
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
		return nil, fmt.Errorf("failed to retrieve reservations: %w", err)
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
