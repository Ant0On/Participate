package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	Count                     int
	Description               string
	ReservationsAccommodation []ReservationAccommodation
	ReservationActivity       []ReservationActivity
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
		return fmt.Errorf("failed to save rating: %w", err)
	}
	return nil
}

func AddGrades() error {
	for _, rating := range RatingList {
		if err := rating.save(); err != nil {
			return fmt.Errorf("failed to add rating: %w", err)
		}
	}
	return nil
}

func GetGrades() ([]Rating, error) {
	var ratings []Rating

	if err := DB.Model(&Rating{}).Scan(&ratings).Error; err != nil {
		return ratings, fmt.Errorf("failed to get ratings: %w", err)
	}
	return ratings, nil
}

func GetGradeByCount(count string) (*Rating, error) {
	var rating *Rating
	if err := DB.Where("count = ?", count).First(&rating).Error; err != nil {
		return nil, fmt.Errorf("failed to get rating by count: %w", err)
	}
	return rating, nil
}
