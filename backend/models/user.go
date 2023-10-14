package models

import (
	"errors"
	"fmt"

	"Participate/utils/token"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Email    string `gorm:"size:100;not null;unique" json:"email"`
	Password string `gorm:"size:100;not null;" json:"password"`
}

func GetCustomerByID(uid uint) (Customer, error) {

	var c Customer

	if err := DB.First(&c, uid).Error; err != nil {
		return c, errors.New("user not found")
	}

	c.PrepareGive()

	return c, nil

}

func (c *Customer) PrepareGive() {
	c.Password = ""
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(email string, password string) (string, error) {

	var err error

	c := Customer{}

	err = DB.Model(Customer{}).Where("email = ?", email).Take(&c).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, c.Password)

	if err != nil && errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return "", err
	}

	t, err := token.GenerateToken(c.ID)

	if err != nil {
		return "", err
	}

	return t, nil

}

func (c *Customer) SaveCustomer() (*Customer, error) {
	fmt.Println("Started saving")

	var err error
	err = DB.Create(&c).Error
	if err != nil {
		return &Customer{}, err
	}
	fmt.Println("Finished saving")

	return c, nil
}

func (c *Customer) BeforeSave(*gorm.DB) error {
	fmt.Println("Starting before save")
	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	c.Password = string(hashedPassword)

	fmt.Println("Finished before save")

	return nil

}
