package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Host struct {
	gorm.Model
	*Customer
	PhoneNumber string `gorm:"size:12;not null;unique" json:"phone_number" binding:"required"`
	BankAccount string `gorm:"size:31;not null;unique" json:"bank_account" binding:"required"`
	Offers      []Offer
}

func NewHost() *Host {
	return &Host{
		Customer: &Customer{
			Role: "Host",
		},
	}
}

func (h *Host) Save() error {
	if err := DB.Create(&h).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}
