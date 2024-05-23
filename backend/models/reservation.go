package models

import (
	"fmt"
	"time"
)

type ReservationState string

type ReservationOperations interface {
	Save() error
	Update() error
	Delete() error
	ChangeState(state string) error
	ChangeCapacity() error
}

const (
	Pending  ReservationState = "pending"
	Accepted ReservationState = "accepted"
	Ongoing  ReservationState = "ongoing"
	Finished ReservationState = "finished"
	Rejected ReservationState = "rejected"
)

type Reservation struct {
	ReservationState ReservationState `gorm:"type:varchar(255);check:reservation_state IN ('pending', 'accepted', 'ongoing', 'finished', 'rejected'); column:reservation_state; not null; default:'pending'" json:"reservation_state,omitempty"`
	NumberOfPeople   int              `gorm:"not null" json:"number_of_people" binding:"required,gt=0"`
	UserID           uint             `gorm:"not null" json:"user_id" binding:"required"`
	PaymentID        uint             `gorm:"not null" json:"payment_id" binding:"required"`
}

// TODO move validation to specific reservations
func (r *Reservation) Validate() error {
	//var offer Offer
	//if err := DB.First(&offer, r.OfferID).Error; err != nil {
	//	return fmt.Errorf("DB.First: %w", err)
	//}
	//
	//if r.NumberOfPeople > offer.Capacity {
	//	return fmt.Errorf("too many people added to reservation")
	//}
	return nil
}

func CheckReservations() error {
	var activityReservation []ReservationActivity
	var accommodationReservation []ReservationAccommodation
	var eventReservation []ReservationEvent
	if err := DB.Model(&ReservationActivity{}).Scan(activityReservation).Error; err != nil {
		return fmt.Errorf("reservations not found: %w", err)
	}
	if err := DB.Model(&ReservationAccommodation{}).Scan(accommodationReservation).Error; err != nil {
		return fmt.Errorf("reservations not found: %w", err)
	}
	if err := DB.Model(&ReservationEvent{}).Scan(eventReservation).Error; err != nil {
		return fmt.Errorf("reservations not found: %w", err)
	}

	for _, reservation := range activityReservation {
		if time.Now().After(reservation.Date) {
			reservation.ReservationState = "finished"
			if err := reservation.Update(); err != nil {
				return fmt.Errorf("reservation update: %w", err)
			}
		}
	}
	for _, reservation := range accommodationReservation {
		if time.Now().After(reservation.DateTo) {
			reservation.ReservationState = "finished"
			if err := reservation.Update(); err != nil {
				return fmt.Errorf("reservation update: %w", err)
			}
		}
	}
	for _, reservation := range eventReservation {
		if time.Now().After(reservation.Date) {
			reservation.ReservationState = "finished"
			if err := reservation.Update(); err != nil {
				return fmt.Errorf("reservation update: %w", err)
			}
		}
	}
	return nil
}
