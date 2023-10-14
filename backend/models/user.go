package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"size:100;not null;unique" json:"email"`
	Password string `gorm:"size:100;not null;" json:"password"`
}

func (u *User) SaveUser() (*User, error) {
	fmt.Println("Started saving")

	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	fmt.Println("Finished saving")

	return u, nil
}

func (u *User) BeforeSave(db *gorm.DB) error {
	fmt.Println("Starting before save")
	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	fmt.Println("Finished before save")

	return nil

}
