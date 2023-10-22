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

type Host struct {
	gorm.Model
	FirstName   string `gorm:"size:30;not null" json:"first_name"`
	LastName    string `gorm:"size:100;not null" json:"last_name"`
	Email       string `gorm:"size:100;not null;unique" json:"email"`
	PhoneNumber string `gorm:"size:12;not null;unique" json:"phone_number"`
	BankAccount string `gorm:"size:31;not null;unique" json:"bank_account"`
	Password    string `gorm:"size:100;not null" json:"password"`
	Offers      []Offer
}

func (h *Host) LoginCheck(email, password string) (string, error) {

	if err := services.CheckHostEmailExists(email, h); err != nil {
		return "", fmt.Errorf("customer with given email doesn't exist")
	}

	if err := verifyPassword(password, h.Password); err != nil && errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return "", fmt.Errorf("verifyPassword: %w", err)
	}

	t, err := token.GenerateToken(h.ID)

	if err != nil {
		return "", fmt.Errorf("token.GenerateToken: %w", err)
	}

	return t, nil
}

func (h *Host) Safe() error {
	if _, err := services.SaveHost(h); err != nil {
		return fmt.Errorf("error with saving customer")
	}
	return nil
}

func (h *Host) AccountType(ctx *gin.Context) (string, error) {
	t, err := h.LoginCheck(h.Email, h.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("models.LoginCheck: %w \nusername or password is incorrect", err)})
		return "", err
	}

	return t, nil
}
