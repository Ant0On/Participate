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

func GetHostById(uid uint) (any, error) {
	var h Host

	if err := DB.First(h, uid).Error; err != nil {
		return &h, errors.New("user not found")
	}

	h.prepareGive()

	return &h, nil
}

func (h *Host) prepareGive() {
	h.Password = ""
}

func (h *Host) CheckHostEmailExists(email string) error {
	if err := DB.Model(Host{}).Where("email = ?", email).Take(h).Error; err != nil {
		return fmt.Errorf("DB.Model.Where.Take: %w", err)
	}
	return nil
}

func (h *Host) LoginCheck(email, password string) (string, error) {

	if err := h.CheckHostEmailExists(email); err != nil {
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

func (h *Host) Save() error {
	if err := DB.Create(&h).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}

func (h *Host) BeforeSave(*gorm.DB) error {
	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(h.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("bcrypt.GenerateFromPassword: %w", err)
	}
	h.Password = string(hashedPassword)

	return nil
}

func (h *Host) AccountType(ctx *gin.Context) (string, error) {
	t, err := h.LoginCheck(h.Email, h.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"models.LoginCheck: username or password is incorrect": err.Error()})
		return "", err
	}

	return t, nil
}
