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
	Finished ReservationState = "finished"
)

type Reservation struct {
	ReservationState ReservationState `gorm:"type:varchar(255);check:reservation_state IN ('pending', 'accepted', 'ongoing', 'finished', 'rejected'); column:reservation_state; not null; default:'pending'" json:"-"`
	NumberOfPeople   int              `gorm:"not null" json:"number_of_people" binding:"required,gt=0"`
	UserID           uint             `gorm:"not null" json:"user_id" binding:"required"`
	PaymentID        uint             `gorm:"not null" json:"payment_id" binding:"required"`
}

func CheckReservations() error {
	var activityReservations []ReservationActivity
	var accommodationReservations []ReservationAccommodation
	var eventReservations []ReservationEvent

	if err := DB.Model(&ReservationActivity{}).Scan(&activityReservations).Error; err != nil {
		return fmt.Errorf("activity reservations not found: %w", err)
	}
	if err := DB.Model(&ReservationAccommodation{}).Scan(&accommodationReservations).Error; err != nil {
		return fmt.Errorf("accommodation reservations not found: %w", err)
	}
	if err := DB.Model(&ReservationEvent{}).Scan(&eventReservations).Error; err != nil {
		return fmt.Errorf("event reservations not found: %w", err)
	}

	now := time.Now()

	for _, reservation := range activityReservations {
		if reservation.ReservationState == "accepted" && now.Format("2006-01-02") == reservation.Date.Format("2006-01-02") {
			reservation.ReservationState = "ongoing"
			if err := reservation.Update(); err != nil {
				return fmt.Errorf("reservation update: %w", err)
			}
		} else if reservation.ReservationState == "ongoing" && now.After(reservation.Date) {
			reservation.ReservationState = "finished"
			if err := reservation.Update(); err != nil {
				return fmt.Errorf("reservation update: %w", err)
			}
		}
	}
	for _, reservation := range accommodationReservations {
		if reservation.ReservationState == "accepted" && now.Format("2006-01-02") == reservation.DateFrom.Format("2006-01-02") {
			reservation.ReservationState = "ongoing"
			if err := reservation.Update(); err != nil {
				return fmt.Errorf("reservation update: %w", err)
			}
		} else if reservation.ReservationState == "ongoing" && now.After(reservation.DateTo) {
			reservation.ReservationState = "finished"
			if err := reservation.Update(); err != nil {
				return fmt.Errorf("reservation update: %w", err)
			}
		}
	}
	for _, reservation := range eventReservations {
		if reservation.ReservationState == "accepted" && now.Format("2006-01-02") == reservation.Date.Format("2006-01-02") {
			reservation.ReservationState = "ongoing"
			if err := reservation.Update(); err != nil {
				return fmt.Errorf("reservation update: %w", err)
			}
		} else if reservation.ReservationState == "ongoing" && now.After(reservation.Date) {
			reservation.ReservationState = "finished"
			if err := reservation.Update(); err != nil {
				return fmt.Errorf("reservation update: %w", err)
			}
		}
	}
	return nil
}
