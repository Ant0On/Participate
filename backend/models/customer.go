package models

import (
	"errors"
	"fmt"
	"net/http"

	"backend/services"
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

func GetCustomerByID(uid uint) (Customer, error) {
	c, _ := services.GetCustomerByID(uid)

	c.prepareGive()

	return *c, nil
}

func (c *Customer) prepareGive() {
	c.Password = ""
}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (c *Customer) LoginCheck(email, password string) (string, error) {

	if err := services.CheckCustomerEmailExists(email, c); err != nil {
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
	if _, err := services.SaveCustomer(c); err != nil {
		return fmt.Errorf("error with saving customer")
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
	t, err := c.LoginCheck(c.Email, c.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("models.LoginCheck: %w \nusername or password is incorrect", err)})
		return "", err
	}

	return t, nil
}
