package models

import (
	"errors"
	"fmt"

	"Participate/backend/utils/token"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	FirstName string `gorm:"size:30;not null" json:"first_name"`
	LastName  string `gorm:"size:100;not null" json:"last_name"`
	Email     string `gorm:"size:100;not null;unique" json:"email"`
	Password  string `gorm:"size:100;not null;" json:"password"`
}

func GetCustomerByID(uid uint) (Customer, error) {
	var c Customer

	if err := DB.First(&c, uid).Error; err != nil {
		return c, errors.New("user not found")
	}

	c.prepareGive()

	return c, nil

}

func (c *Customer) prepareGive() {
	c.Password = ""
}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(email string, password string) (string, error) {
	c := Customer{}

	if err := DB.Model(Customer{}).Where("email = ?", email).Take(&c).Error; err != nil {
		return "", fmt.Errorf("DB.Model.Where.Take: %w", err)
	}

	if err := verifyPassword(password, c.Password); err != nil && errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return "", fmt.Errorf("verifyPassword: %w", err)
	}

	t, err := token.GenerateToken(c.ID)

	if err != nil {
		return "", fmt.Errorf("token.GenerateToken: %w", err)
	}

	return t, nil

}

func (c *Customer) SaveCustomer() (*Customer, error) {
	if err := DB.Create(&c).Error; err != nil {
		return &Customer{}, fmt.Errorf("DB.Create: %w", err)
	}

	return c, nil
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
