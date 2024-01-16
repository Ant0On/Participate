package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Host struct {
	gorm.Model
	*Customer
	Description string `gorm:"not null;" json:"description" binding:"required"`
	PhoneNumber string `gorm:"size:12;not null;unique" json:"phone_number" binding:"required"`
	BankAccount string `gorm:"size:31;not null;unique" json:"bank_account" binding:"required"`
	Offers      []Offer
}

func NewHost() *Host {
	return &Host{
		Customer: &Customer{
			Role: "host",
		},
	}
}

func (h *Host) Save() error {
	if err := DB.Create(&h).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}

func GetHost(host *Host, id string) (*Host, error) {
	if err := DB.First(host, id).Error; err != nil {
		return nil, fmt.Errorf("DB.First: %w", err)
	}

	return host, nil
}
