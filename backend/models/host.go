package models

import "gorm.io/gorm"

type Host struct {
	gorm.Model
	FirstName   string `gorm:"size:30;not null" json:"first_name"`
	LastName    string `gorm:"size:100;not null" json:"last_name"`
	Email       string `gorm:"size:100;not null;unique" json:"email"`
	PhoneNumber string `gorm:"size:12;not null;unique" json:"phone_number"`
	BankAccount string `gorm:"size:31;not null;unique" json:"bank_account"`
	Password    string `gorm:"size:100;not null" json:"password"`
}
