package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDatabase() {

	DbHost := os.Getenv("DATABASE_HOST")
	DbUser := os.Getenv("POSTGRES_USER")
	DbPassword := os.Getenv("POSTGRES_PASSWORD")
	DbName := os.Getenv("POSTGRES_DB")
	DbPort := os.Getenv("DATABASE_PORT")

	DBURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", DbHost, DbUser, DbPassword, DbName, DbPort)

	var err error

	DB, err = gorm.Open(postgres.Open(DBURL), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println("Cannot connect to database")
		log.Fatal("Connection error:", err)
	}
	fmt.Println("We are connected to the database")

	allCountries, err := GetAllCountries()
	if err != nil {
		log.Fatal("GetAllCountries", err)
	}
	if len(allCountries) == 0 {
		if err := AddCountries(); err != nil {
			log.Fatal("AddCountries:", err)
		}
	}

	allTownTypes, err := GetAllTownTypes()
	if err != nil {
		log.Fatal("GetAllTownTypes", err)
	}
	if len(allTownTypes) == 0 {
		if err := AddTownTypes(); err != nil {
			log.Fatal("AddTownTypes:", err)
		}
	}

	allGrades, err := GetGrades()
	if err != nil {
		log.Fatal("GetGrades", err)
	}
	if len(allGrades) == 0 {
		if err := AddGrades(); err != nil {
			log.Fatal("AddGrades:", err)
		}
	}

	allPaymentsMethod, err := GetAllPayments()
	if err != nil {
		log.Fatal("GetAllPayments", err)
	}
	if len(allPaymentsMethod) == 0 {
		if err := AddPayments(); err != nil {
			log.Fatal("AddPayments:", err)
		}
	}

	if err = DB.AutoMigrate(&Country{}, &TownType{}, &Grade{}, &Payment{}, &Customer{}, &Host{},
		&Town{}, &Discount{},
		&Animal{}, &Offer{}); err != nil {
		log.Fatal("Migration error:", err)
	}
}
