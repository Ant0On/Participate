package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Host struct {
	gorm.Model
	*Customer
	Description string  `gorm:"not null;" form:"description" binding:"required,min=15,max=255"`
	PhoneNumber string  `gorm:"size:12;not null;unique" form:"phone_number" binding:"required,min=9,max=15"`
	BankAccount string  `gorm:"size:31;not null;unique" form:"bank_account" binding:"required,min=9,max=12"`
	Offers      []Offer `gorm:"foreignKey:HostID"`
}

func NewHost() Host {
	return Host{
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

func GetHost(id string) (*Host, error) {
	var h Host
	if err := DB.Model(&Host{}).Where("id = ?", id).Scan(&h).Error; err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	return &h, nil
}
