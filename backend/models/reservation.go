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
	DateFrom         time.Time        `gorm:"not null" json:"date_from" binding:"required"`
	DateTo           time.Time        `gorm:"not null" json:"date_to" binding:"required,gtfield=DateFrom"`
	ReservationState ReservationState `gorm:"type:varchar(255);check:reservation_state IN ('pending', 'accepted', 'ongoing', 'finished', 'rejected'); column:reservation_state; not null" json:"reservation_state" binding:"required,oneof=pending accepted ongoing finished rejected"`
	NumberOfPeople   int              `gorm:"not null" json:"number_of_people" binding:"required,gt=0"`
	UserID           uint             `gorm:"not null" json:"user_id" binding:"required"`
	OfferID          uint             `gorm:"not null" json:"offer_id" binding:"required"`
	GradeID          uint             `json:"grade_id"`
	PaymentID        uint             `gorm:"not null" json:"payment_id" binding:"required"`
	AnimalID         uint
}

func (r *Reservation) ValidateDates() error {
	if r.DateFrom.Before(time.Now()) || r.DateTo.Before(time.Now()) {
		return fmt.Errorf("reservation dates cannot be in the past")
	}

	return nil
}

func (r *Reservation) Validate() error {
	var offer Offer
	if err := DB.First(&offer, r.OfferID).Error; err != nil {
		return fmt.Errorf("DB.First: %w", err)
	}

	if r.NumberOfPeople > offer.MaxPeople {
		return fmt.Errorf("too many people added to reservation")
	}
	return nil
}

func (r *Reservation) Save() error {
	if err := r.Validate(); err != nil {
		return fmt.Errorf("r.Validate: %v", err)
	}

	if err := DB.Create(&r).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}

func (r *Reservation) Update() error {
	if err := r.Validate(); err != nil {
		return fmt.Errorf("r.Validate: %v", err)
	}
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

func GetReservationsByState(state string) ([]Reservation, error) {
	var reservations []Reservation
	if err := DB.Model(&Reservation{}).Where("reservation_state = ?", state).Scan(reservations).Error; err != nil {
		return nil, fmt.Errorf("reservation not found: %w", err)
	}
	return reservations, nil
}

func CheckReservations() error {
	var reservations []Reservation
	if err := DB.Model(&Reservation{}).Scan(reservations).Error; err != nil {
		return fmt.Errorf("reservations not found: %w", err)
	}

	for _, reservation := range reservations {
		if time.Now().After(reservation.DateTo) {
			reservation.ReservationState = "finished"
			if err := reservation.Update(); err != nil {
				return fmt.Errorf("reservation update: %w", err)
			}
		}
	}
	return nil
}
