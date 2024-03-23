package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	Count        int
	Description  string
	Reservations []Reservation
}

var RatingList = []Rating{
	{Count: 1, Description: "Terrible"},
	{Count: 2, Description: "Bad"},
	{Count: 3, Description: "Mediocre"},
	{Count: 4, Description: "Good"},
	{Count: 5, Description: "Excellent"},
}

func (r *Rating) save() error {
	if err := DB.Create(&r).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}
	return nil
}

func AddGrades() error {
	for _, rating := range RatingList {
		if err := rating.save(); err != nil {
			return fmt.Errorf("grades.Save: %w", err)
		}
	}
	return nil
}

func GetGrades() ([]Rating, error) {
	var rating []Rating

	if err := DB.Model([]Rating{}).Scan(&rating).Error; err != nil {
		return rating, fmt.Errorf("DB.Model(&[]Rating{}).Scan: %w", err)
	}
	return rating, nil
}

func GetGradeByCount(count string) (*Rating, error) {
	var r *Rating
	if err := DB.Where("count = ?", count).First(&r).Error; err != nil {
		return nil, fmt.Errorf("DB.First: %w", err)
	}
	return r, nil
}
