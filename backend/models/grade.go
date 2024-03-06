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

var GradesList = []Rating{
	{Count: 1, Description: "Terrible"},
	{Count: 2, Description: "Bad"},
	{Count: 3, Description: "Mediocre"},
	{Count: 4, Description: "Good"},
	{Count: 5, Description: "Excellent"},
}

func (g *Rating) save() error {
	if err := DB.Create(&g).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}
	return nil
}

func AddGrades() error {
	for _, grade := range GradesList {
		if err := grade.save(); err != nil {
			return fmt.Errorf("grades.Save: %w", err)
		}
	}
	return nil
}

func GetGrades() ([]Rating, error) {
	var grade []Rating

	if err := DB.Model([]Rating{}).Scan(&grade).Error; err != nil {
		return grade, fmt.Errorf("DB.Model(&[]Rating{}).Scan: %w", err)
	}
	return grade, nil
}

func GetGradeByCount(count string) (*Rating, error) {
	var g *Rating
	if err := DB.Where("count = ?", count).First(&g).Error; err != nil {
		return nil, fmt.Errorf("DB.First: %w", err)
	}
	return g, nil
}
