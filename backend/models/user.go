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
	FirstName                 string `gorm:"not null" form:"first_name" binding:"required,min=2,max=30"`
	LastName                  string `gorm:"not null" form:"last_name"`
	Email                     string `gorm:"not null;unique" form:"email" binding:"required,email"`
	ImagePath                 string `gorm:"default:images/users/default_image.png" form:"image_path" binding:"omitempty,url"`
	Password                  string `gorm:"not null" form:"password" binding:"required,min=8"`
	PasswordConfirmation      string `gorm:"-" form:"password_confirmation" binding:"required,eqfield=Password"`
	Role                      string `gorm:"not null;default:'customer'" form:"-" binding:"required,oneof=customer host"`
	Description               string `gorm:"not null;" form:"description"`
	PhoneNumber               string `gorm:"not null;unique" form:"phone_number"`
	BankAccount               string `gorm:"not null;unique" form:"bank_account"`
	Accommodations            []Accommodation
	Activities                []Activity
	Events                    []Event
	ReservationsAccommodation []ReservationAccommodation
	ReservationsActivity      []ReservationActivity
	ReservationsEvent         []ReservationEvent
	ReservationsRoom          []ReservationRoom
}

func (*User) TableName() string {
	return "app_user"
}

func GetUserByEmail(email string) (*User, error) {
	var u User

	if err := DB.Where("email = ?", email).First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with email %s not found", email)
		}
		return nil, fmt.Errorf("failed to get user by email")
	}

	return &u, nil
}

func GetUserById(id string) (*User, error) {
	var u User
	if err := DB.Model(&User{}).Where("id = ?", id).Scan(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with ID %s not found: %w", id, err)
		}
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}
	return &u, nil
}

func (u *User) Save() error {
	if err := DB.Create(&u).Error; err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}
	return nil
}

func (u *User) Update() error {
	if err := DB.Save(&u).Error; err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (u *User) Delete() error {
	if err := DB.Delete(&u).Error; err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

func LoginCheck(email, password string) (string, *User, error) {
	user, err := GetUserByEmail(email)
	if err != nil {
		return "", nil, fmt.Errorf("user with given email does not exist")
	}

	if err := passHelper.VerifyPassword(password, user.Password); err != nil {
		return "", nil, fmt.Errorf("wrong password")
	}

	t, err := token.GenerateToken(user.ID, email, user.Role)
	if err != nil {
		return "", nil, fmt.Errorf("failed to generate token")
	}

	return t, user, nil
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password")
	}
	u.Password = string(hashedPassword)

	return nil
}

func (u *User) HandleUserImageUploads(context *gin.Context, userID uint) (string, bool, error) {
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
	dst := filepath.Join("images/users", filename)

	if err := context.SaveUploadedFile(file, dst); err != nil {
		return "", false, fmt.Errorf("error saving uploaded file: %v", err)
	}

	return dst, true, nil
}
