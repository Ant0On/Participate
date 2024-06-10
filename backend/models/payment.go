package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Type                      string
	ReservationsAccommodation []ReservationAccommodation
	ReservationsActivity      []ReservationActivity
	ReservationsEvent         []ReservationEvent
	ReservationsRoom          []ReservationRoom
}

var PaymentList = []Payment{
	{Type: "PayPal"},
	{Type: "Credit Card"},
	{Type: "Bitcoin"},
}

func (p *Payment) save() error {
	if err := DB.Create(&p).Error; err != nil {
		return fmt.Errorf("failed to create payment: %w", err)
	}
	return nil
}

func AddPayments() error {
	for _, payment := range PaymentList {
		if err := payment.save(); err != nil {
			return fmt.Errorf("failed to add payment: %w", err)
		}
	}
	return nil
}

func GetAllPayments() ([]Payment, error) {
	var payments []Payment

	if err := DB.Model(&Payment{}).Scan(&payments).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve payments: %w", err)
	}
	return payments, nil
}
