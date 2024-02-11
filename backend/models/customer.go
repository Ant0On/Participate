package models

import (
	"fmt"
	"path/filepath"
	"unicode/utf8"

	"backend/pkg/passHelper"
	"backend/utils/token"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	FirstName    string `gorm:"size:30;not null" form:"first_name" binding:"required,min=2,max=30"`
	LastName     string `gorm:"size:100" form:"last_name"`
	Email        string `gorm:"size:100;not null;unique" form:"email" binding:"required,email"`
	ImagePath    string `gorm:"default:images/customers/default_image.png" form:"image_path" binding:"omitempty,url"`
	Password     string `gorm:"size:100;not null" form:"password" binding:"required,min=8"`
	Role         string `gorm:"size:20;not null;default:'customer'" form:"-"`
	Reservations []Reservation
	Messages     []Message
}

func validateLastName(lastName string) error {
	if utf8.RuneCountInString(lastName) == 0 || (utf8.RuneCountInString(lastName) >= 2 && utf8.RuneCountInString(lastName) <= 100) {
		return nil
	}

	return fmt.Errorf("LastName must be empty or have a length between 2 and 100 characters")
}

func GetUserByEmail(email string) (any, error) {
	var c Customer
	var h Host

	if err := DB.Model(&Customer{}).Where("email = ?", email).Scan(&c).Error; err != nil {
		if err = DB.Model(&Host{}).Where("email = ?", email).Scan(&h).Error; err != nil {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		h.Password = ""
		return &h, nil
	}
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
	if err := validateLastName(c.LastName); err != nil {
		return fmt.Errorf("validateLastName: %v", err)
	}
	if err := DB.Create(&c).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}

func (c *Customer) Update() error {
	if err := validateLastName(c.LastName); err != nil {
		return fmt.Errorf("validateLastName: %v", err)
	}
	if err := DB.Save(&c).Error; err != nil {
		return err
	}
	return nil
}

func (c *Customer) Delete() error {
	if err := DB.Delete(&c).Error; err != nil {
		return err
	}
	return nil
}

func LoginCheck(email, password string) (string, any, error) {
	var role, uPassword, t string
	var userID uint
	var user any
	var err error
	c := Customer{}

	if userID, role, uPassword, err = c.checkIfEmailExist(email); err != nil {
		return "", nil, fmt.Errorf("user with given email doesn't exist: %w", err)
	}

	if err = passHelper.VerifyPassword(password, uPassword); err != nil {
		return "", nil, fmt.Errorf("VerifyPassword: wrong password")
	}

	if t, err = token.GenerateToken(userID, email, role); err != nil {
		return "", nil, fmt.Errorf("token.GenerateToken: %w", err)
	}

	if user, err = GetUser(email); err != nil {
		return "", nil, fmt.Errorf("GetUser: %w", err)
	}

	return t, user, nil
}

func (c *Customer) checkIfEmailExist(email string) (uint, string, string, error) {
	h := NewHost()

	if err := DB.Where("email = ?", email).First(&c).Error; err != nil {
		if err = DB.Where("email = ?", email).First(&h).Error; err != nil {
			return 0, "", "", fmt.Errorf("user not found: %w", err)
		}
		return h.ID, h.Role, h.Password, nil
	}
	return c.ID, c.Role, c.Password, nil
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

func (c *Customer) BeforeCreate(*gorm.DB) error {
	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("bcrypt.GenerateFromPassword: %w", err)
	}
	c.Password = string(hashedPassword)

	return nil
}

func (c *Customer) HandleUserImageUploads(context *gin.Context, userID uint, role string) (string, bool, error) {
	form, err := context.MultipartForm()
	if err != nil {
		return "", false, fmt.Errorf("multipart form error: %v", err)
	}

	files := form.File["image"]

	if len(files) > 1 {
		return "", false, fmt.Errorf("only one image can be uploaded, but %d images were provided", len(files))
	}

	if len(files) == 0 {
		fmt.Println("Warning: No image uploaded, using default one instead")
		return "", false, nil
	}

	file := files[0]

	filename := fmt.Sprintf("%d.jpeg", userID)
	var dst string

	if role == "customer" {
		dst = filepath.Join("images/customers", filename)
	} else {
		dst = filepath.Join("images/hosts", filename)
	}

	if err := context.SaveUploadedFile(file, dst); err != nil {
		return "", false, fmt.Errorf("error saving uploaded file: %v", err)
	}

	return dst, true, nil
}
