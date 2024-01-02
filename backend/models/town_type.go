package models

import (
	"fmt"
	"math"

	"gorm.io/gorm"
)

type TownType struct {
	gorm.Model
	Name          string
	MinPopulation int
	MaxPopulation int
	Towns         []Town
}

var TownTypesList = []TownType{
	{Name: "Village", MinPopulation: 0, MaxPopulation: 5000},
	{Name: "Small Town", MinPopulation: 5001, MaxPopulation: 50000},
	{Name: "Medium Town", MinPopulation: 50001, MaxPopulation: 250000},
	{Name: "Big Town", MinPopulation: 250001, MaxPopulation: math.MaxInt},
}

func (tt *TownType) save() error {
	if err := DB.Create(&tt).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}
	return nil
}

func AddTownTypes() error {
	for _, townType := range TownTypesList {
		if err := townType.save(); err != nil {
			return fmt.Errorf("townType.Save: %w", err)
		}
	}
	return nil
}

func GetAllTownTypes() ([]TownType, error) {
	var tt []TownType

	if err := DB.Model([]TownType{}).Scan(&tt).Error; err != nil {
		return tt, fmt.Errorf("DB.Model(&[]TownType{}).Scan: %w", err)
	}
	return tt, nil
}
