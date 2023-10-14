package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm/schema"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {

	err := godotenv.Load("../.env")

	if err != nil {
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
	} else {
		fmt.Println("We are connected to the database")
	}

	err = DB.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Migration error:", err)
		return
	}

}
