package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Type         string
	Reservations []Reservation
}

var PaymentList = []Payment{
	{Type: "PayPal"},
	{Type: "Credit Card"},
	{Type: "Bitcoin"},
}

func (p *Payment) save() error {
	if err := DB.Create(&p).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}
	return nil
}

func AddPayments() error {
	for _, payment := range PaymentList {
		if err := payment.save(); err != nil {
			return fmt.Errorf("payment.Save: %w", err)
		}
	}
	return nil
}

func GetAllPayments() ([]Payment, error) {
	var p []Payment

	if err := DB.Model([]Payment{}).Scan(&p).Error; err != nil {
		return p, fmt.Errorf("DB.Model([]Payment{}).Scan: %w", err)
	}
	return p, nil
}
