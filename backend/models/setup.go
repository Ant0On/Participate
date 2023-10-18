package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDatabase() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", DbHost, DbUser, DbPassword, DbName, DbPort)

	DB, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println("Cannot connect to database")
		log.Fatal("Connection error:", err)
	}
	fmt.Println("We are connected to the database")

	if err = DB.AutoMigrate(&Customer{}); err != nil {
		log.Fatal("Migration error:", err)
		return
	}
}
