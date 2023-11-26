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
	FirstName    string `gorm:"size:30;not null" json:"first_name" binding:"required"`
	LastName     string `gorm:"size:100;not null" json:"last_name" binding:"required"`
	Email        string `gorm:"size:100;not null;unique" json:"email" binding:"required"`
	Password     string `gorm:"size:100;not null;" json:"password" binding:"required"`
	Role         string `gorm:"size:100;not null;default:Customer"`
	Reservations []Reservation
}

func GetUserByEmail(email string) (any, error) {
	var c Customer
	var h Host

	if err := DB.Model(&Customer{}).Where("email = ?", email).Scan(&c).Error; err != nil {
		if err = DB.Model(&Host{}).Where("email = ?", email).Scan(&h).Error; err != nil {
			return nil, fmt.Errorf("user not found: %w", err)
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

func (c *Customer) LoginCheck(email, password string) (string, error) {
	var role string
	var err error
	if role, err = c.checkIfEmailExist(email); err != nil {
		return "", fmt.Errorf("user with given email doesn't exist: %w", err)
	}

	if err = passHelper.VerifyPassword(password, c.Password); err != nil && errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return "", fmt.Errorf("VerifyPassword: %w", err)
	}

	t, err := token.GenerateToken(c.Email, role)

	if err != nil {
		return "", fmt.Errorf("token.GenerateToken: %w", err)
	}

	return t, nil
}
func (c *Customer) checkIfEmailExist(email string) (string, error) {
	var h Host

	if err := DB.Model(&Customer{}).Where("email = ?", email).Scan(&c).Error; err != nil {
		if err = DB.Model(&Host{}).Where("email = ?", email).Scan(&h).Error; err != nil {
			return "", fmt.Errorf("user not found: %w", err)
		}
		return h.Role, nil
	}
	return c.Role, nil
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
