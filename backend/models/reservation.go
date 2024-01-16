package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type ReservationState string

const (
	Pending  ReservationState = "pending"
	Accepted ReservationState = "accepted"
	Ongoing  ReservationState = "ongoing"
	Finished ReservationState = "finished"
	Rejected ReservationState = "rejected"
)

type Reservation struct {
	gorm.Model
	DateFrom         time.Time        `gorm:"not null" json:"date_from"`
	DateTo           time.Time        `gorm:"not null" json:"date_to"`
	ReservationState ReservationState `gorm:"type:varchar(255);check:reservation_state IN ('pending', 'accepted', 'ongoing', 'finished', 'rejected'); column:reservation_state; not null" json:"reservation_state"`
	CustomerID       uint             `gorm:"not null"`
	OfferID          uint             `gorm:"not null"`
	GradeID          uint
	DiscountID       uint
	PaymentID        uint `gorm:"not null"`
	AnimalID         uint
}

func (r *Reservation) Save() error {
	if err := DB.Create(&r).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}

func (r *Reservation) Update() error {
	if err := DB.Save(&r).Error; err != nil {
		return err
	}
	return nil
}

func GetReservationById(id string) (*Reservation, error) {
	var r Reservation
	if err := DB.Model(&Reservation{}).Where("id = ?", id).Scan(&r).Error; err != nil {
		return nil, fmt.Errorf("reservation not found: %w", err)
	}
	return &r, nil
}
