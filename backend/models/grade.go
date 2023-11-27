package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Grade struct {
	gorm.Model
	Count        int
	Description  string
	Reservations []Reservation
}

var GradesList = []Grade{
	{Count: 1, Description: "Terrible"},
	{Count: 2, Description: "Bad"},
	{Count: 3, Description: "Mediocre"},
	{Count: 4, Description: "Good"},
	{Count: 5, Description: "Excellent"},
}

func (g *Grade) save() error {
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

func GetGrades() (*[]Grade, error) {
	var grade []Grade

	if err := DB.Model(&[]Grade{}).Scan(&grade).Error; err != nil {
		return &grade, fmt.Errorf("DB.Model(&[]Grade{}).Scan: %w", err)
	}
	return &grade, nil
}
