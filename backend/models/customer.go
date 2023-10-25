package models

import (
	"errors"
	"fmt"
	"net/http"

	"backend/utils/token"

	"github.com/gin-gonic/gin"
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

func GetCustomerById(uid uint) (any, error) {
	var c Customer

	if err := DB.First(c, uid).Error; err != nil {
		return &c, errors.New("user not found")
	}

	c.prepareGive()

	return &c, nil
}

func (c *Customer) prepareGive() {
	c.Password = ""
}

func (c *Customer) CheckCustomerEmailExists(email string) error {
	if err := DB.Model(Customer{}).Where("email = ?", email).Take(&c).Error; err != nil {
		return fmt.Errorf("DB.Model.Where.Take: %w", err)
	}
	return nil
}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (c *Customer) loginCheck(email, password string) (string, error) {

	if err := c.CheckCustomerEmailExists(email); err != nil {
		return "", fmt.Errorf("customer with given email doesn't exist")
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

func (c *Customer) Save() error {
	if err := DB.Create(&c).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
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

func (c *Customer) AccountType(ctx *gin.Context) (string, error) {
	t, err := c.loginCheck(c.Email, c.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"customer.loginCheck: username or password is incorrect": err.Error()})
		return "", err
	}

	return t, nil
}
