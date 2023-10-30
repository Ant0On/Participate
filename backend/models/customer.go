package models

import (
	"errors"
	"fmt"

	"backend/pkg/passHelper"
	"backend/utils/token"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	FirstName    string `gorm:"size:30;not null" json:"first_name"`
	LastName     string `gorm:"size:100;not null" json:"last_name"`
	Email        string `gorm:"size:100;not null;unique" json:"email"`
	Password     string `gorm:"size:100;not null;" json:"password"`
	Reservations []Reservation
}

func GetUserByEmail(email string) (any, error) {
	var c Customer
	var h Host

	if err := DB.Model(&Customer{}).Where("email = ?", email).Scan(&c).Error; err != nil {
		if err := DB.Model(&Host{}).Where("email = ?", email).Scan(&h).Error; err != nil {
			return &h, fmt.Errorf("user not found: %w", err)
		}
		//Reset password to hide it from JSON
		h.Password = ""
		return &h, nil
	}

	//Reset password to hide it from JSON
	c.Password = ""

	return &c, nil
}

func (c *Customer) Save() error {
	if err := DB.Create(&c).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}

func (c *Customer) LoginCheck(email, password, table string) (string, error) {
	if err := c.checkIfEmailExist(email, table); err != nil {
		return "", fmt.Errorf("user with given email doesn't exist: %w", err)
	}

	if err := passHelper.VerifyPassword(password, c.Password); err != nil && errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return "", fmt.Errorf("VerifyPassword: %w", err)
	}

	t, err := token.GenerateToken(c.Email)

	if err != nil {
		return "", fmt.Errorf("token.GenerateToken: %w", err)
	}

	return t, nil
}
func (c *Customer) checkIfEmailExist(email, table string) error {
	if err := DB.Table(table).Where("email = ?", email).First(&c).Error; err != nil {
		return fmt.Errorf("DB.Table.Where.First: %w", err)
	}
	return nil
}

func (c *Customer) BeforeSave(*gorm.DB) error {
	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("bcrypt.GenerateFromPassword: %w", err)
	}
	c.Password = string(hashedPassword)

	return nil
}
