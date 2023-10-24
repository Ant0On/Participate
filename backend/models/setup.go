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

	if err = DB.AutoMigrate(&Customer{}); err != nil {
		log.Fatal("Migration error:", err)
		return
	}
}
