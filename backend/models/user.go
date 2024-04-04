package models

import (
	"errors"
	"fmt"
	"path/filepath"

	"backend/pkg/passHelper"
	"backend/utils/token"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Tabler interface {
	TableName() string
}

type User struct {
	gorm.Model
	FirstName                 string `gorm:"size:30;not null" form:"first_name" binding:"required,min=2,max=30"`
	LastName                  string `gorm:"size:100" form:"last_name"`
	Email                     string `gorm:"size:100;not null;unique" form:"email" binding:"required,email"`
	ImagePath                 string `gorm:"default:images/customers/default_image.png" form:"image_path" binding:"omitempty,url"`
	Password                  string `gorm:"size:100;not null" form:"password" binding:"required,min=8"`
	PasswordConfirmation      string `gorm:"-" form:"password_confirmation" binding:"required,eqfield=Password"`
	Role                      string `gorm:"size:20;not null;default:'customer'" form:"-" validate:"required,oneof=customer host"`
	Description               string `gorm:"not null;" form:"description"`
	PhoneNumber               string `gorm:"size:12;not null" form:"phone_number"`
	BankAccount               string `gorm:"size:31;not null" form:"bank_account"`
	Accommodations            []Accommodation
	Activities                []Activity
	Events                    []Event
	ReservationsAccommodation []ReservationAccommodation
	ReservationsActivity      []ReservationActivity
	ReservationsEvent         []ReservationEvent
	Messages                  []Message
}

func (User) TableName() string {
	return "app_user"
}

func GetUserByEmail(email string) (*User, error) {
	var u User

	if err := DB.Where("email = ?", email).First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("DB.Model.Where.Scan: %w", err)
	}

	return &u, nil
}

func GetUserById(id string) (*User, error) {
	var u User
	if err := DB.Model(&User{}).Where("id = ? AND role = host", id).Scan(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("DB.Model.Where.Scan: %w", err)
	}
	return &u, nil
}

func (u *User) Save() error {
	if err := DB.Create(&u).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}

	return nil
}

func (u *User) Update() error {
	if err := DB.Save(&u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) Delete() error {
	if err := DB.Delete(&u).Error; err != nil {
		return err
	}
	return nil
}

func LoginCheck(email, password string) (string, *User, error) {
	user, err := GetUserByEmail(email)
	if err != nil {
		return "", nil, fmt.Errorf("GetUserByEmail: %w", err)
	}

	if err := passHelper.VerifyPassword(password, user.Password); err != nil {
		return "", nil, fmt.Errorf("VerifyPassword: wrong password")
	}

	t, err := token.GenerateToken(user.ID, email, user.Role)
	if err != nil {
		return "", nil, fmt.Errorf("token.GenerateToken: %w", err)
	}

	return t, user, nil
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("bcrypt.GenerateFromPassword: %w", err)
	}
	u.Password = string(hashedPassword)

	return nil
}

func (u *User) HandleUserImageUploads(context *gin.Context, userID uint, role string) (string, bool, error) {
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
