package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Equipment struct {
	gorm.Model
	Name     string
	Activity []Activity `gorm:"many2many:activity_equipment;"`
}

var EquipmentList = []Equipment{
	{Name: "Life Jacket"},
	{Name: "Kayak"},
	{Name: "Paddle"},
	{Name: "Helmet"},
	{Name: "Snowboard"},
	{Name: "Sled"},
	{Name: "Snowshoes"},
	{Name: "Tent"},
	{Name: "Sleeping Bag"},
	{Name: "Backpack"},
	{Name: "Hiking Boots"},
	{Name: "Compass"},
	{Name: "Map"},
	{Name: "Binoculars"},
	{Name: "Flashlight"},
	{Name: "First Aid Kit"},
	{Name: "Water Bottle"},
	{Name: "Climbing Harness"},
	{Name: "Climbing Rope"},
	{Name: "Carabiners"},
	{Name: "Rock Climbing Shoes"},
	{Name: "Fishing Rod"},
	{Name: "Bait"},
	{Name: "Camera"},
	{Name: "Swimsuit"},
	{Name: "Snorkel Gear"},
	{Name: "Tennis Racket"},
	{Name: "Golf Clubs"},
	{Name: "Bicycle"},
}

func (e *Equipment) save() error {
	if err := DB.Create(&e).Error; err != nil {
		return fmt.Errorf("DB.Create: %w", err)
	}
	return nil
}

func AddEquipment() error {
	for _, equipment := range EquipmentList {
		if err := equipment.save(); err != nil {
			return fmt.Errorf("country.Save: %w", err)
		}
	}
	return nil
}

func GetAllEquipment() ([]Equipment, error) {
	var equipment []Equipment

	if err := DB.Order("name").Find(&equipment).Error; err != nil {
		return nil, fmt.Errorf("DB.Order().Find(): %w", err)
	}
	return equipment, nil
}

func GetEquipmentByID(id string) (Equipment, error) {
	var equipment Equipment

	if err := DB.First(&equipment, id).Error; err != nil {
		return Equipment{}, fmt.Errorf("DB.First: %w", err)
	}
	return equipment, nil
}
