package models

import (
	"fmt"

	"backend/pkg/passHelper"
	"backend/utils/token"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	FirstName    string `gorm:"size:30;not null" form:"first_name" binding:"required"`
	LastName     string `gorm:"size:100" form:"last_name"`
	Email        string `gorm:"size:100;not null;unique" form:"email" binding:"required"`
	ImagePath    string `gorm:"default:images/customers/default_image.png" form:"image_path"`
	Password     string `gorm:"size:100;not null;" form:"password" binding:"required"`
	Role         string `gorm:"default:customer"`
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

func GetCustomer(id string) (*Customer, error) {
	var c Customer
	if err := DB.Model(&Customer{}).Where("id = ?", id).Scan(&c).Error; err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	return &c, nil
}

func (c *Customer) Save() error {
	if err := DB.Create(&c).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}

func (c *Customer) Update() error {
	if err := DB.Save(&c).Error; err != nil {
		return err
	}
	return nil
}

func LoginCheck(email, password string) (string, any, error) {
	var role, uPassword, t string
	var user any
	var err error
	c := Customer{}

	if role, uPassword, err = c.checkIfEmailExist(email); err != nil {
		return "", nil, fmt.Errorf("user with given email doesn't exist: %w", err)
	}

	if err = passHelper.VerifyPassword(password, uPassword); err != nil {
		return "", nil, fmt.Errorf("VerifyPassword: %s", role)

	}

	if t, err = token.GenerateToken(c.Email, role); err != nil {
		return "", nil, fmt.Errorf("token.GenerateToken: %w", err)
	}

	if user, err = GetUser(email); err != nil {
		return "", nil, fmt.Errorf("GetUser: %w", err)
	}

	return t, user, nil
}
func (c *Customer) checkIfEmailExist(email string) (string, string, error) {
	h := NewHost()

	if err := DB.Where("email = ?", email).First(&c).Error; err != nil {
		if err = DB.Where("email = ?", email).First(&h).Error; err != nil {
			return "", "", fmt.Errorf("user not found: %w", err)
		}
		return h.Role, h.Password, nil
	}
	return c.Role, c.Password, nil
}

func GetUser(email string) (any, error) {
	var c *Customer
	h := NewHost()

	if err := DB.Where("email = ?", email).First(&c).Error; err != nil {
		if err = DB.Where("email = ?", email).First(&h).Error; err != nil {
			return "", fmt.Errorf("user not found: %w", err)
		}
		// hide sensitive data
		h.Password = ""
		return h, nil
	}
	// hide sensitive data
	c.Password = ""
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
